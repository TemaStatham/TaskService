package app

import (
	"context"
	"github.com/TemaStatham/TaskService/internal/config"
	"github.com/TemaStatham/TaskService/internal/handler"
	"github.com/TemaStatham/TaskService/internal/repository"
	"github.com/TemaStatham/TaskService/internal/repository/postgres"
	"github.com/TemaStatham/TaskService/internal/service"
	"github.com/TemaStatham/TaskService/pkg/db"
	"github.com/TemaStatham/TaskService/pkg/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
}

func New() *App {
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
	repository.Migrate(dbg)

	taskRep := postgres.NewTaskPostgresRepository(dbg)
	commRep := postgres.NewCommentsRepository(dbg)
	respRep := postgres.NewResponseRepository(dbg)
	userRep := postgres.NewUserRepository(dbg)

	taskServ := service.NewTaskService(taskRep, userRep)
	commServ := service.NewCommentService(commRep)
	respServ := service.NewResponseService(respRep)
	// userServ := service.NewUserService(userRep)

	// todo: продумать логику получения организаций по кафке
	// todo: запихнуть создание топиков в докер
	// todo: протестить вебсокеты
	// todo: написать фенкциональные тесты
	// todo: разобраться с связью организаций и пользователей
	kafkaServ := service.NewKafkaService(*userRep)
	go kafkaServ.StartConsume(cfg.KConfig)

	hand := handler.NewTaskHandler(taskServ, respServ, commServ)

	serve := new(server.Server)
	go func() {
		if err := serve.Run(cfg.SConfig.Port, hand.Init(cfg.JWTSecret)); err != nil {
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
