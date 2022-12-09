package auth

/*
登录GRPC服务
*/
import (
	"context"
	authpb "coolcar/author/api/gen/v1/author"
	"coolcar/author/dao"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service
type Service struct {
	OpenIDResolver OpenIDResolver
	MyMongo        *dao.MyMongo
	TokenGenerator TokenGenerator
	TokenExpire    time.Duration
	Logger         *zap.Logger
}

// to get open id
type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

//TokenGenerator generates a token for the specified account
type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, error)
}

// Login logs a user in
func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	openID, err := s.OpenIDResolver.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot get openid:%v", err)
	}

	accountID, err := s.MyMongo.ResolveAccountID(c, openID)
	if err != nil {
		s.Logger.Error("Login ResolveAccountID fail id", zap.Error(err))
		return nil, status.Error(codes.Internal, "Login ResolveAccountID fail")
	}

	tkn, err := s.TokenGenerator.GenerateToken(accountID, s.TokenExpire)
	if err != nil {
		s.Logger.Error("cannot generate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		AccessToken:  tkn,
		ExpiresInSec: int32(s.TokenExpire.Seconds()),
	}, nil
}
