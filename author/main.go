package main

import (
	authpb "coolcar/author/api/gen/v1/author"
	"coolcar/author/auth"
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

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	//创建一个rpc服务对象
	s := grpc.NewServer()

	//authpb注册一个服务
	authpb.RegisterAuthServiceServer(s, &auth.Service{
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
