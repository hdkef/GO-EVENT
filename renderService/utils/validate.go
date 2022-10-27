package utils

import (
	"fmt"
	"html"
	"net/mail"
	"renderService/layer"
)

const VALIDATE_TYPE_CREATE = uint8(1)
const VALIDATE_TYPE_UPDATE = uint8(2)

func errEmpty(name string) error {
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
		return errEmpty(name)
	}
	return nil
}

func mustNotEmptyUint(name string, value *uint32) error {
	if value == nil {
		return errEmpty(name)
	}
	return nil
}

func mustNotEmptyBool(name string, value *bool) error {
	if value == nil {
		return errEmpty(name)
	}
	return nil
}

func ValidateUser(payload *layer.User, validateType uint8) error {
	//escape string XSS
	escapeString(payload.Name, payload.Description, payload.Email, payload.Password)

	var err error
	switch validateType {
	case VALIDATE_TYPE_CREATE:
		//name is a must
		err = mustNotEmptyString("name", payload.Name)
		if err != nil {
			return err
		}
		//email is a must & must valid
		err = mustNotEmptyString("email", payload.Email)
		if err != nil {
			return err
		}
		_, err = mail.ParseAddress(*payload.Email)
		if err != nil {
			return err
		}
		//desc is a must
		err = mustNotEmptyString("description", payload.Description)
		if err != nil {
			return err
		}
		//password is a must
		err = mustNotEmptyString("password", payload.Password)
		if err != nil {
			return err
		}
	case VALIDATE_TYPE_UPDATE:
		//email must be right format
		if payload.Email != nil {
			_, err = mail.ParseAddress(*payload.Email)
		}
		if err != nil {
			return err
		}
	}
	return err
}

func ValidateEvent(payload *layer.Event, validateType uint8) error {
	//escape string XSS
	escapeString(
		payload.Name,
		payload.Description,
		payload.EventImg,
		payload.Requirement,
		payload.LocationAddress,
		payload.LocationCity,
		payload.LocationProvince,
		payload.PresenceQuestion,
		payload.MediaLink,
	)

	var err error
	switch validateType {
	case VALIDATE_TYPE_CREATE:
		//name is a must
		err = mustNotEmptyString("name", payload.Name)
		if err != nil {
			return err
		}
		//desc is a must
		err = mustNotEmptyString("description", payload.Description)
		if err != nil {
			return err
		}
		//requirement is a must
		err = mustNotEmptyString("requirement", payload.Requirement)
		if err != nil {
			return err
		}
		//needPayment is a must
		err = mustNotEmptyBool("need payment", payload.NeedPayment)
		if err != nil {
			return err
		}
		//needID is a must
		err = mustNotEmptyBool("need id", payload.Need_ID)
		if err != nil {
			return err
		}
		//isOffline is a must
		err = mustNotEmptyBool("is offline", payload.IsOffline)
		if err != nil {
			return err
		}
		//publisher id is a must
		err = mustNotEmptyUint("publisher id", payload.Publisher_ID)
		if err != nil {
			return err
		}
		//start date is a must
		err = mustNotEmptyString("start date", payload.StartDate)
		if err != nil {
			return err
		}
		//finish date is a must
		err = mustNotEmptyString("finish date", payload.FinishDate)
		if err != nil {
			return err
		}
		//presence question is a must
		err = mustNotEmptyString("presence question", payload.PresenceQuestion)
		if err != nil {
			return err
		}
		//media link is a must
		err = mustNotEmptyString("media link", payload.MediaLink)
		//event category is a must
		if payload.EventCategory == nil {
			err = errEmpty("event category")
		}
		if err != nil {
			return err
		}
	case VALIDATE_TYPE_UPDATE:
		return err
	}

	return err
}

func ValidateRegister(payload *layer.Register, validateType uint8) error {
	escapeString(payload.Requirement, payload.IDImg, payload.PaymentImg)
	var err error

	switch validateType {
	case VALIDATE_TYPE_CREATE:
		err = mustNotEmptyString("requirement", payload.Requirement)
		if err != nil {
			return err
		}
		err = mustNotEmptyUint("event id", payload.Event_ID)
		if err != nil {
			return err
		}
		err = mustNotEmptyUint("use id", payload.User_ID)
		if err != nil {
			return err
		}
	}
	return err
}

func ValidateCertificate(payload *layer.Certificate, validateType uint8) error {
	var err error
	switch validateType {
	case VALIDATE_TYPE_CREATE:
		err = mustNotEmptyUint("user id", payload.User_ID)
		if err != nil {
			return err
		}
		err = mustNotEmptyUint("event id", payload.Event_ID)
		if err != nil {
			return err
		}
		err = mustNotEmptyString("file url", payload.FileUrl)
		if err != nil {
			return err
		}
	}
	return nil
}
