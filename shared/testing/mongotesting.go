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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	image         = "mongo:4.4"
	containerPort = "27017/tcp"
)

var mongoURI string

const defaultMongoURI = "mongodb://localhost:27017"

func RunWithMongoInDocker(m *testing.M) int {
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
		if err != nil {
			fmt.Printf("销毁容器失败:error removing container:%v", err)
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
	mongoURI = fmt.Sprintf("mongodb://%s:%s", hostPort.HostIP, hostPort.HostPort)
	fmt.Printf("listening at %+v\n", inspRes.NetworkSettings.Ports[containerPort][0])

	return m.Run()
}

//创建新的docker容器中的mongo客户端
func NewClient(c context.Context) (*mongo.Client, error) {
	if mongoURI == "" {
		return nil, fmt.Errorf("mongo uri not set. Please run RunWithMongoInDocker in TestMain")
	}
	return mongo.Connect(c, options.Client().ApplyURI(mongoURI))
}

//创建本地的mongo客户端
func NewDefaultClient(c context.Context) (*mongo.Client, error) {
	return mongo.Connect(c, options.Client().ApplyURI(defaultMongoURI))
}

//建立索引
func SetupIndexes(c context.Context, d *mongo.Database) error {
	_, err := d.Collection("account").Indexes().CreateOne(c, mongo.IndexModel{
		Keys: bson.D{
			{Key: "open_id", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	_, err = d.Collection("trip").Indexes().CreateOne(c, mongo.IndexModel{
		Keys: bson.D{
			{Key: "trip.accountid", Value: 1},
			{Key: "trip.status", Value: 1},
		},
		Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
			"trip.status": 1,
		}),
	})
	return err
}
