package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	//var ch chan struct{}
	//ch <- struct{}{}
	//<-ch

	//leak()
	finish()
}

func leak() {
	taskChan := make(chan int)

	consumer := func() {
		for task := range taskChan {
			_ = task
		}
	}

	producer := func() {
		for i := 0; i < 10; i++ {
			taskChan <- i
		}
	}

	go consumer()
	go producer()

	f, err := os.Create("./cpuProfile")

	if err != nil {
		fmt.Println(err)
	}

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
}

func finish() {
	ch := make(chan bool, 1)
	go func() {
		ch <- true
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(1):
		fmt.Println("timeout")
	}
}
