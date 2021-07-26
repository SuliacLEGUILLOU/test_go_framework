package main

import (
	"github.com/SuliacLEGUILLOU/test_go_framework/lib"
	"github.com/SuliacLEGUILLOU/test_go_framework/stats"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 *	FizzOrBuzz handler
 */
func getFizzy(c *gin.Context){
	var int1 = utils.GetIntDefault(c, "int1", 3)
	var int2 = utils.GetIntDefault(c, "int2", 5)
	var limit = utils.GetIntDefault(c, "limit", 15)
	var str1 = utils.GetStrDefault(c, "str1", "fizz")
	var str2 = utils.GetStrDefault(c, "str2", "buzz")

	if (limit > 10000){
		c.AbortWithStatus(413)
	} else {
		var res =  make([]string, limit)
		for i := 0; i < limit; i++ {
			str := utils.FizzOrBuzz(i + 1, int1, int2, str1, str2)
	
			res[i] = str
		}
		c.IndentedJSON(http.StatusOK, res)
	}
}

func main() {
	router := gin.Default()

	router.Use(stats.Middleware())

	router.GET("/:limit/:int1/:int2/:str1/:str2", getFizzy)
	router.GET("/:limit/:int1/:int2", getFizzy)
	router.GET("/:limit", getFizzy)
	fmt.Println("Starting...")
	router.Run("localhost:8080")
}
