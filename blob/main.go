package main

import (
	"context"
	blobpb "coolcar/blob/api/gen/v1/blob"
	"coolcar/blob/blob"
	"coolcar/blob/dao"
	"coolcar/shared/sharedserver"
	"flag"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var mongoURI = flag.String("mongo_uri", "mongodb://localhost:27017", "mongo uri")

func main() {
	logger, err := sharedserver.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger:%v", err)
	}

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI(*mongoURI))
	if err != nil {
		logger.Fatal(("cannot connect mongodb"))
	}
	db := mongoClient.Database("coolcar")

	logger.Sugar().Fatal(sharedserver.RunGRPCServer(&sharedserver.GRPCConfig{
		Name:   "blob",
		Addr:   ":8083",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			blobpb.RegisterBlobServiceServer(s, &blob.Service{
				Mongo:  dao.NewMongo(db),
				Logger: logger,
			})
		},
	}))
}
