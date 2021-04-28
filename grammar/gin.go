package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
	"net/http"
)

func main() {
	eng := gin.Default()

	req := Req{group: singleflight.Group{}}
	eng.GET("/test", req.test)
	pprof.Register(eng)
	eng.Run(":8081")
}

type Req struct {
	group singleflight.Group
}

func (r *Req) test(c *gin.Context) {
	do, err, shared := r.group.Do(c.Request.RequestURI, func() (interface{}, error) {
		fmt.Println("settings data")
		return "data", nil
	})
	fmt.Println(do.(string), err, shared)
	c.JSON(http.StatusOK, "ok")
}

/*

 */
