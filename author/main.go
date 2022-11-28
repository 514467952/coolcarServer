package main

import (
	"context"
	authpb "coolcar/author/api/gen/v1/author"
	"coolcar/author/auth"
	"coolcar/author/dao"
	"coolcar/author/wechat"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	//建立与mongoDB数据库的链接
	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://106.54.49.241:27017/?readPreference=primary&ssl=false&directConnection=true"))
	if err != nil {
		logger.Fatal("auth main cannot connect mongodb", zap.Error(err))
	}

	//创建一个rpc服务对象
	s := grpc.NewServer()

	//authpb注册auth服务
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID:     "wx2e157f9c5eef2403",
			AppSecret: "b9db8bdde1b98a0e1749b936c3549c43",
		},
		MyMongo: dao.NewMongo(mongoClient.Database("coolcar")),
		Logger:  logger,
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
