// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golifez/zkit/internal/biz"
	"github.com/golifez/zkit/internal/biz/aws"
	"github.com/golifez/zkit/internal/client"
	"github.com/golifez/zkit/internal/conf"
	"github.com/golifez/zkit/internal/data/aws"
	"github.com/golifez/zkit/internal/server"
	"github.com/golifez/zkit/internal/service"
	aws2 "github.com/golifez/zkit/internal/service/aws"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, config *conf.Config, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	autherUsecase := biz.NewAutherUsecase(config, logger)
	authService := service.NewAuthService(autherUsecase, logger)
	entClient := client.NewEntClient(confData, logger)
	clientData, cleanup, err := client.NewData(entClient, logger)
	if err != nil {
		return nil, nil, err
	}
	awsIamRepo := data.NewAwsIamRepo(clientData, config, logger)
	awsIamUsecase := aws.NewAwsIamUsecase(awsIamRepo, config, logger)
	iamService := aws2.NewIamService(awsIamUsecase, logger)
	serviceGrpcContainer := server.NewServiceGrpcContainer(authService, iamService)
	grpcServer := server.NewGRPCServer(confServer, serviceGrpcContainer, logger)
	serviceContainer := server.NewServiceContainer(authService)
	httpServer := server.NewHTTPServer(confServer, config, serviceContainer, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, httpServer, registrar, registry)
	return app, func() {
		cleanup()
	}, nil
}
