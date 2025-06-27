//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/golifez/zkit/internal/biz"
	"github.com/golifez/zkit/internal/conf"
	"github.com/golifez/zkit/internal/server"
	"github.com/golifez/zkit/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
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
	// panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
