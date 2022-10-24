package repository

import (
	"context"
	"dbservice/layer"
	"dbservice/model"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	DB *gorm.DB
	*layer.Event
}

func setEventModel(u *Event) model.Event {
	mdl := model.Event{}
	if u.Name != nil {
		mdl.Name = *u.Name
	}
	if u.Desc != nil {
		mdl.Desc = *u.Desc
	}
	if u.Password != nil {
		mdl.Desc = *u.Desc
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
		tmp, err := time.Parse(time.UnixDate, *u.StartDate)
		if err != nil {
			tmp = time.Now()
		}
		mdl.StartDate = tmp
	}
	if u.FinishDate != nil {
		tmp, err := time.Parse(time.UnixDate, *u.StartDate)
		if err != nil {
			tmp = time.Now()
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
	return mdl
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
	mdl := setEventModel(u)
	setEventUpdatedAtCreatedAt(&mdl)
	//set status into 1 = open for reg
	mdl.Status = uint8(1)
	return u.DB.Create(&mdl).Error
}

func (u *Event) Edit(ctx *context.Context, selectq []string, ID *uint32) error {
	mdl := setEventModel(u)
	setEventUpdatedAt(&mdl)
	return u.DB.Model(&model.Event{}).Select(selectq).Where("id = ?", *ID).Updates(&mdl).Error
}
