package utils

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetStrDefault(c *gin.Context, key string, def string) string {
	ret := c.Param(key)

	if (len(ret) == 0) {
		return def
	}
	return ret
}

func GetIntDefault(c *gin.Context, key string, def int) int {
	val := c.Param(key)

	if (len(val) == 0) {
		return def
	}
	ret, err := strconv.Atoi(val)
	if (err != nil) {
		return def
	}
	return ret
}