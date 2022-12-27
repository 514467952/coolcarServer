package sharedserver

import (
	sharedauth "coolcar/shared/auth"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Name              string             //服务名称
	Addr              string             //服务端口地址
	AuthPublicKeyFile string             //公钥文件名称
	RegisterFunc      func(*grpc.Server) //对外暴露每个服务自己得注册函数
	Logger            *zap.Logger
}

func RunGRPCServer(c *GRPCConfig) error {
	nameField := zap.String("name", c.Name)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("cannot listen", nameField, zap.Error(err))
	}

	var opts []grpc.ServerOption
	if c.AuthPublicKeyFile != "" {
		in, err := sharedauth.Interceptor(c.AuthPublicKeyFile)
		if err != nil {
			c.Logger.Fatal("cannot create auth interceptor", nameField, zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}

	//创建一个rpc服务对象
	s := grpc.NewServer(opts...)
	c.RegisterFunc(s)

	c.Logger.Info("sharedserver started", nameField, zap.String("addr", c.Addr))
	return s.Serve(lis)
}
