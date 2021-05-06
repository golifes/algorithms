package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
	"net/http"
)

var group = new(singleflight.Group)

func main() {
	eng := gin.Default()
	key := "gin"

	/*
			ab.exe  -n100 -c50 http://127.0.0.1:8081/test


			 var wg sync.WaitGroup
		    for i := 0; i < n; i++ {
		        wg.Add(1)
		        go func() { // n个协程同时调用了g.Do，fn中的逻辑只会被一个协程执行
		            v, err := g.Do("key", fn)
		            if err != nil {
		                t.Errorf("Do error: %v", err)
		            }
		            if v.(string) != "bar" {
		                t.Errorf("got %q; want %q", v, "bar")
		            }
		            wg.Done()
		        }()
		    }

	*/
	eng.GET("/test", func(c *gin.Context) {
		data, _ := getSetData(group, key)

		//if err != nil {
		//	group.Forget(key)
		//}
		//go func() {
		//	time.Sleep(1 * time.Second)
		//}()
		c.JSON(http.StatusOK, data)
	})
	pprof.Register(eng)
	eng.Run(":8081")
}

func getSetData(sg *singleflight.Group, key string) (string, error) {
	do, err, shared := sg.Do(key, func() (interface{}, error) {
		fmt.Println("settings data", key)
		return "data", nil
	})
	fmt.Println("get data", shared)
	return do.(string), err
}

func getSetDataContext(ctx context.Context, sg *singleflight.Group, key string) (string, error) {
	doChan := sg.DoChan(key, func() (interface{}, error) {
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
