package profile

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	"coolcar/rental/profile/dao"

	sharedauth "coolcar/shared/auth"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Mongo  *dao.Mongo
	Logger *zap.Logger
}

//获取个人信息
func (s *Service) GetProfile(c context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	p, err := s.Mongo.GetProfile(c, aid)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &rentalpb.Profile{}, nil
		}
		s.Logger.Error("Service cannot get profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return p, nil
}

//提交个人信息
func (s *Service) SubmitProfile(c context.Context, i *rentalpb.Identity) (*rentalpb.Profile, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	p := &rentalpb.Profile{
		Identity:       i,
		IdentityStatus: rentalpb.IdentityStatus_PENDING,
	}

	err = s.Mongo.UpdateProfile(c, aid, rentalpb.IdentityStatus_UNSUBMITTED, p)
	if err != nil {
		s.Logger.Error("SubmitProfile cannot update profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return p, nil
}

//清空个人信息
func (s *Service) ClearProfile(c context.Context, req *rentalpb.ClearProfileRequest) (*rentalpb.Profile, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	p := &rentalpb.Profile{}

	err = s.Mongo.UpdateProfile(c, aid, rentalpb.IdentityStatus_VERIFIED, p)
	if err != nil {
		s.Logger.Error("ClearProfile cannot update profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return p, nil
}
