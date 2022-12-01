package mongotesting

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

const (
	image         = "mongo:4.4"
	containerPort = "27017/tcp"
)

func RunWithMongoInDocker(m *testing.M, mongoURI *string) int {
	c, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	//创建一个容器
	resp, err := c.ContainerCreate(ctx, &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			containerPort: {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding{
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

	containerID := resp.ID
	//销毁容器在整个方法执行完后销毁
	defer func() {
		fmt.Println("杀掉容器,killing container")
		err := c.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
			Force: true,
		})
		fmt.Printf("销毁容器失败:error removing container:%v", err)
		if err != nil {
			panic(err)
		}
	}()

	//启动容器
	err = c.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("容器开始运行:container started")
	time.Sleep(5 * time.Second)

	//因为创建容器时是随机端口，需要找到对应容器在哪个端口上启动
	inspRes, err := c.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}

	hostPort := inspRes.NetworkSettings.Ports[containerPort][0]
	//mongoURI告诉外面mongoDB在哪里启动
	*mongoURI = fmt.Sprintf("mongodb://%s:%s", hostPort.HostIP, hostPort.HostPort)
	fmt.Printf("listening at %+v\n", inspRes.NetworkSettings.Ports[containerPort][0])

	return m.Run()
}
