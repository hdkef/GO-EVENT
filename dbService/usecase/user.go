package usecase

import (
	"context"
	"dbservice/layer"
	"dbservice/repository"
	"dbservice/utils"
)

type UserService struct {
	*layer.UnimplementedUserLayerServer
}

func (e *UserService) SignIn(ctx context.Context, payload *layer.LoginPayload) (*layer.Token, error) {
	return nil, nil
}

func (e *UserService) GetUserID(ctx context.Context, id *layer.IDPayload) (*layer.User, error) {
	return nil, nil
}

func (e *UserService) Create(ctx context.Context, payload *layer.User) (*layer.Empty, error) {
	//validate payload
	err := utils.ValidateUser(payload)
	if err != nil {
		return &layer.Empty{}, err
	}

	//create user repo
	repo := repository.User{
		User: payload,
	}

	//save
	err = repo.CreateUser(ctx)
	if err != nil {
		return &layer.Empty{}, err
	}

	return &layer.Empty{}, nil
}

func (e *UserService) Edit(ctx context.Context, payload *layer.UserEditPayload) (*layer.Empty, error) {

	return &layer.Empty{}, nil
}
