package app

import (
	"context"
	"github.com/TemaStatham/TaskService/internal/config"
	"github.com/TemaStatham/TaskService/internal/handler"
	"github.com/TemaStatham/TaskService/internal/repository/postgres"
	"github.com/TemaStatham/TaskService/internal/service"
	"github.com/TemaStatham/TaskService/pkg/db"
	"github.com/TemaStatham/TaskService/pkg/kafka"
	"github.com/TemaStatham/TaskService/pkg/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
}

func New(env string) *App {
	return &App{}
}

func (a *App) MustRun(cfg *config.Config) {
	if err := a.Run(cfg); err != nil {
		panic(err)
	}
}

func (a *App) Run(cfg *config.Config) error {
	dbg, err := db.NewPostgresGormDB(db.Config{
		Host:     cfg.DBConfig.Host,
		Port:     cfg.DBConfig.Port,
		Username: cfg.DBConfig.Username,
		Password: cfg.DBConfig.Password,
		DBName:   cfg.DBConfig.DBName,
		SSLMode:  cfg.DBConfig.SSLMode,
	})

	if err != nil {
		return err
	}

	taskRep := postgres.NewTaskPostgresRepository(dbg)
	commRep := postgres.NewCommentsRepository(dbg)
	respRep := postgres.NewResponseRepository(dbg)
	userRep := postgres.NewUserRepository(dbg)

	taskServ := service.NewTaskService(taskRep, userRep)
	commServ := service.NewCommentService(commRep)
	respServ := service.NewResponseService(respRep)
	userServ := service.NewUserService(userRep)

	kafka.StartKafkaConsumer([]string{"localhost:9092"}, "users-topic", "", func(key, value []byte) {
		err := userServ.CreateFromKafka(key, value)
		if err != nil {
			log.Println(err)
		}
	})

	hand := handler.NewTaskHandler(taskServ, respServ, commServ)

	serve := new(server.Server)
	go func() {
		if err := serve.Run(cfg.SConfig.Port, hand.Init()); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := serve.Shutdown(context.Background()); err != nil {
		return err
	}
	sqlDB, err := dbg.DB()
	if err != nil {
		log.Fatal("Ошибка получения SQL DB:", err)
	}
	if err := sqlDB.Close(); err != nil {
		return err
	}

	return nil
}
