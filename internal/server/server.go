package server

import (
	"context"
	"fmt"
	"github.com/valyala/fasthttp"
	"go-recipe/internal/config"
)

type Server struct {
	httpServer *fasthttp.Server
	r          func(ctx *fasthttp.RequestCtx)
}

func New(r func(ctx *fasthttp.RequestCtx)) *Server {
	return &Server{
		r: r,
	}
}

func (s *Server) Run(cfg *config.Config) error {
	s.httpServer = &fasthttp.Server{
		Handler: s.r,
	}

	return s.httpServer.ListenAndServe(fmt.Sprintf(":%d", cfg.HttpPort))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.ShutdownWithContext(ctx)
}
