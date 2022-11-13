package auth

/*
登录GRPC服务
*/
import (
	"context"
	authpb "coolcar/author/api/gen/v1/author"

	"go.uber.org/zap"
)

// Service
type Service struct {
	Logger *zap.Logger
}

// Login logs a user in
func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code", zap.String("code", req.Code))
	return &authpb.LoginResponse{
		AccessToken:  "token for" + req.Code,
		ExpiresInSec: 7200,
	}, nil
}
