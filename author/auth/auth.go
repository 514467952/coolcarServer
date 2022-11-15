package auth

/*
登录GRPC服务
*/
import (
	"context"
	authpb "coolcar/author/api/gen/v1/author"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service
type Service struct {
	OpenIDResolver OpenIDResolver
	Logger         *zap.Logger
}

// to get open id
type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

// Login logs a user in
func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	openID, err := s.OpenIDResolver.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot get openid:%v", err)
	}
	return &authpb.LoginResponse{
		AccessToken:  "token for" + openID,
		ExpiresInSec: 7200,
	}, nil
}
