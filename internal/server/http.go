package server

import (
	"github.com/nextmicro/logger"
	v1 "github.com/nextmicro/next-layout/api/helloworld/v1"
	"github.com/nextmicro/next-layout/internal/service"

	"github.com/nextmicro/next/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(greeter *service.GreeterService, logger logger.Logger) *http.Server {
	srv := http.NewServer()
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
