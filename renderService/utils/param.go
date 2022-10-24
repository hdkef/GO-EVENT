package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamID(c *gin.Context) (uint32, error) {
	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 32)
	if err != nil {
		return 0, nil
	}
	return uint32(id), nil
}
