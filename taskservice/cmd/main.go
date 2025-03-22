package main

import (
	"fmt"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/handler"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/server"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/config"
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
