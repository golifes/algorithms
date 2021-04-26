package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
	"net/http"
)

var group singleflight.Group

func main() {
	eng := gin.Default()

	eng.GET("/test", func(c *gin.Context) {
		key := "gin"
		do, err, shared := group.Do(key, func() (interface{}, error) {
			fmt.Println("settings data")
			return "data", nil
		})
		fmt.Println(do.(string), err, shared)
		c.JSON(http.StatusOK, "ok")
	})
	pprof.Register(eng)
	eng.Run(":8081")
}
