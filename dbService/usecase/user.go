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

func (e *UserService) GetByEmail(ctx context.Context, payload *layer.EmailPayload) (*layer.User, error) {
	//create user repo
	repo := repository.User{
		DB: e.DB,
	}

	//find user by email
	return repo.GetByEmail(&ctx, &payload.Email)
}

func (e *UserService) GetByID(ctx context.Context, id *layer.IDPayload) (*layer.User, error) {
	//create user repo
	repo := repository.User{
		DB: e.DB,
	}

	return repo.GetByID(&ctx, &id.ID)
}

func (e *UserService) Create(ctx context.Context, payload *layer.User) (*layer.Empty, error) {
	//create user repo
	repo := repository.User{
		DB:   e.DB,
		User: payload,
	}

	//create
	err := repo.Create(&ctx)
	if err != nil {
		return &layer.Empty{}, err
	}

	return &layer.Empty{}, nil
}

func (e *UserService) Edit(ctx context.Context, payload *layer.UserEditPayload) (*layer.Empty, error) {
	//create user repo
	repo := repository.User{
		DB:   e.DB,
		User: payload.User,
	}

	//edit
	err := repo.Edit(&ctx, payload.Select, &payload.ID)
	if err != nil {
		return &layer.Empty{}, err
	}

	return &layer.Empty{}, nil
}
