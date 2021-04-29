package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
	"net/http"
	"time"
)

func main() {
	eng := gin.Default()
	key := "gin"
	var group singleflight.Group
	eng.GET("/test", func(c *gin.Context) {
		data, err := getSetData(&group, key)

		if err != nil {
			group.Forget(key)
		}
		go func() {
			time.Sleep(1 * time.Second)

		}()
		c.JSON(http.StatusOK, data)
	})
	pprof.Register(eng)
	eng.Run(":8081")
}

func getSetData(sg *singleflight.Group, key string) (string, error) {
	do, err, _ := sg.Do(key, func() (interface{}, error) {
		select {}
		fmt.Println("settings data", key)
		return "data", nil
	})
	fmt.Println("get data")
	return do.(string), err
}

func getSetDataContext(ctx context.Context, sg *singleflight.Group, key string) (string, error) {
	doChan := sg.DoChan(key, func() (interface{}, error) {
		select {}
		fmt.Println("settings data", key)
		return "data", nil
	})
	select {
	case r := <-doChan:
		return r.Val.(string), r.Err
	case <-ctx.Done():
		return "", ctx.Err()

	}
}

/*

 */
