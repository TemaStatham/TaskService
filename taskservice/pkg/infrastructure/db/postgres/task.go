package postgres

import (
	"context"
	"errors"
	"fmt"
	organizationmodel "github.com/TemaStatham/TaskService/taskservice/pkg/app/organization/model"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/task/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/task/model"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/lib/paginate"
	"gorm.io/gorm"
)

var (
	ErrTaskNotFound   = errors.New("Задача не найдена")
	ErrUpdateTask     = errors.New("Не удалось обновить задачу")
	ErrDeleteTask     = errors.New("Не удалось удалить задачу")
	ErrCreateTask     = errors.New("Ошибка при создании задачи")
	ErrAddCoordinator = errors.New("Ошибка при добавлении координатора")
	ErrAddCategory    = errors.New("Ошибка при добавлении категории")
)

type TaskRepository struct {
	db               *gorm.DB
	taskuserrepo     model.TaskUserRepositoryInterface
	taskcategoryrepo model.TaskCategoryRepositoryInterface
}

func NewTaskPostgresRepository(
	db *gorm.DB,
	taskuserrepo model.TaskUserRepositoryInterface,
	taskcategoryrepo model.TaskCategoryRepositoryInterface,
) *TaskRepository {
	return &TaskRepository{
		db:               db,
		taskuserrepo:     taskuserrepo,
		taskcategoryrepo: taskcategoryrepo,
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
		return 0, fmt.Errorf("%w: %v", ErrCreateTask, res.Error)
	}

	for _, coordinator := range task.Coordinators {
		if err := t.taskuserrepo.Add(ctx, coordinator, taskModel.ID, true); err != nil {
			fmt.Printf("%s: %d\n", ErrAddCoordinator, coordinator)
		}
	}

	for _, category := range task.Categories {
		if err := t.taskcategoryrepo.Create(ctx, taskModel.ID, category); err != nil {
			fmt.Printf("%s: %d\n", ErrAddCategory, category)
		}
	}

	return taskModel.ID, nil
}

func (t *TaskRepository) Update(ctx context.Context, task *data.UpdateTask) error {
	res := t.db.WithContext(ctx).Model(&model.TaskModel{}).Where("id = ?", task.ID).Updates(task)
	if res.Error != nil {
		return fmt.Errorf("%w: %v", ErrUpdateTask, res.Error)
	}

	return nil
}

func (t *TaskRepository) Delete(ctx context.Context, id uint) error {
	res := t.db.WithContext(ctx).Delete(&model.TaskModel{}, id)
	if res.Error != nil {
		return fmt.Errorf("%w: %v", ErrDeleteTask, res.Error)
	}
	return nil
}

func (t *TaskRepository) Get(ctx context.Context, id uint) (*model.TaskModel, error) {
	var task model.TaskModel
	res := t.db.WithContext(ctx).First(&task, "id = ?", id)
	if res.Error != nil {
		return nil, fmt.Errorf("%w: %v", ErrTaskNotFound, res.Error)
	}
	return &task, nil
}

func (t *TaskRepository) GetAll(
	ctx context.Context,
	dto data.GetAllTasks,
	user uint,
	organizations []organizationmodel.Organization,
) (*paginate.Pagination, error) {
	var tasks []*model.TaskModel

	// todo вынести в сервис
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

	var total int64
	if err := query.Model(&model.TaskModel{}).Count(&total).Error; err != nil {
		return nil, fmt.Errorf("Ошибка при подсчете задач: %v", err)
	}

	// todo: вынести в сервис
	limit := dto.Limit
	if limit <= paginate.MinLimit {
		limit = paginate.DefaultLimit
	}

	page := dto.Page
	if page <= paginate.MinPage {
		page = paginate.DefaultPage
	}

	offset := (page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	if err := query.Find(&tasks).Error; err != nil {
		return nil, fmt.Errorf("Ошибка при получении списка задач: %v", err)
	}

	return &paginate.Pagination{
		TotalPages: total,
		Page:       page,
		Limit:      limit,
		Rows:       tasks,
	}, nil
}
