package utils

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getID(c *gin.Context, what string) (*uint32, error) {
	idstr := c.Param(what)
	id, err := strconv.ParseUint(idstr, 10, 32)
	if err != nil {
		return nil, err
	}
	idConverted := uint32(id)
	return &idConverted, nil
}

func GetParamSubscriptionID(c *gin.Context) (*uint32, error) {
	return getID(c, "subscription-id")
}

func GetParamPublisherID(c *gin.Context) (*uint32, error) {
	return getID(c, "publisher-id")
}

func GetParamRegisterID(c *gin.Context) (*uint32, error) {
	return getID(c, "register-id")
}

func GetParamEventID(c *gin.Context) (*uint32, error) {
	return getID(c, "event-id")
}

func GetParamUserID(c *gin.Context) (*uint32, error) {
	return getID(c, "user-id")
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
		query = convertQuery(&queryStr)
		return
	} else {
		query = nil
	}
	return
}

func convertQuery(q *string) *string {
	splitQuery := strings.Split(*q, "|")
	stringArrayQuery := []string{}

	for index, Query := range splitQuery {
		QueryData := strings.Split(Query, ",")
		if len(splitQuery)-1 == index {
			QueryConvertToSnakeCase := QueryData[0] + " LIKE " + "'%" + QueryData[1] + "%'"
			stringArrayQuery = append(stringArrayQuery, QueryConvertToSnakeCase)
		} else {
			QueryConvertToSnakeCase := QueryData[0] + " LIKE '%" + QueryData[1] + "%' AND"
			stringArrayQuery = append(stringArrayQuery, QueryConvertToSnakeCase)
		}
	}

	QueryFix := strings.Join(stringArrayQuery, " ")
	return &QueryFix

}
