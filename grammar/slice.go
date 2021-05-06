package main

import (
	"fmt"
)

func main() {
	data := []int{10, 20, 30, 4, 5, 6}
	s := data[0:2]
	fmt.Println(len(s), cap(s))
	//for i, x := range data {
	//	fmt.Println(data)
	//	if i == 0 {
	//		data[0] += 100
	//		data[1] += 100
	//		data[2] += 100
	//	}
	//	fmt.Printf("x : %d,data:%d\n", x, data[i])
	//}
	//
	//for i, x := range data[:] {
	//	fmt.Println(data)
	//	if i == 0 {
	//		data[0] += 100
	//		data[1] += 100
	//		data[2] += 100
	//	}
	//	fmt.Printf("x : %d,data:%d\n", x, data[i])
	//}
}
