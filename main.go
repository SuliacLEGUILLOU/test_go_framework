package main

import (
	"github.com/SuliacLEGUILLOU/test_go_framework/lib"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

// TODO error management
func getFizzy(c *gin.Context){
	var int1, _ = strconv.Atoi(c.Param("int1"))
	var int2, _ = strconv.Atoi(c.Param("int2"))
	var limit, _ = strconv.Atoi(c.Param("limit"))
	var str1 = c.Param("str1")
	var str2 = c.Param("str2")

	var res =  make([]string, limit)
	for i := 0; i < limit; i++ {
		str := fizzOrBuzz.FizzOrBuzz(i + 1, int1, int2, str1, str2)

		res[i] = str
	}
	c.IndentedJSON(http.StatusOK, res)
}

func main() {
	router := gin.Default()

	router.GET("/:int1/:int2/:limit/:str1/:str2", getFizzy)
	fmt.Println("Starting...")
	router.Run("localhost:8080")
}
