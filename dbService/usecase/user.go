package usecase

import (
	"context"
	"dbservice/layer"
	"dbservice/repository"

	"gorm.io/gorm"
)

type UserService struct {
	*layer.UnimplementedUserLayerServer
	DB *gorm.DB
}

func (e *UserService) SignIn(ctx context.Context, payload *layer.LoginPayload) (*layer.Token, error) {
	return nil, nil
}

func (e *UserService) GetUserID(ctx context.Context, id *layer.IDPayload) (*layer.User, error) {
	return nil, nil
}

func (e *UserService) Create(ctx context.Context, payload *layer.User) (*layer.Empty, error) {
	//create user repo
	repo := repository.User{
		DB:   e.DB,
		User: payload,
	}

	//save
	err := repo.CreateUser(ctx)
	if err != nil {
		return &layer.Empty{}, err
	}

	return &layer.Empty{}, nil
}

func (e *UserService) Edit(ctx context.Context, payload *layer.UserEditPayload) (*layer.Empty, error) {

	return &layer.Empty{}, nil
}
