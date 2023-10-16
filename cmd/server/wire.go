//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/nextmicro/logger"
	"github.com/nextmicro/next"
	"github.com/nextmicro/next-layout/internal/biz"
	"github.com/nextmicro/next-layout/internal/conf"
	"github.com/nextmicro/next-layout/internal/data"
	"github.com/nextmicro/next-layout/internal/server"
	"github.com/nextmicro/next-layout/internal/service"

	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, logger.Logger) (*next.Next, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
