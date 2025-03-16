package postgres

import (
	"context"
	"errors"
	organizationmodel "github.com/TemaStatham/TaskService/client/pkg/app/organization/model"
	"github.com/TemaStatham/TaskService/client/pkg/app/paginate"
	"github.com/TemaStatham/TaskService/client/pkg/app/task/data"
	"github.com/TemaStatham/TaskService/client/pkg/app/task/model"
	"gorm.io/gorm"
)

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrUserNotFound = errors.New("user not found")
	ErrCreateTask   = errors.New("failed to create task")
	ErrUpdateTask   = errors.New("failed to update task")
	ErrDeleteTask   = errors.New("failed to delete task")
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskPostgresRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (t *TaskRepository) Create(ctx context.Context, task *data.CreateTask) (uint, error) {
	taskModel := model.TaskModel{
		OrganizationID:    task.Organization,
		Name:              task.Name,
		TypeID:            task.TaskType,
		Description:       task.Description,
		Location:          task.Location,
		TaskDate:          task.TaskDate,
		ParticipantsCount: task.ParticipantsCount,
		MaxScore:          task.MaxScore,
		StatusID:          task.TaskStatus,
	}
	res := t.db.WithContext(ctx).Create(&taskModel)
	if res.Error != nil {
		return 0, res.Error
	}

	return taskModel.ID, nil
}

func (t *TaskRepository) Update(ctx context.Context, task *data.UpdateTask) error {
	taskModel := model.TaskModel{
		ID:                task.ID,
		OrganizationID:    task.Organization,
		Name:              task.Name,
		TypeID:            task.TaskType,
		Description:       task.Description,
		Location:          task.Location,
		TaskDate:          task.TaskDate,
		ParticipantsCount: task.ParticipantsCount,
		MaxScore:          task.MaxScore,
		StatusID:          task.TaskStatus,
	}

	res := t.db.WithContext(ctx).Model(&task).Updates(taskModel)
	if res.Error != nil {
		return ErrUpdateTask
	}

	return nil
}

func (t *TaskRepository) Delete(ctx context.Context, id uint) error {
	res := t.db.WithContext(ctx).Delete(&model.TaskModel{}, id)
	if res.Error != nil {
		return ErrDeleteTask
	}

	return nil
}

func (t *TaskRepository) Get(ctx context.Context, id uint) (*model.TaskModel, error) {
	var task model.TaskModel

	res := t.db.WithContext(ctx).First(&task, "id = ?", id)
	if res.Error != nil {
		return nil, ErrTaskNotFound
	}

	return &task, nil
}

func (t *TaskRepository) GetAll(
	ctx context.Context,
	pagination *paginate.Pagination,
	user uint,
	organizations []organizationmodel.OrganizationModel,
) (*paginate.Pagination, error) {
	var tasks []*model.TaskModel

	// todo: вынести в сервис
	var orgIDs []uint
	for _, org := range organizations {
		orgIDs = append(orgIDs, org.ID)
	}

	query := t.db.WithContext(ctx).
		Joins("JOIN task_type tt on tt.id = task.type_id").
		Joins("JOIN task_user tu ON tu.task_id = task.id").
		Where("tu.user_id = ?", user)

	if len(orgIDs) > 0 {
		query = query.Where("tt.name = 'Открытый' OR (tt.name = 'Закрытый' AND task.organization_id IN ?)", orgIDs)
	} else {
		query = query.Where("tt.name = 'Открытый'")
	}

	query = query.Scopes(paginate.Paginate(tasks, pagination, t.db))

	res := query.Find(&tasks)
	if res.Error != nil {
		return nil, res.Error
	}

	pagination.Rows = tasks

	return pagination, nil
}
