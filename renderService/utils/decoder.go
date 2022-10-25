package utils

import (
	"renderService/layer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DecodeUser(c *gin.Context, user *layer.User) ([]string, error) {
	//variable to store select query
	selectQ := []string{}
	err := c.Request.ParseForm()
	if err != nil {
		return nil, err
	}

	form := c.Request.Form

	key, exist := form["name"]
	if exist {
		user.Name = &key[0]
		selectQ = append(selectQ, "name")
	}

	key, exist = form["desc"]
	if exist {
		user.Desc = &key[0]
		selectQ = append(selectQ, "desc")
	}

	key, exist = form["email"]
	if exist {
		user.Email = &key[0]
		selectQ = append(selectQ, "email")
	}

	key, exist = form["password"]
	if exist {
		user.Password = &key[0]
		selectQ = append(selectQ, "password")
	}

	return selectQ, nil
}

func DecodeEvent(c *gin.Context, event *layer.Event) ([]string, error) {
	//variable to store select query
	selectQ := []string{}
	err := c.Request.ParseForm()
	if err != nil {
		return nil, err
	}

	form := c.Request.Form

	key, exist := form["name"]
	if exist {
		event.Name = &key[0]
		selectQ = append(selectQ, "name")
	}

	key, exist = form["desc"]
	if exist {
		event.Desc = &key[0]
		selectQ = append(selectQ, "desc")
	}

	key, exist = form["event_img"]
	if exist {
		event.EventImg = &key[0]
		selectQ = append(selectQ, "event_img")
	}

	key, exist = form["password"]
	if exist {
		event.Password = &key[0]
		selectQ = append(selectQ, "password")
	}

	key, exist = form["requirement"]
	if exist {
		event.Requirement = &key[0]
		selectQ = append(selectQ, "requirement")
	}

	key, exist = form["need_payment"]
	if exist {
		b, _ := strconv.ParseBool(key[0])
		event.NeedPayment = &b
		selectQ = append(selectQ, "need_payment")
	}

	key, exist = form["need_id"]
	if exist {
		b, _ := strconv.ParseBool(key[0])
		event.Need_ID = &b
		selectQ = append(selectQ, "need_id")
	}

	key, exist = form["payment_id"]
	if exist {
		b, _ := strconv.ParseUint(key[0], 10, 32)
		a := uint32(b)
		event.Payment_ID = &a
		selectQ = append(selectQ, "payment_id")
	}

	key, exist = form["creator_id"]
	if exist {
		b, _ := strconv.ParseUint(key[0], 10, 32)
		c := uint32(b)
		event.Creator_ID = &c
		selectQ = append(selectQ, "creator_id")
	}

	key, exist = form["payment_price"]
	if exist {
		b, _ := strconv.ParseFloat(key[0], 64)
		event.PaymentPrice = &b
		selectQ = append(selectQ, "payment_price")
	}

	key, exist = form["event_category"]
	if exist {
		b, _ := strconv.ParseInt(key[0], 10, 32)
		c := layer.EventCategory(b)
		event.EventCategory = &c
		selectQ = append(selectQ, "event_category")
	}

	key, exist = form["is_offline"]
	if exist {
		b, _ := strconv.ParseBool(key[0])
		event.IsOffline = &b
		selectQ = append(selectQ, "is_offline")
	}

	key, exist = form["location_address"]
	if exist {
		event.LocationAddress = &key[0]
		selectQ = append(selectQ, "location_address")
	}

	key, exist = form["location_city"]
	if exist {
		event.LocationCity = &key[0]
		selectQ = append(selectQ, "location_city")
	}

	key, exist = form["location_province"]
	if exist {
		event.LocationProvince = &key[0]
		selectQ = append(selectQ, "location_province")
	}

	key, exist = form["start_date"]
	if exist {
		event.StartDate = &key[0]
		selectQ = append(selectQ, "start_date")
	}

	key, exist = form["finish_date"]
	if exist {
		event.FinishDate = &key[0]
		selectQ = append(selectQ, "finish_date")
	}

	key, exist = form["status"]
	if exist {
		b, _ := strconv.ParseInt(key[0], 10, 32)
		c := layer.EventStatus(b)
		event.Status = &c
		selectQ = append(selectQ, "status")
	}

	key, exist = form["presence_question"]
	if exist {
		event.PresenceQuestion = &key[0]
		selectQ = append(selectQ, "presence_question")
	}

	key, exist = form["media_link"]
	if exist {
		event.MediaLink = &key[0]
		selectQ = append(selectQ, "media_link")
	}

	return selectQ, nil
}
