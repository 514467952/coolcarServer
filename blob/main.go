package main

import (
	"context"
	blobpb "coolcar/blob/api/gen/v1/blob"
	"coolcar/blob/blob"
	"coolcar/blob/cos"
	"coolcar/blob/dao"
	"coolcar/shared/sharedserver"
	"flag"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
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

	st, err := cos.NewClient("https://coolcar-1300226989.cos.ap-beijing.myqcloud.com", "AKIDRfieXFjpgGjz1sToupWyUklW0DCuKTM5", "dXEm5KL9sSmoa8cu55pIUdd4NWdXBZia")
	if err != nil {
		logger.Fatal("cannot create cos server", zap.Error(err))
	}
	logger.Sugar().Fatal(sharedserver.RunGRPCServer(&sharedserver.GRPCConfig{
		Name:   "blob",
		Addr:   ":8083",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			blobpb.RegisterBlobServiceServer(s, &blob.Service{
				Storage: st,
				Mongo:   dao.NewMongo(db),
				Logger:  logger,
			})
		},
	}))
}
