package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamID(c *gin.Context) (*uint32, error) {
	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 32)
	if err != nil {
		return nil, err
	}
	idConverted := uint32(id)
	return &idConverted, nil
}

func GetPagination(c *gin.Context) (lid *uint32, lim *uint32, query *string, err error) {
	//extract last-id and limit
	lastIDStr, exist := c.GetQuery("last-id")
	if !exist {
		err = errors.New("last-id doesn't exist")
		return
	}
	limitStr, exist := c.GetQuery("limit")
	if !exist {
		err = errors.New("last-id doesn't exist")
		return
	}
	lastid64, err := strconv.ParseUint(lastIDStr, 10, 64)
	if err != nil {
		return
	}
	lastid := uint32(lastid64)
	lid = &lastid
	limit64, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		return
	}
	limit := uint32(limit64)
	lim = &limit
	//extract query
	queryStr, exist := c.GetQuery("query")
	if exist {
		//convert to query SQL
		fmt.Println(queryStr)
		return
	} else {
		query = nil
	}
	return
}
