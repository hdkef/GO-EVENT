package repository

import (
	"context"
	"dbservice/layer"
	"dbservice/model"
	"fmt"

	"gorm.io/gorm"
)

type Subscription struct {
	DB *gorm.DB
	*layer.Subscription
}

func setSubscriptionModel(u *Subscription) (model.Subscription, error) {
	mdl := model.Subscription{}
	if u.Consumer_ID != nil {
		mdl.ConsumerID = *u.Consumer_ID
	}
	if u.Publisher_ID != nil {
		mdl.PublisherID = *u.Publisher_ID
	}
	return mdl, nil
}

func (u *Subscription) Create(ctx *context.Context) error {
	mdl, err := setSubscriptionModel(u)
	if err != nil {
		return err
	}

	return u.DB.Create(&mdl).Error
}

func (u *Subscription) GetByID(ctx *context.Context, id *uint32) (*layer.Subscription, error) {

	return u.getOneSubscriptionByField(ctx, *id, "id")
}

func (u *Subscription) Get(ctx *context.Context, pagination *layer.Pagination) (*layer.SubscriptionList, error) {
	result := layer.SubscriptionList{}
	if pagination.Query != nil {
		err := u.DB.Model(&model.Subscription{}).Where(fmt.Sprintf("id > ? AND %s", *pagination.Query), *pagination.LastID).Limit(int(*pagination.Limit)).Find(&result.List).Error
		if err != nil {
			return &layer.SubscriptionList{}, err
		}
		return &result, nil
	}
	err := u.DB.Model(&model.Subscription{}).Where("id > ?", *pagination.LastID).Limit(int(*pagination.Limit)).Find(&result.List).Error
	if err != nil {
		return &layer.SubscriptionList{}, err
	}
	return &result, nil
}

func (u *Subscription) Delete(ctx *context.Context, id *uint32) error {
	return u.DB.Delete(&model.Subscription{}, *id).Error
}

func (u *Subscription) getOneSubscriptionByField(ctx *context.Context, value interface{}, what string) (*layer.Subscription, error) {
	err := u.DB.Where(fmt.Sprintf("%s = ?", what), value).First(&u.Subscription).Error
	if err != nil {
		return nil, err
	}
	return u.Subscription, nil
}
