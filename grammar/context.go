package main

import (
	"context"
	"fmt"
	"time"
)

//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
//	defer cancel()
//
//	go handle(ctx, 100*time.Millisecond)
//
//	select {
//
//	case <-ctx.Done():
//		fmt.Println("main ", ctx.Err())
//	default:
//		println("default")
//	}
//}
//
//func handle(ctx context.Context, duration time.Duration) {
//	select {
//	case <-ctx.Done():
//		fmt.Println("handler ", ctx.Err())
//	case <-time.After(duration):
//		fmt.Println("process request with ", duration)
//	}
//}
func sleepRandom_1(stopChan chan struct{}) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("This is sleep Random 1: %d\n", i)

		i++
		if i == 5 {
			fmt.Println("cancel sleep random 1")
			stopChan <- struct{}{}
			break
		}
	}
}

//https://zhuanlan.zhihu.com/p/76555349
/*
	改进引子里的例子。 sleepRandom_1 结束后，
	会触发 cancelParent() 被调用。所以 sleepRandom_2 中的 ctx.Done() 会被关闭。
	sleepRandom_2 执行退出。
	????

*/
func sleepRandom_2(ctx context.Context) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("This is sleep Random 2: %d\n", i)
		i++

		select {
		case <-ctx.Done():
			fmt.Printf("Why? %s\n", ctx.Err())
			fmt.Println("cancel sleep random 2")
			return
		default:
		}
	}
}

//func main() {
//
//	ctxParent, cancelParent := context.WithCancel(context.Background())
//	ctxChild, _ := context.WithCancel(ctxParent)
//	cancelParent()
//
//	stopChan := make(chan struct{})
//
//	go sleepRandom_1(stopChan)
//	go sleepRandom_2(ctxChild)
//
//	select {
//	case <-stopChan:
//		fmt.Println("stopChan received")
//	}
//
//	for {
//		time.Sleep(1 * time.Second)
//		fmt.Println("Continue...")
//	}
//}
//

func main() {

	/*
		https://zhuanlan.zhihu.com/p/163529509
		channel未关闭
	*/
	ch := make(chan int)
	go func() {
		fmt.Println("hello inline")
		ch <- 1
	}()

	go printHello(ch)

	fmt.Println("Hello from main")
	i := <-ch

	fmt.Println("recieved ", i)
	<-ch

}

func printHello(ch chan int) {
	fmt.Println(" hello from printHello")
	ch <- 2
	//close(ch)
}
