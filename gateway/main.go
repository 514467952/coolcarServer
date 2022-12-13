package main

/*
gateWay服务，将grpc服务变成http请求，让小程序可以调用到
*/

import (
	"context"
	authpb "coolcar/author/api/gen/v1/author"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers: true,
				UseProtoNames:  false,
			},
		},
	))

	//监听8081的auth服务
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(
		c, mux, "localhost:8081",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalf("cannnot register auth service,%v", err)
	}

	//注册8082端口的rental服务
	err = rentalpb.RegisterTripServiceHandlerFromEndpoint(
		c, mux, "localhost:8082",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalf("cannnot register rental service,%v", err)
	}
	log.Fatal(http.ListenAndServe(":8080", mux))
}
