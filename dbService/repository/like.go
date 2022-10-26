package repository

import (
	"context"
	"dbservice/layer"
	"dbservice/model"
	"fmt"

	"gorm.io/gorm"
)

type Like struct {
	DB *gorm.DB
	*layer.Like
}

func setLikeModel(u *Like) (model.Like, error) {
	mdl := model.Like{}
	if u.User_ID != nil {
		mdl.UserID = *u.User_ID
	}
	if u.Event_ID != nil {
		mdl.EventID = *u.Event_ID
	}
	return mdl, nil
}

func (u *Like) Create(ctx *context.Context) error {
	mdl, err := setLikeModel(u)
	if err != nil {
		return err
	}

	return u.DB.Create(&mdl).Error
}

func (u *Like) Get(ctx *context.Context, pagination *layer.Pagination) (*layer.LikeList, error) {
	result := layer.LikeList{}
	if pagination.Query != nil {
		err := u.DB.Model(&model.Like{}).Where(fmt.Sprintf("id > ? AND %s", *pagination.Query), *pagination.LastID).Limit(int(*pagination.Limit)).Find(&result.List).Error
		if err != nil {
			return &layer.LikeList{}, err
		}
		return &result, nil
	}
	err := u.DB.Model(&model.Like{}).Where("id > ?", *pagination.LastID).Limit(int(*pagination.Limit)).Find(&result.List).Error
	if err != nil {
		return &layer.LikeList{}, err
	}
	return &result, nil
}

func (u *Like) Delete(ctx *context.Context, id *uint32) error {
	return u.DB.Delete(&model.Like{}, *id).Error
}

func (u *Like) GetByID(ctx *context.Context, id *uint32) (*layer.Like, error) {
	return u.getOneLikeByField(ctx, *id, "id")
}

func (u *Like) getOneLikeByField(ctx *context.Context, value interface{}, what string) (*layer.Like, error) {
	err := u.DB.Where(fmt.Sprintf("%s = ?", what), value).First(&u.Like).Error
	if err != nil {
		return nil, err
	}
	return u.Like, nil
}
