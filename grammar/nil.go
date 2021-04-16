package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var m map[string]int
	println(m == nil)
	fmt.Println(m == nil)
	fmt.Println(unsafe.Sizeof(m))
}
