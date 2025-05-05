package sdk

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(ctx context.Context, port string, handler http.Handler) Server {
	return Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%v", port),
			Handler: handler,
			BaseContext: func(net.Listener) context.Context {
				return ctx
			},
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
