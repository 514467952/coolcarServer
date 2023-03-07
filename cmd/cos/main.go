package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func main() {
	u, _ := url.Parse("https://coolcar-1300226989.cos.ap-beijing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}

	secID := "AKIDRfieXFjpgGjz1sToupWyUklW0DCuKTM5"
	secKey := "dXEm5KL9sSmoa8cu55pIUdd4NWdXBZia"
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secID,
			SecretKey: secKey,
		},
	})

	//文件名
	name := "abc.jpg"
	//预签名URL
	presignedURL, err := client.Object.GetPresignedURL(
		context.Background(),
		http.MethodPut,
		name,
		secID,
		secKey,
		1*time.Hour,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(presignedURL)
}
