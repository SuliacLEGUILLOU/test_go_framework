package main

import (
	// "github.com/SuliacLEGUILLOU/test_go_framework/lib/fizzOrBuzz.go"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

// TODO Move to package
func FizzOrBuzz(n, a, b int, str1, str2 string) string {
	if (n % a == 0 && n % b == 0) {
		return str1+str2
	} else if (n % a == 0) {
		return str1
	} else if (n % b == 0) {
		return str2
	} else {
		return strconv.Itoa(n)
	}
}

// TODO error management
func getFizzy(c *gin.Context){
	var int1, _ = strconv.Atoi(c.Param("int1"))
	var int2, _ = strconv.Atoi(c.Param("int2"))
	var limit, _ = strconv.Atoi(c.Param("limit"))
	var str1 = c.Param("str1")
	var str2 = c.Param("str2")

	var res =  make([]string, limit)
	for i := 0; i < limit; i++ {
		str := FizzOrBuzz(i + 1, int1, int2, str1, str2)

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
