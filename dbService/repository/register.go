package repository

import (
	"context"
	"dbservice/layer"
	"dbservice/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Register struct {
	DB *gorm.DB
	*layer.Register
}

func setRegisterModel(u *Register) (model.Register, error) {
	mdl := model.Register{}
	if u.Requirement != nil {
		mdl.Requirement = *u.Requirement
	}
	if u.PaymentImg != nil {
		mdl.PaymentImg = *u.PaymentImg
	}
	if u.IDImg != nil {
		mdl.IDImg = *u.IDImg
	}
	if u.User_ID != nil {
		mdl.UserID = *u.User_ID
	}
	if u.Event_ID != nil {
		mdl.EventID = *u.Event_ID
	}
	return mdl, nil
}

func (u *Register) Create(ctx *context.Context) error {
	mdl, err := setRegisterModel(u)
	if err != nil {
		return err
	}

	// set created at now
	mdl.CreatedAt = time.Now()
	return u.DB.Create(&mdl).Error
}

func (u *Register) Get(ctx *context.Context, pagination *layer.Pagination) (*layer.RegisterList, error) {
	result := layer.RegisterList{}
	if pagination.Query != nil {
		err := u.DB.Model(&model.Register{}).Where(fmt.Sprintf("id > ? AND %s", *pagination.Query), *pagination.LastID).Limit(int(*pagination.Limit)).Find(&result.List).Error
		if err != nil {
			return &layer.RegisterList{}, err
		}
		return &result, nil
	}
	err := u.DB.Model(&model.Register{}).Where("id > ?", *pagination.LastID).Limit(int(*pagination.Limit)).Find(&result.List).Error
	if err != nil {
		return &layer.RegisterList{}, err
	}
	return &result, nil
}

func (u *Register) GetByID(ctx *context.Context, ID *uint32) (*layer.Register, error) {
	return u.getOneRegisterByField(ctx, *ID, "id")
}

func (u *Register) getOneRegisterByField(ctx *context.Context, value interface{}, what string) (*layer.Register, error) {
	err := u.DB.Where(fmt.Sprintf("%s = ?", what), value).First(&u.Register).Error
	if err != nil {
		return nil, err
	}
	return u.Register, nil
}
