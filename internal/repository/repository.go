package repository

import "github.com/TemaStatham/TaskService/internal/repository/postgres"

type Repository struct {
	postgres.TaskRepository
}

func NewRepository() {

}
