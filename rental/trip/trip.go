package trip

import (
	rentalpb "coolcar/rental/api/gen/v1/rental"
	tripdao "coolcar/rental/trip/dao"
	"coolcar/shared/id"

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

	// 创建行程: 写入数据库 + 计费
	poi, err := s.POIManager.Resolve(c, req.Start)
	if err != nil {
		s.Logger.Info("cannot resolve poi", zap.Stringer("location", req.Start), zap.Error(err))
	}

	ls := &rentalpb.LocationStatus{
		Loaction: req.Start,
		PoiName:  poi,
	}

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
	return nil, status.Error(codes.Unimplemented, "")
}
func (s *Service) GetTrips(c context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
func (s *Service) UpdateTrip(c context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	tid := id.TripID(req.Id)
	tr, err := s.Mongo.GetTrip(c, id.TripID(req.Id), aid)
	//更新行程，算钱
	if req.Current != nil {
		tr.Trip.Current = s.calcCurrentStatus(tr.Trip, req.Current)
	}
	if req.EndTrip {
		tr.Trip.End = tr.Trip.Current
		tr.Trip.Status = rentalpb.TripStatus_FINISHED
	}

	s.Mongo.UpdateTrip(c, tid, aid, tr.UpdateAt, tr.Trip)
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) calcCurrentStatus(trip *rentalpb.Trip, cur *rentalpb.Loaction) *rentalpb.LocationStatus {
	return nil
}
