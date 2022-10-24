package utils

import (
	"fmt"
	"html"
	"net/mail"
	"renderService/layer"
)

const VALIDATE_TYPE_CREATE = uint8(1)
const VALIDATE_TYPE_UPDATE = uint8(2)

func errEmptyString(name string) error {
	return fmt.Errorf("%s is empty", name)
}

func escapeString(argv ...*string) {
	for _, v := range argv {
		if v != nil {
			*v = html.EscapeString(*v)
		}
	}
}

func mustNotEmptyString(name string, value *string) error {
	if value == nil {
		return errEmptyString(name)
	}
	return nil
}

func ValidateUser(payload *layer.User, validateType uint8) error {
	//escape string XSS
	escapeString(payload.Name, payload.Desc, payload.Email, payload.Password)

	var err error
	switch validateType {
	case VALIDATE_TYPE_CREATE:
		//name is a must
		err = mustNotEmptyString("name", payload.Name)
		//email is a must & must valid
		err = mustNotEmptyString("email", payload.Email)
		_, err = mail.ParseAddress(*payload.Email)
		//desc is a must
		err = mustNotEmptyString("description", payload.Desc)
		//password is a must
		err = mustNotEmptyString("password", payload.Password)
		return err
	case VALIDATE_TYPE_UPDATE:
		//email must be right format
		if payload.Email != nil {
			_, err = mail.ParseAddress(*payload.Email)
		}
		return err
	}
	return err
}
