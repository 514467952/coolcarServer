package main

import (
	"context"
	blobpb "coolcar/blob/api/gen/v1/blob"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("106.54.49.241:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	c := blobpb.NewBlobServiceClient(conn)

	ctx := context.Background()
	// 测试上传图片
	// res, err := c.CreateBlob(ctx, &blobpb.CreateBlobRequest{
	// 	AccountId:           "account_2",
	// 	UploadUrlTimeoutSec: 1000,
	// })

	//测试获取图片
	// res, err := c.GetBlob(ctx, &blobpb.GetBlobRequest{
	// 	Id: "6420f51571ad01063a33fb0c",
	// })

	//测试展示图片
	res, err := c.GetBlobURL(ctx, &blobpb.GetBlobURLRequest{
		Id:         "6420f51571ad01063a33fb0c",
		TimeoutSec: 100,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
}
