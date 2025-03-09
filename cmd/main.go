package main

import (
	"github.com/TemaStatham/TaskService/internal/app"
	"github.com/TemaStatham/TaskService/internal/config"
	"log"
	"time"
)

// go run ./cmd/main.go --config="./config.yaml"

func main() {
	log.Println("asd")
	cfg := config.Config{
		Env:          "development",
		StoragePaths: "./storage",
		TokenTTL:     time.Duration(24 * time.Hour),
		Port:         8080,
		SConfig: config.ServerConfig{
			Port: "8080",
			Host: "127.0.0.1",
		},
		DBConfig: config.DataBaseConfig{
			Host:     "127.0.0.1",
			Port:     "5432",
			Username: "postgres",
			Password: "dapomogu_password",
			SSLMode:  "disable",
			DBName:   "dapomogu_task_db",
		},
	}

	a := app.New(cfg.Env)

	a.MustRun(&cfg)
}
