package utils

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

/**
 *	Return a string param from the request or a default string if empty
 */
func GetStrDefault(c *gin.Context, key string, def string) string {
	ret := c.Param(key)

	if (len(ret) == 0) {
		return def
	}
	return ret
}

/**
 *	Return a int param from the request or a default int if empty
 */
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