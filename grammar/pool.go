package main

import (
	"bytes"
	"fmt"
	"sync"
)

var buffers = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func main() {
	pool := sync.Pool{New: func() interface{} { return make([]int, 2) }}
	pool.Put(1)
	get := pool.Get()
	fmt.Println(get)
	pool.Put(2)
	pool.Put(3)

}

func GetBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}

func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	buffers.Put(buf)
}
