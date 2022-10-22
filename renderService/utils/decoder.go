package utils

import (
	"renderService/layer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DecodeUser(c *gin.Context, user *layer.User) error {
	err := c.Request.ParseForm()
	if err != nil {
		return err
	}

	form := c.Request.Form

	key, exist := form["name"]
	if exist {
		user.Name = key[0]
	}

	key, exist = form["desc"]
	if exist {
		user.Desc = key[0]
	}

	key, exist = form["email"]
	if exist {
		user.Email = key[0]
	}

	key, exist = form["password"]
	if exist {
		user.Password = key[0]
	}

	return nil
}

func DecodeEvent(c *gin.Context, event *layer.Event) error {
	err := c.Request.ParseForm()
	if err != nil {
		return err
	}

	form := c.Request.Form

	key, exist := form["name"]
	if exist {
		event.Name = key[0]
	}

	key, exist = form["desc"]
	if exist {
		event.Desc = key[0]
	}

	key, exist = form["requirement"]
	if exist {
		event.Requirement = &key[0]
	}

	key, exist = form["need_payment"]
	if exist {
		b, _ := strconv.ParseBool(key[0])
		event.NeedPayment = b
	}

	key, exist = form["need_id"]
	if exist {
		b, _ := strconv.ParseBool(key[0])
		event.Need_ID = b
	}

	key, exist = form["payment_id"]
	if exist {
		b, _ := strconv.ParseUint(key[0], 10, 32)
		a := uint32(b)
		event.Payment_ID = &a
	}

	key, exist = form["creator_id"]
	if exist {
		b, _ := strconv.ParseUint(key[0], 10, 32)
		event.Creator_ID = uint32(b)
	}

	key, exist = form["payment_price"]
	if exist {
		b, _ := strconv.ParseFloat(key[0], 10)
		event.PaymentPrice = b
	}

	key, exist = form["event_category"]
	if exist {
		b, _ := strconv.ParseInt(key[0], 10, 32)
		event.EventCategory = layer.EventCategory(int32(b))
	}

	key, exist = form["is_offline"]
	if exist {
		b, _ := strconv.ParseBool(key[0])
		event.IsOffline = b
	}

	key, exist = form["location_address"]
	if exist {
		event.LocationAddress = key[0]
	}

	key, exist = form["location_city"]
	if exist {
		event.LocationCity = key[0]
	}

	key, exist = form["location_province"]
	if exist {
		event.LocationProvince = key[0]
	}

	key, exist = form["start_date"]
	if exist {
		event.StartDate = key[0]
	}

	key, exist = form["finish_date"]
	if exist {
		event.FinishDate = key[0]
	}

	key, exist = form["status"]
	if exist {
		b, _ := strconv.ParseInt(key[0], 10, 32)
		event.Status = layer.EventStatus(int32(b))
	}

	key, exist = form["presence_question"]
	if exist {
		event.PresenceQuestion = key[0]
	}

	return nil
}
