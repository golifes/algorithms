package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked = 1 << iota
)

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) TryLock() bool {
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		fmt.Println(*(*int32)(unsafe.Pointer(&m.Mutex)))
		return true
	}
	return false
}
func main() {
	var wg sync.WaitGroup
	var mu Mutex

	var count int32

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := mu.TryLock()
			if res {
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("执行结果:%v\n", count)
}
