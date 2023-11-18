//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/nextmicro/logger"
	"github.com/nextmicro/next-layout/internal/biz"
	"github.com/nextmicro/next-layout/internal/conf"
	"github.com/nextmicro/next-layout/internal/data"
	"github.com/nextmicro/next-layout/internal/server"
	"github.com/nextmicro/next-layout/internal/service"
	"github.com/nextmicro/next-layout/internal/svc"

	"github.com/google/wire"
)

// wireApp init next application.
func wireApp(*conf.Data, logger.Logger) (*Injector, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, svc.ProviderSet, InjectorSet, newApp))

	return new(Injector), nil, nil
}
