package usecase

import "context"

type UserService struct {
	*UnimplementedUserLayerServer
}

func (e *UserService) Login(ctx context.Context, payload *LoginPayload) (*Token, error) {
	return nil, nil
}

func (e *UserService) GetUserByID(ctx context.Context, id *IDPayload) (*User, error) {
	return nil, nil
}

func (e *UserService) Register(ctx context.Context, user *User) (*Empty, error) {
	return nil, nil
}
