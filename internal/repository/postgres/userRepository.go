package postgres

import (
	"context"
	"errors"
	"github.com/TemaStatham/TaskService/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Get(ctx context.Context, id uint) (*model.UserModel, error) {
	var user model.UserModel

	res := u.db.First(&user, "id = ?", id)
	if res.Error != nil {
		return nil, errors.New("user not found" + res.Error.Error())
	}

	return &user, nil
}

func (u *UserRepository) IsCoordinatorForTask(ctx context.Context, userID, taskID uint) (bool, error) {
	userTask := model.TaskUser{}
	res := u.db.First(&userTask, "user_id = ? AND task_id = ?", userID, taskID)
	if res.Error != nil {
		return false, errors.New("user not found" + res.Error.Error())
	}

	return userTask.IsCoordinator, nil
}

func (u *UserRepository) IsAdmin(ctx context.Context, userID uint) (bool, error) {
	user := model.UserModel{}
	res := u.db.First(&user, "id = ?", userID)
	if res.Error != nil {
		return false, errors.New("user not found" + res.Error.Error())
	}

	return user.IsAdmin, nil
}

func (u *UserRepository) IsOwner(ctx context.Context, userID uint) (bool, error) {
	// todo: непонятно как отличить владельца организации от других ролей
	return false, nil
}

func (u *UserRepository) Create(ctx context.Context, user *model.UserModel) error {
	res := u.db.Save(user)
	if res.Error != nil {
		res = u.db.Updates(user)
		if res.Error != nil {
			return errors.New("user not found" + res.Error.Error())
		}
	}

	return nil
}
