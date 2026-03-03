package main

import (
	"github.com/MiKance/ToDoAPI/internal/config"
	"github.com/MiKance/ToDoAPI/internal/handlers"
	"github.com/MiKance/ToDoAPI/internal/models"
	"github.com/MiKance/ToDoAPI/internal/repository"
	"github.com/MiKance/ToDoAPI/internal/repository/postgres"
	"github.com/MiKance/ToDoAPI/internal/service"

	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg := config.MustNewConfig("./configs/local.yaml")

	ctx := context.Background()
	ctxSt, cancel := context.WithTimeout(ctx, 5*time.Second)
	_ = postgres.NewStorage(ctxSt, cfg.Storage)

	repo := repository.NewRepository()
	serv := service.NewService(repo)
	handler := handlers.NewHandler(serv)

	router := handler.InitRouter()

	server := new(models.Server)
	go func() {
		server.Start(cfg.Server, router)
	}()

	fmt.Println("Server started")

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.GracefulShutdown(ctx); err != nil {
		fmt.Println("Graceful shutdown failed:", err)
	}
	fmt.Println("Server gracefully stopped")
}
