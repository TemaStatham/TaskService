package main

import (
	"fmt"
	"github.com/TemaStatham/TaskService/client/pkg/infrastructure/transport/handler"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/TemaStatham/TaskService/client/pkg/infrastructure"
	"github.com/TemaStatham/TaskService/client/pkg/infrastructure/config"
	"github.com/TemaStatham/TaskService/client/pkg/infrastructure/server"
)

// go run ./cmd/main.go --config="./config.yaml"

func main() {
	cfg := config.MustLoad()
	container := infrastructure.NewContainer(*cfg)

	hand := handler.NewTaskHandler(
		container.ResponseQuery,
		container.ResponseService,
		container.CommentQuery,
		container.CommentService,
		container.TaskQuery,
		container.TaskService,
		container.ApproveService,
		container.TaskUserService,
		container.TaskUserQuery,
	)

	fmt.Println("Client started")
	serve := new(server.Server)
	go func() {
		if err := serve.Run(cfg.SConfig.Port, hand.Init(cfg.JWTSecret)); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	return
}
