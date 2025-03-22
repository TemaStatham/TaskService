package query

import (
	"context"
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/user/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type UserQueryInterface interface {
	GetUser(ctx context.Context, userID uint64) (model.User, error)
	GetUsersByIDS(ctx context.Context, userIDS []uint64) ([]model.User, error)
}

// Grpcs методы
type ClientUserInterface interface {
	GetUser(ctx context.Context, userID uint64) (model.User, error)
}

type UserQuery struct {
	client ClientUserInterface
}

func (u *UserQuery) GetUsersByIDS(ctx context.Context, userIDS []uint64) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserQuery(client ClientUserInterface) UserQueryInterface {
	return &UserQuery{
		client: client,
	}
}

func (u *UserQuery) GetUser(ctx context.Context, userID uint64) (model.User, error) {
	if userID <= 0 {
		return model.User{}, grpc.Errorf(codes.InvalidArgument, "invalid user id")
	}

	return u.client.GetUser(ctx, userID)
}
