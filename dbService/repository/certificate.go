package repository

import (
	"context"
	"dbservice/layer"
	"dbservice/model"
	"fmt"

	"gorm.io/gorm"
)

type Certificate struct {
	DB *gorm.DB
	*layer.Certificate
}

func setCertificateModel(u *Certificate) (model.Certificate, error) {
	mdl := model.Certificate{}
	if u.User_ID != nil {
		mdl.UserID = *u.User_ID
	}
	if u.Event_ID != nil {
		mdl.EventID = *u.Event_ID
	}
	if u.FileUrl != nil {
		mdl.FileURL = *u.FileUrl
	}
	return mdl, nil
}

func (u *Certificate) Create(ctx *context.Context) error {
	mdl, err := setCertificateModel(u)
	if err != nil {
		return err
	}

	return u.DB.Create(&mdl).Error
}

func (u *Certificate) Get(ctx *context.Context, pagination *layer.Pagination) (*layer.CertificateList, error) {

	result := layer.CertificateList{}
	if pagination.Query != nil {
		err := u.DB.Model(&model.Certificate{}).Where(fmt.Sprintf("id > ? AND %s", *pagination.Query), *pagination.LastID).Limit(int(*pagination.Limit)).Find(&result.List).Error
		if err != nil {
			return &layer.CertificateList{}, err
		}
		return &result, nil
	}
	err := u.DB.Model(&model.Certificate{}).Where("id > ?", *pagination.LastID).Limit(int(*pagination.Limit)).Find(&result.List).Error
	if err != nil {
		return &layer.CertificateList{}, err
	}
	return &result, nil
}
