package utils

import (
	"errors"
	"net/mail"
	"renderService/layer"
)

var ErrNameNotExist = errors.New("name cannot be empty")
var ErrEmailNotExist = errors.New("email cannot be empty")
var ErrDescNotExist = errors.New("description cannot be empty")
var ErrPasswordNotExist = errors.New("password cannot be empty")

func ValidateUser(payload *layer.User) error {
	//name is a must
	if payload.Name == "" {
		return ErrNameNotExist
	}
	//email is a must & must valid
	if payload.Email == "" {
		return ErrEmailNotExist
	}
	if _, err := mail.ParseAddress(payload.Email); err != nil {
		return err
	}
	//desc is a must
	if payload.Desc == "" {
		return ErrDescNotExist
	}
	//password is a must
	if payload.Password == "" {
		return ErrPasswordNotExist
	}
	return nil
}
