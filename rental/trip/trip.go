package trip

import (
	"context"

	rentalpb "coolcar/rental/api/gen/v1/rental"
	sharedauth "coolcar/shared/auth"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
}

func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
	// get accountID from context
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		s.Logger.Info("create trip fail", zap.String("start", req.Start), zap.String("account_id", aid))
		return nil, err
	}
	s.Logger.Info("create trip success", zap.String("start", req.Start), zap.String("account_id", aid))
	return nil, status.Error(codes.Unimplemented, "")
}
