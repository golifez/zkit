//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"

	// "github.com/go-kratos/kratos/v2/registry"

	"github.com/golifez/zkit/internal/biz"
	"github.com/golifez/zkit/internal/client"
	"github.com/golifez/zkit/internal/conf"
	"github.com/golifez/zkit/internal/data"
	"github.com/golifez/zkit/internal/server"
	"github.com/golifez/zkit/internal/service"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(
	*conf.Server,
	*conf.Data,
	*conf.Config,
	*conf.Registry,
	log.Logger,
) (*kratos.App, func(), error) {
	panic(wire.Build(
		// 工具类
		// utils.ProviderSet,

		// 基础设施层
		client.ProviderSet,
		server.SrvProviderSet,

		// 数据访问层
		data.DataProviderSet,

		// 业务逻辑层
		biz.BizProviderSet,

		// 服务实现层
		service.SvcProviderSet,

		// 应用组装
		newApp,
	))
}
