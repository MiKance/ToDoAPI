package models

import (
	"context"
	"net/http"

	"github.com/MiKance/ToDoAPI/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func (server *Server) Start(cfg *config.ServerConfig, handler http.Handler) error {
	server.httpServer = &http.Server{
		Addr:           cfg.Host + ":" + cfg.Port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
	}
	return server.httpServer.ListenAndServe()
}

func (server *Server) GracefulShutdown(ctx context.Context) error {
	return server.httpServer.Shutdown(ctx)
}
