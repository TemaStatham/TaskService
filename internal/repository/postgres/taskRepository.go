package postgres

import (
	"context"
	"errors"
	"github.com/TemaStatham/TaskService/internal/handler/request"
	"github.com/TemaStatham/TaskService/internal/model"
	"github.com/TemaStatham/TaskService/pkg/paginate"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskPostgresRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (t *TaskRepository) Create(ctx context.Context, dto *request.CreateTaskRequest) (uint, error) {
	task := &model.TaskModel{
		OrganizationID:    dto.Organization.ID,
		Name:              dto.Name,
		TypeID:            dto.TaskType.ID,
		Description:       dto.Description,
		Location:          dto.Location,
		TaskDate:          dto.TaskDate,
		ParticipantsCount: dto.ParticipantsCount,
		MaxScore:          dto.MaxScore,
		StatusID:          dto.TaskStatus.ID,
	}

	res := t.db.Create(&task)
	if res.Error != nil {
		return 0, res.Error
	}

	return task.ID, nil
}

func (t *TaskRepository) Update(ctx context.Context, dto *request.UpdateTaskRequest) error {
	var task model.TaskModel

	res := t.db.First(&task, "id = ?", dto.ID)
	if res.Error != nil {
		return errors.New("user not found" + res.Error.Error())
	}

	res = t.db.Model(&task).Updates(dto)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (t *TaskRepository) Delete(ctx context.Context, id uint) error {
	res := t.db.Delete(&model.TaskModel{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (t *TaskRepository) Get(ctx context.Context, id uint) (*model.TaskModel, error) {
	var task model.TaskModel

	res := t.db.First(&task, "id = ?", id)
	if res.Error != nil {
		return nil, errors.New("user not found" + res.Error.Error())
	}

	return &task, nil
}

func (t *TaskRepository) GetAll(
	ctx context.Context,
	pagination *paginate.Pagination,
	user uint,
) (*paginate.Pagination, error) {
	var tasks []*model.TaskModel

	// todo : добавить фильтр, что показывать нужно только общие все и закрытые задания
	res := t.db.Scopes(paginate.Paginate(tasks, pagination, t.db)).Find(&tasks)
	if res.Error != nil {
		return nil, res.Error
	}

	pagination.Rows = tasks

	return pagination, nil
}
