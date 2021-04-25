package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"golang.org/x/sync/singleflight"
	"log"
	"strconv"
	"time"
)

/**
https://blog.csdn.net/weixin_39895181/article/details/112116006
https://segmentfault.com/a/1190000019661223
https://pkg.go.dev/golang.org/x/sync
https://writings.sh/post/consistent-hashing-algorithms-part-4-maglev-consistent-hash


func xxx () {
	// 大部分协程都卡在这里(11个)
	// 这个锁的效果主要是流控，limit 值初始化赋值，可以是 1，也可以是其他；
	// locker 为 blockingKeyCountLimit 类型
	limitLocker.Acquire( key )
	defer limitLocker.Release( key )
	// 获取数据    getBytesNolc( key , ...)}
	func getBytesNolc () {    // ...
	// 下面就是 singleflight.Group 的用法，防穿透
	// 同一时间只允许一个人去后端更新    ret, err = x.Group.Do(id, func() (interface{}, error)
	{        // 去服务后台获取，更新数据；    })    // ...}


*/
var single singleflight.Group

func SingleF(key string, Id int) (string, error) {
	fmt.Println("start to get and set cache", Id)
	weighted := semaphore.NewWeighted(100)
	weighted.Acquire(context.Background(), int64(Id))

	fn := func() (interface{}, error) {
		fmt.Println("request is setting cache", Id)
		time.Sleep(1 * time.Second)
		if Id > 10 {
			panic("err")
		}
		return strconv.Itoa(Id), nil
	}
	weighted.Release(int64(Id))

	do, _, _ := single.Do(key, fn)

	return do.(string), nil

}

func main() {
	key := "request_key"

	for i := 0; i <= 200; i++ {
		go func(i int) {
			singleF, _ := SingleF(key, i)
			log.Printf("id is %v,data is %v---", i, singleF)
		}(i)

	}
	time.Sleep(2 * time.Second)
}
