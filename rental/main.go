package main

import (
	rentalpb "coolcar/rental/api/gen/v1/rental"
	"coolcar/rental/trip"
	sharedauth "coolcar/shared/auth"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//创建日志对象
	logger, err := newZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger:%v", err)
	}

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	in, err := sharedauth.Interceptor("shared/auth/public.key")
	if err != nil {
		logger.Fatal("cannot create auth interceptor", zap.Error(err))
	}

	//创建一个rpc服务对象
	s := grpc.NewServer(grpc.UnaryInterceptor(in))
	//注册rental服务
	rentalpb.RegisterTripServiceServer(s, &trip.Service{
		Logger: logger,
	})

	s.Serve(lis)
	logger.Fatal("cannot server", zap.Error(err))
}

//自定义日志
func newZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
}
