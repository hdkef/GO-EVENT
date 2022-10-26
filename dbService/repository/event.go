package repository

import (
	"context"
	"dbservice/layer"
	"dbservice/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	DB *gorm.DB
	*layer.Event
}

func setEventModel(u *Event) (model.Event, error) {
	mdl := model.Event{}
	if u.Name != nil {
		mdl.Name = *u.Name
	}
	if u.Description != nil {
		mdl.Description = *u.Description
	}
	if u.Password != nil {
		mdl.Password = *u.Password
	}
	if u.Requirement != nil {
		mdl.Requirement = *u.Requirement
	}
	if u.NeedPayment != nil {
		mdl.NeedPayment = *u.NeedPayment
	}
	if u.Need_ID != nil {
		mdl.NeedID = *u.Need_ID
	}
	if u.Payment_ID != nil {
		mdl.PaymentID = *u.Payment_ID
	}
	if u.Creator_ID != nil {
		mdl.CreatorID = *u.Creator_ID
	}
	if u.PaymentPrice != nil {
		mdl.PaymentPrice = *u.PaymentPrice
	}
	if u.EventCategory != nil {
		mdl.EventCategory = uint8(*u.EventCategory)
	}
	if u.IsOffline != nil {
		mdl.IsOffline = *u.IsOffline
	}
	if u.LocationAddress != nil {
		mdl.LocationAddress = *u.LocationAddress
	}
	if u.LocationCity != nil {
		mdl.LocationCity = *u.LocationCity
	}
	if u.LocationProvince != nil {
		mdl.LocationProvince = *u.LocationProvince
	}
	if u.StartDate != nil {
		tmp, err := time.Parse("2006-01-02T15:04:05Z07:00", *u.StartDate)
		if err != nil {
			panic(err.Error())
		}
		mdl.StartDate = tmp
	}
	if u.FinishDate != nil {
		tmp, err := time.Parse("2006-01-02T15:04:05Z07:00", *u.FinishDate)
		if err != nil {
			return model.Event{}, err
		}
		mdl.FinishDate = tmp
	}
	if u.Status != nil {
		mdl.Status = uint8(*u.Status)
	}
	if u.PresenceQuestion != nil {
		mdl.PresenceQuestion = *u.PresenceQuestion
	}
	if u.MediaLink != nil {
		mdl.MediaLink = *u.MediaLink
	}
	return mdl, nil
}

func setEventUpdatedAtCreatedAt(u *model.Event) {
	setEventCreatedAt(u)
	setEventUpdatedAt(u)
}

func setEventUpdatedAt(u *model.Event) {
	u.UpdatedAt = time.Now()
}

func setEventCreatedAt(u *model.Event) {
	u.UpdatedAt = time.Now()
}

func (u *Event) Create(ctx *context.Context) error {
	mdl, err := setEventModel(u)
	if err != nil {
		return err
	}
	setEventUpdatedAtCreatedAt(&mdl)
	// set status into 1 = open for reg
	mdl.Status = 1
	return u.DB.Create(&mdl).Error
}

func (u *Event) Edit(ctx *context.Context, selectq []string, ID *uint32) error {
	mdl, err := setEventModel(u)
	if err != nil {
		return err
	}
	setEventUpdatedAt(&mdl)
	return u.DB.Model(&model.Event{}).Select(selectq).Where("id = ?", *ID).Updates(&mdl).Error
}

func (u *Event) Get(ctx *context.Context, pagination *layer.Pagination) (*layer.EventList, error) {

	result := layer.EventList{}
	if pagination.Query != nil {
		err := u.DB.Model(&model.Event{}).Where(fmt.Sprintf("id > ? AND %s", *pagination.Query), *pagination.LastID).Limit(int(*pagination.Limit)).Find(&result.List).Error
		if err != nil {
			return &layer.EventList{}, err
		}
		return &result, nil
	}
	err := u.DB.Model(&model.Event{}).Where("id > ?", *pagination.LastID).Limit(int(*pagination.Limit)).Find(&result.List).Error
	if err != nil {
		return &layer.EventList{}, err
	}
	return &result, nil
}

func (u *Event) GetByID(ctx *context.Context, ID *uint32) (*layer.Event, error) {
	return u.getOneEventByField(ctx, *ID, "id")
}

func (u *Event) getOneEventByField(ctx *context.Context, value interface{}, what string) (*layer.Event, error) {
	err := u.DB.Where(fmt.Sprintf("%s = ?", what), value).First(&u.Event).Error
	if err != nil {
		return nil, err
	}
	return u.Event, nil
}

func (u *Event) Delete(ctx *context.Context, id *uint32) error {
	return u.DB.Delete(&model.Event{}, *id).Error
}
