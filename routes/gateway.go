package routes

import (
	"context"
	"net/http"

	configs "github.com/crowdeco/skeleton/configs"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type GRpcGateway struct {
	Servers []configs.Server
}

func (g *GRpcGateway) Handle(ctx context.Context, server *http.ServeMux, client *grpc.ClientConn) *http.ServeMux {
	mux := runtime.NewServeMux()

	for _, server := range g.Servers {
		server.GRpcHandler(ctx, mux, client)
	}

	server.Handle("/", mux)

	return server
}
