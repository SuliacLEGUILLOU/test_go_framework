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

/**
 *	Stats handler
 */
 func stats_handler(c *gin.Context){
	// Those variable should always be set by a middleware
	path, _ := c.Get("stat_top")
	hit, _ := c.Get("stat_hit")

	c.JSON(200, gin.H{
		"path": path,
		"hit": hit,
	})
}

/**
 *	Healthcheck handler for the load balancer
 */
func healthcheck(c *gin.Context){
	c.JSON(200, gin.H{
		"code": "OK",
	})
}

func main() {
	router := gin.Default()

	router.GET("/healthcheck", healthcheck)
	router.GET("/stats", stats_handler)

	// Route declared after this point are include in statistics
	router.Use(stats.Middleware())
	router.GET("/fb/:limit/:int1/:int2/:str1/:str2", getFizzy)
	router.GET("/fb/:limit/:int1/:int2", getFizzy)
	router.GET("/fb/:limit", getFizzy)
	fmt.Println("Starting...")
	router.Run("localhost:8080")
}
