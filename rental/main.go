package main

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	profile "coolcar/rental/profile"
	profiledao "coolcar/rental/profile/dao"
	"coolcar/rental/trip"
	poi "coolcar/rental/trip/client"
	"coolcar/rental/trip/client/car"
	profileClient "coolcar/rental/trip/client/profile"
	tripdao "coolcar/rental/trip/dao"
	"coolcar/shared/sharedserver"
	"log"

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

	db := mongoClient.Database("coolcar")

	profService := &profile.Service{
		Mongo:  profiledao.NewMongo(db),
		Logger: logger,
	}

	err = sharedserver.RunGRPCServer(&sharedserver.GRPCConfig{
		Name:              "rental",
		Addr:              ":8082",
		AuthPublicKeyFile: "shared/auth/public.key",
		Logger:            logger,
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				CarManager: &car.Manager{},
				ProfileManager: &profileClient.Manager{
					Fetcher: profService,
				},
				POIManager: &poi.Manager{},
				Mongo:      tripdao.NewMongo(db),
				Logger:     logger,
			})
			rentalpb.RegisterProfileServiceServer(s, profService)
		},
	})

	logger.Fatal("cannot start rental server", zap.Error(err))
}
