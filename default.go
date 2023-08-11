package serviceHttpGrpc

import (
	"github.com/atom-providers/app"
	grpcServer "github.com/atom-providers/grpc-server"
	httpFiber "github.com/atom-providers/http-fiber"
	"github.com/atom-providers/log"
	"github.com/rogeecn/atom/container"
)

func Default(providers ...container.ProviderContainer) container.Providers {
	return append(container.Providers{
		app.DefaultProvider(),
		log.DefaultProvider(),
		httpFiber.DefaultProvider(),
		grpcServer.DefaultProvider(),
	}, providers...)
}
