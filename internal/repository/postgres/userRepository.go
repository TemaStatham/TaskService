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

func (u *UserRepository) Get(ctx context.Context, id uint) (*model.User, error) {
	var user model.User

	res := u.db.First(&user, "id = ?", id)
	if res.Error != nil {
		return nil, errors.New("user not found" + res.Error.Error())
	}

	return &user, nil
}

func (u *UserRepository) Create(ctx context.Context, user *model.User) error {
	res := u.db.Create(user)
	if res.Error != nil {
		return errors.New("user not found" + res.Error.Error())
	}

	return nil
}
