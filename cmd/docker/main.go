package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	c, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	//创建一个容器
	resp, err := c.ContainerCreate(ctx, &container.Config{
		Image: "mongo:4.4",
		ExposedPorts: nat.PortSet{
			"27017/tcp": {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"27017/tcp": []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: "0",
				},
			},
		},
	}, nil, nil, "lxy_coolcar")

	if err != nil {
		panic(err)
	}

	//启动容器
	err = c.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("容器开始运行:container started")
	time.Sleep(5 * time.Second)

	//因为创建容器时是随机端口，需要找到对应容器在哪个端口上启动
	inspRes, err := c.ContainerInspect(ctx, resp.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listening at %+v\n", inspRes.NetworkSettings.Ports["27017/tcp"][0])

	fmt.Println("杀掉容器,killing container")
	err = c.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}
