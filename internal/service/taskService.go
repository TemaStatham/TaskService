package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/TemaStatham/TaskService/internal/handler/request"
	"github.com/TemaStatham/TaskService/internal/model"
	"github.com/TemaStatham/TaskService/pkg/paginate"
	"github.com/TemaStatham/TaskService/pkg/roles"
)

var (
	ErrTaskIdIsEmpty      = errors.New("task id is empty")
	ErrTaskNameIsEmpty    = errors.New("task name is empty")
	ErrUserIsNotValid     = errors.New("user is not found")
	ErrUserIsNotValidRole = errors.New("user is not found")
)

type TaskRepository interface {
	Get(ctx context.Context, id uint) (*model.Task, error)
	GetAll(ctx context.Context, paginat *paginate.Pagination, user uint) (*paginate.Pagination, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, dto request.UpdateTaskRequest) error
	Create(ctx context.Context, dto request.CreateTaskRequest) (uint, error)
}

type TaskService struct {
	TaskRepository
	UserRepository
}

func NewTaskService(taskRepository TaskRepository, userRepository UserRepository) *TaskService {
	return &TaskService{
		taskRepository,
		userRepository,
	}
}

func (t *TaskService) Get(ctx context.Context, id uint, user uint) (*model.Task, error) {
	if id < 0 {
		return &model.Task{}, ErrTaskIdIsEmpty
	}

	taskPtr, err := t.TaskRepository.Get(ctx, id)
	if err != nil {
		return &model.Task{}, fmt.Errorf("%s", err.Error())
	}

	return taskPtr, nil
}

func (t *TaskService) Show(ctx context.Context, paginat *paginate.Pagination, user uint) (*paginate.Pagination, error) {
	if paginat.Page < 0 {
		return &paginate.Pagination{}, fmt.Errorf("Error page is less then 0. Page: %d", paginat.Page)
	}

	newPaginat, err := t.TaskRepository.GetAll(ctx, paginat, user)
	if err != nil {
		return &paginate.Pagination{}, fmt.Errorf("%s", err.Error())
	}

	return newPaginat, nil
}

func (t *TaskService) Create(ctx context.Context, dto request.CreateTaskRequest, user uint) (uint, error) {
	if dto.Title == "" {
		return 0, ErrTaskNameIsEmpty
	}

	err := t.Validate(ctx, user, map[int16]bool{
		roles.AdminRole:       true,
		roles.OrganizatorRole: true,
	})
	if err != nil {
		return 0, err
	}

	id, err := t.TaskRepository.Create(ctx, dto)
	if err != nil {
		return 0, fmt.Errorf("%s", err.Error())
	}

	return id, nil
}

func (t *TaskService) Update(ctx context.Context, dto request.UpdateTaskRequest, user uint) error {
	if dto.ID < 0 {
		return ErrTaskIdIsEmpty
	}

	err := t.Validate(ctx, user, map[int16]bool{
		roles.AdminRole:       true,
		roles.OrganizatorRole: true,
	})
	if err != nil {
		return err
	}

	err = t.TaskRepository.Update(ctx, dto)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	return nil
}

func (t *TaskService) Delete(ctx context.Context, id uint, user uint) error {
	if id < 0 {
		return ErrTaskIdIsEmpty
	}

	err := t.Validate(ctx, user, map[int16]bool{
		roles.AdminRole:       true,
		roles.OrganizatorRole: true,
	})
	if err != nil {
		return err
	}

	err = t.TaskRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	return nil
}

func (t *TaskService) Validate(ctx context.Context, user uint, validRoles map[int16]bool) (err error) {
	auth, err := t.UserRepository.Get(ctx, user)
	if err != nil {
		return ErrUserIsNotValid
	}

	if validRoles[auth.Role] {
		return ErrUserIsNotValidRole
	}

	return nil
}
