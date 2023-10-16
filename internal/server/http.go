package server

import (
	"github.com/nextmicro/logger"
	v1 "github.com/nextmicro/next-layout/api/helloworld/v1"
	"github.com/nextmicro/next-layout/internal/conf"
	"github.com/nextmicro/next-layout/internal/service"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/nextmicro/next/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger logger.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
