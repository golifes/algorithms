package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

/*
	https://zhuanlan.zhihu.com/p/32529039
*/
func main() {
	//var ch chan struct{}
	//ch <- struct{}{}
	//<-ch

	//leak()
	//finish()
	//traceX()
	nilChannel()

}

func leak() {
	taskChan := make(chan int)
	consumer := func() {
		for task := range taskChan {
			i := task
			fmt.Println("consumer--->", i)
		}
	}
	producer := func() {
		for i := 0; i < 10; i++ {
			taskChan <- i
			fmt.Println("producer--->", i)
		}
		//close(taskChan)
	}
	fmt.Println("start g num", runtime.NumGoroutine())

	go consumer()
	go producer()

	go func() {
		for {
			fmt.Println("g num", runtime.NumGoroutine())
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(5 * time.Second)

	f, _ := os.Create("./goroutine.prof")
	lookup := pprof.Lookup("goroutine")
	lookup.WriteTo(f, 0)

}

func finish() {
	ch := make(chan bool)
	go func() {
		ch <- true
		close(ch)
	}()
	time.Sleep(1)
	select {
	case result := <-ch:
		fmt.Println(result)
		return
	//case <-time.After(1):
	//	fmt.Println("timeout")
	default:
		println("error")
	}
	//channel 什么时候关闭
}

func traceX() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan string)

	go func() {
		ch <- "go搞事"
	}()

	<-ch
}

func nilChannel() {
	//data := make(chan int, 1)
	//
	//for value := range data {
	//	fmt.Println(value)
	//}
	//trace.Start(os.Stderr)
	//defer trace.Stop()

	ch := make(chan string)
	f := func() {
		ch <- "Go语言编程之旅"
	}
	//f()  channel必须和goroutine搭配？？？？
	go f()
	<-ch
}
