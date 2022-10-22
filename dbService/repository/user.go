package repository

import (
	"context"
	"dbservice/layer"
	"dbservice/model"

	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
	*layer.User
}

func (u *User) CreateUser(ctx context.Context) error {
	mdl := model.User{
		ID: *u.ID,
	}
	return u.DB.Create(&mdl).Error
}
