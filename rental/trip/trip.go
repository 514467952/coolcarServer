package trip

import (
	rentalpb "coolcar/rental/api/gen/v1/rental"
	tripdao "coolcar/rental/trip/dao"
	"coolcar/shared/id"
	"math/rand"
	"time"

	sharedauth "coolcar/shared/auth"

	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	ProfileManager ProfileManager
	CarManager     CarManager
	POIManager     POIManager
	Mongo          *tripdao.Mongo
	Logger         *zap.Logger
}

//ProfileManager defines the ACL (Anti Corruption Layer) 反腐败层
// for profile verification logic (资格认证的逻辑)
type ProfileManager interface {
	//认证方法:通过Verify函数告诉我们是否有资格
	//返回IdentityID 租车时，可能会重新审查，换身份信息
	Verify(context.Context, id.AccountID) (id.IdentityID, error)
}

//地点管理
type POIManager interface {
	Resolve(context.Context, *rentalpb.Loaction) (string, error)
}

type CarManager interface {
	//检查车辆是否被租用
	Verify(context.Context, id.CarID, *rentalpb.Loaction) error
	//开锁
	Unlock(context.Context, id.CarID) error
}

func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	if req.CarId == "" || req.Start == nil {
		return nil, status.Error(codes.InvalidArgument, "CreateTrip error")
	}

	// 验证驾驶者身份
	iID, err := s.ProfileManager.Verify(c, aid)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	// 检查车辆状态
	carID := id.CarID(req.CarId)
	err = s.CarManager.Verify(c, carID, req.Start)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	//复用计算费用的方法，
	ls := s.calcCurrentStatus(c, &rentalpb.LocationStatus{
		Loaction:     req.Start,
		TimestampSec: nowFunc(),
	}, req.Start)

	tr, err := s.Mongo.CreateTrip(c, &rentalpb.Trip{
		AccountID:  aid.String(),
		CarID:      carID.String(),
		IdentityId: iID.String(),
		Status:     rentalpb.TripStatus_IN_PROGRESS,
		Start:      ls,
		Current:    ls,
	})
	if err != nil {
		s.Logger.Warn("cannot create trip", zap.Error(err))
		return nil, status.Error(codes.AlreadyExists, "")
	}

	// 车辆开锁
	go func() {
		err = s.CarManager.Unlock(context.Background(), carID)
		if err != nil {
			s.Logger.Error("cannot unlock car", zap.Error(err))
		}
	}()

	return &rentalpb.TripEntity{
		Id:   tr.ID.Hex(),
		Trip: tr.Trip,
	}, nil
}
func (s *Service) GetTrip(c context.Context, req *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	tr, err := s.Mongo.GetTrip(c, id.TripID(req.Id), aid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}

	return tr.Trip, nil
}
func (s *Service) GetTrips(c context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	trips, err := s.Mongo.GetTrips(c, aid, req.Status)
	if err != nil {
		s.Logger.Error("cannot get trips", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	res := &rentalpb.GetTripsResponse{}

	for _, tr := range trips {
		res.Trips = append(res.Trips, &rentalpb.TripEntity{
			Id:   tr.ID.Hex(),
			Trip: tr.Trip,
		})
	}
	return res, nil
}
func (s *Service) UpdateTrip(c context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	tid := id.TripID(req.Id)
	tr, err := s.Mongo.GetTrip(c, id.TripID(req.Id), aid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "GetTrip not found")
	}

	if tr.Trip.Current == nil {
		s.Logger.Error("trip without current set", zap.String("id", tid.String()))
		return nil, status.Error(codes.Internal, "tr.Trip.Current is null")
	}

	//更新行程，算钱
	cur := tr.Trip.Current.Loaction
	if req.Current != nil {
		cur = req.Current
	}

	tr.Trip.Current = s.calcCurrentStatus(c, tr.Trip.Current, cur)

	if req.EndTrip {
		tr.Trip.End = tr.Trip.Current
		tr.Trip.Status = rentalpb.TripStatus_FINISHED
	}

	err = s.Mongo.UpdateTrip(c, tid, aid, tr.UpdateAt, tr.Trip)
	if err != nil {
		s.Logger.Error("s.Mongo.UpdateTrip failed", zap.String("id", tid.String()))
		return nil, status.Error(codes.Aborted, "")
	}
	return tr.Trip, nil
}

var nowFunc = func() int64 {
	return time.Now().Unix()
}

const (
	centsPerSec = 0.7
	kmPerSec    = 0.02 //设定每s中跑0.02km
)

func (s *Service) calcCurrentStatus(c context.Context, last *rentalpb.LocationStatus, cur *rentalpb.Loaction) *rentalpb.LocationStatus {
	now := nowFunc()
	elapsedSec := float64(now - last.TimestampSec)
	poi, err := s.POIManager.Resolve(c, cur)
	if err != nil {
		s.Logger.Info("cannot resolve poi", zap.Stringer("location", cur), zap.Error(err))
	}

	//Float64时0~1之间的随机数，长期来说，是0.5，所以*2为1
	//长期看是想当于没有乘
	return &rentalpb.LocationStatus{
		Loaction:     cur,
		FeeCent:      last.FeeCent + int32(centsPerSec*elapsedSec*2*rand.Float64()),
		KmDriven:     last.KmDriven + kmPerSec*elapsedSec*2*rand.Float64(),
		TimestampSec: now,
		PoiName:      poi,
	}
}
