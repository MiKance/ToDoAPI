package main

import (
	"ToDoAPI/internal/config"
	"ToDoAPI/internal/handlers"
	"ToDoAPI/internal/models"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg := config.MustNewConfig("./configs/local.yaml")

	handl := new(handlers.Handler)
	router := handl.InitRouter()

	server := new(models.Server)
	go func() {
		server.Start(cfg.Server, router)
	}()

	fmt.Println("Server started")

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.GracefulShutdown(ctx); err != nil {
		fmt.Println("Graceful shutdown failed:", err)
	}
	fmt.Println("Server gracefully stopped")
}
