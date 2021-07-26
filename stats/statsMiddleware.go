package stats

import (
	// "sync"

	"github.com/gin-gonic/gin"
)

/**
 * TODO: Move this to an async db query, as this will not work with multiple instances
 */
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
		// m.unlock()
		c.Set("stat_top", topUsed)
		c.Set("stat_hit", stats[topUsed])
        c.Next()
    }
}