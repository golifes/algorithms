package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
)

func main() {
	//var ch chan struct{}
	//ch <- struct{}{}
	//<-ch

	//leak()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait()
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

func fo() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
