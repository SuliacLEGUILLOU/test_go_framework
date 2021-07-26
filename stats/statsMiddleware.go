package stats

import (
	"fmt"
	// "sync"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	var stats = make(map[string]int)
	var topUsed string
	// var m = sync.Mutex()

	return func(c *gin.Context) {
		// m.lock()
		path := c.Request.URL.String()
		_, ok := stats[path]
		if !ok{
			stats[path] = 0
		}
		stats[path]++
		if(stats[topUsed] < stats[path]){
			topUsed = path
		}
		fmt.Println(topUsed)
		// m.unlock()
        c.Next()
    }
}