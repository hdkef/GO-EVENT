package repository

import (
	"context"
	"dbservice/layer"
	"dbservice/model"
	"time"

	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
	*layer.User
}

func (u *User) CreateUser(ctx context.Context) error {
	mdl := model.User{
		ID:         u.GetID(),
		Name:       u.GetName(),
		Desc:       u.GetDesc(),
		ProfileImg: u.GetProfileImg(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Email:      u.GetEmail(),
		Password:   u.GetPassword(),
	}
	return u.DB.Create(&mdl).Error
}
