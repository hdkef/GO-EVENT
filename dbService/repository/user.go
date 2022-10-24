package repository

import (
	"context"
	"dbservice/layer"
	"dbservice/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
	*layer.User
}

func setUserModel(u *User) model.User {
	mdl := model.User{}
	if u.ID != nil {
		mdl.ID = *u.ID
	}
	if u.Name != nil {
		mdl.Name = *u.Name
	}
	if u.Desc != nil {
		mdl.Desc = *u.Desc
	}
	if u.ProfileImg != nil {
		mdl.ProfileImg = *u.ProfileImg
	}
	if u.Email != nil {
		mdl.Email = *u.Email
	}
	if u.Password != nil {
		mdl.Password = *u.Password
	}
	return mdl
}

func setUpdatedAtCreatedAt(u *model.User) {
	setCreatedAt(u)
	setUpdatedAt(u)
}

func setUpdatedAt(u *model.User) {
	u.UpdatedAt = time.Now()
}

func setCreatedAt(u *model.User) {
	u.UpdatedAt = time.Now()
}

func (u *User) Create(ctx *context.Context) error {
	mdl := setUserModel(u)
	setUpdatedAtCreatedAt(&mdl)
	return u.DB.Create(&mdl).Error
}

func (u *User) Edit(ctx *context.Context, selectq []string, ID *uint32) error {
	mdl := setUserModel(u)
	setUpdatedAt(&mdl)
	return u.DB.Model(&model.User{}).Select(selectq).Where("id = ?", *ID).Updates(&mdl).Error
}

func (u *User) GetByEmail(ctx *context.Context, email *string) (*layer.User, error) {
	return u.getOneByField(ctx, *email, "email")
}

func (u *User) GetByID(ctx *context.Context, ID *uint32) (*layer.User, error) {
	return u.getOneByField(ctx, *ID, "id")
}

func (u *User) getOneByField(ctx *context.Context, value interface{}, what string) (*layer.User, error) {
	err := u.DB.Where(fmt.Sprintf("%s = ?", what), value).First(&u.User).Error
	if err != nil {
		return nil, err
	}
	return u.User, nil
}
