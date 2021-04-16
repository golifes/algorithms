package main

import (
	"fmt"
	"unsafe"
)

func main() {
	const x, y int = 123, 456
	fmt.Println(x, y)

	//x = 6

	sizeof := unsafe.Sizeof(uintptr(x))
	fmt.Println(sizeof)
}
