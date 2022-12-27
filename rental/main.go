package main

import (
	rentalpb "coolcar/rental/api/gen/v1/rental"
	"coolcar/rental/trip"
	"coolcar/shared/sharedserver"
	"log"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//创建日志对象
	logger, err := sharedserver.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger:%v", err)
	}

	err = sharedserver.RunGRPCServer(&sharedserver.GRPCConfig{
		Name:              "rental",
		Addr:              ":8082",
		AuthPublicKeyFile: "shared/auth/public.key",
		Logger:            logger,
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				Logger: logger,
			})
		},
	})

	logger.Fatal("cannot start rental server", zap.Error(err))
}
