package repository

import (
	"fmt"
	"github.com/TemaStatham/TaskService/internal/model"
	"github.com/TemaStatham/TaskService/internal/repository/postgres"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	postgres.TaskRepository
}

func NewRepository() {

}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
	err = db.AutoMigrate(&model.Response{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
	err = db.AutoMigrate(&model.Comment{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
	err = db.AutoMigrate(&model.Category{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
	err = db.AutoMigrate(&model.Task{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
	err = db.AutoMigrate(&model.Organization{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
	fmt.Println("migrate success")
}
