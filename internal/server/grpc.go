package server

import (
	v1 "github.com/golifez/zkit/api/auth/v1"
	"github.com/golifez/zkit/internal/conf"
	"github.com/golifez/zkit/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type ServiceGrpcContainer struct {
	AuthService *service.AuthService
	// UserService  *service.UserService
	// 其他服务...
}

func NewServiceGrpcContainer(
	auth *service.AuthService,
	// register *service.RegisterService,
	// 其他服务...
) *ServiceGrpcContainer {
	return &ServiceGrpcContainer{
		AuthService: auth,
		// UserService:  user,
	}
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, sc *ServiceGrpcContainer, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterAuthServiceServer(srv, sc.AuthService)
	return srv
}
