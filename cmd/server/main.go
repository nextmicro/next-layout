package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/nextmicro/logger"
	"github.com/nextmicro/next"
	"github.com/nextmicro/next-layout/internal/conf"
	"github.com/nextmicro/next/config"
	"github.com/nextmicro/next/transport/grpc"
	"github.com/nextmicro/next/transport/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger logger.Logger, gs *grpc.Server, hs *http.Server) (*next.Next, error) {
	hs.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{
		DisableCompression: true,
		EnableOpenMetrics:  true,
	}))

	return next.New(
		next.ID(id),
		next.Name(Name),
		next.Version(Version),
		next.Metadata(map[string]string{}),
		next.Logger(logger),
		next.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	c, err := config.Init(filepath.Join(flagconf, config.BizConfFile()))
	if err != nil {
		panic(err)
	}
	defer c.Close()

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	injector, cleanup, err := wireApp(bc.Data, logger.DefaultLogger)
	if err != nil {
		panic(err)
	}

	err = injector.serviceContext.Run()
	if err != nil {
		panic(err)
	}

	defer cleanup()

	// start and wait for stop signal
	if err = injector.Next.Run(); err != nil {
		panic(err)
	}
}
