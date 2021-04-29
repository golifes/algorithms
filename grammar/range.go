package main

import (
	"fmt"
)

func main() {
	//arr := []int{1, 2, 3}
	//
	//for idx, v := range arr {
	//	arr = append(arr, v)
	//	fmt.Println(arr, idx)
	//}
	//
	//arr = arr[0:0]
	//fmt.Println(arr)
	//
	//for i := range arr {
	//	arr[i] = 0
	//}
	//
	//fmt.Println(arr)
	count()
}

func count() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}

	counter := 0
	for k := range m {
		if k == "A" {
			delete(m, "A")
		}
		counter++
	}
	fmt.Println(counter, m)
}
