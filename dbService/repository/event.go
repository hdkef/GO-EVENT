package repository

import (
	"context"
	"dbservice/layer"
	"dbservice/model"

	"gorm.io/gorm"
)

type Event struct {
	DB *gorm.DB
	*layer.Event
}

func (u *Event) Create(ctx *context.Context) error {
	mdl := model.Event{}

	return u.DB.Create(&mdl).Error
}
