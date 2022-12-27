package main

import (
	"context"
	authpb "coolcar/author/api/gen/v1/author"
	"coolcar/author/auth"
	"coolcar/author/dao"
	token "coolcar/author/token"
	"coolcar/author/wechat"
	"coolcar/shared/sharedserver"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//创建日志对象
	logger, err := sharedserver.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger:%v", err)
	}

	//建立与mongoDB数据库的链接
	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://106.54.49.241:27017/?readPreference=primary&ssl=false&directConnection=true"))
	if err != nil {
		logger.Fatal("auth main cannot connect mongodb", zap.Error(err))
	}

	//读取私钥
	pkFile, err := os.Open("author/private.key")
	if err != nil {
		logger.Fatal("cannot open private key", zap.Error(err))
	}
	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("cannot read private key", zap.Error(err))
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}

	err = sharedserver.RunGRPCServer(&sharedserver.GRPCConfig{
		Name:   "auth",
		Addr:   ":8081",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			authpb.RegisterAuthServiceServer(s, &auth.Service{
				OpenIDResolver: &wechat.Service{
					AppID:     "wx2e157f9c5eef2403",
					AppSecret: "b9db8bdde1b98a0e1749b936c3549c43",
				},
				MyMongo:        dao.NewMongo(mongoClient.Database("coolcar")),
				Logger:         logger,
				TokenExpire:    2 * time.Hour,
				TokenGenerator: token.NewJWTTokenGen("coolcar/auth", privateKey),
			})
		},
	})

	logger.Fatal("cannot start auth server", zap.Error(err))
}
