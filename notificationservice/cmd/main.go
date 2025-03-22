package cmd

import (
	"fmt"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/db/postgres"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/server"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/TemaStatham/TaskService/notificationservice/pkg/app/notification/service"
	"github.com/TemaStatham/TaskService/notificationservice/pkg/infra/config"
	postgresnotification "github.com/TemaStatham/TaskService/notificationservice/pkg/infra/db/postgres"
	grpc "github.com/TemaStatham/TaskService/notificationservice/pkg/infra/transport/internalapi"
	handler "github.com/TemaStatham/TaskService/notificationservice/pkg/infra/transport/publicapi"
)

func main() {
	cfg := config.NewConfig()

	db, err := postgres.NewPostgresGormDB(postgres.Config{
		Host:     cfg.DBConfig.Host,
		Port:     cfg.DBConfig.Port,
		Username: cfg.DBConfig.Username,
		Password: cfg.DBConfig.Password,
		DBName:   cfg.DBConfig.DBName,
		SSLMode:  cfg.DBConfig.SSLMode,
	})
	if err != nil {
		panic(err)
	}

	rep := postgresnotification.NewNotificationPostgres(db)
	srv := service.NewNotificationService(rep)
	hnd := handler.NewHandler(srv)

	fmt.Println("Client started")
	serve := new(server.Server)
	go func() {
		if err := serve.Run(cfg.HandlerPort, hnd.Init()); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		grpc.InitServer(cfg.ServerPort)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	return
}
