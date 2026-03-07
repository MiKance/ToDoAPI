package models

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MiKance/ToDoAPI/internal/config"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres"
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
	fmt.Println("Starting server on http://" + cfg.Host + ":" + cfg.Port)
	return server.httpServer.ListenAndServe()
}

func (server *Server) GracefulShutdown(ctx context.Context, storage *postgres.Storage) error {
	storage.Close()
	return server.httpServer.Shutdown(ctx)
}
