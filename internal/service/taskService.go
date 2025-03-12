package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/TemaStatham/TaskService/internal/handler/request"
	"github.com/TemaStatham/TaskService/internal/model"
	"github.com/TemaStatham/TaskService/pkg/paginate"
)

var (
	ErrTaskIdIsEmpty      = errors.New("task id is empty")
	ErrTaskNameIsEmpty    = errors.New("task name is empty")
	ErrUserIsNotValid     = errors.New("user is not found valid")
	ErrUserIsNotValidRole = errors.New("user is not found valid role")
)

type TaskRepository interface {
	Get(ctx context.Context, id uint) (*model.TaskModel, error)
	GetAll(
		ctx context.Context,
		pagination *paginate.Pagination,
		user uint,
	) (*paginate.Pagination, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, dto *request.UpdateTaskRequest) error
	Create(ctx context.Context, dto *request.CreateTaskRequest) (uint, error)
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

func (t *TaskService) Get(ctx context.Context, id uint) (*model.TaskModel, error) {
	if id < 0 {
		return &model.TaskModel{}, ErrTaskIdIsEmpty
	}

	taskPtr, err := t.TaskRepository.Get(ctx, id)
	if err != nil {
		return &model.TaskModel{}, fmt.Errorf("%s", err.Error())
	}

	return taskPtr, nil
}

func (t *TaskService) Show(
	ctx context.Context,
	pagination *paginate.Pagination,
	user uint,
) (*paginate.Pagination, error) {
	if pagination.Page < 0 {
		return &paginate.Pagination{}, fmt.Errorf("Error page is less then 0. Page: %d", pagination.Page)
	}

	newPagination, err := t.TaskRepository.GetAll(ctx, pagination, user)
	if err != nil {
		return &paginate.Pagination{}, fmt.Errorf("%s", err.Error())
	}

	return newPagination, nil
}

func (t *TaskService) Create(ctx context.Context, dto *request.CreateTaskRequest) (uint, error) {
	if dto.Name == "" {
		return 0, ErrTaskNameIsEmpty
	}

	/*err := t.Validate(ctx, user, map[int16]bool{
		roles.AdminRole:       true,
		roles.OrganizatorRole: true,
	})
	if err != nil {
		return 0, err
	}*/

	id, err := t.TaskRepository.Create(ctx, dto)
	if err != nil {
		return 0, fmt.Errorf("%s", err)
	}

	return id, nil
}

func (t *TaskService) Update(ctx context.Context, dto *request.UpdateTaskRequest) error {
	if dto.ID < 0 {
		return ErrTaskIdIsEmpty
	}

	/*err := t.Validate(ctx, user, map[int16]bool{
		roles.AdminRole:       true,
		roles.OrganizatorRole: true,
	})
	if err != nil {
		return err
	}*/

	err := t.TaskRepository.Update(ctx, dto)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

func (t *TaskService) Delete(ctx context.Context, id uint) error {
	if id < 0 {
		return ErrTaskIdIsEmpty
	}

	/*err := t.Validate(ctx, user, map[int16]bool{
		roles.AdminRole:       true,
		roles.OrganizatorRole: true,
	})
	if err != nil {
		return err
	}*/

	err := t.TaskRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}
