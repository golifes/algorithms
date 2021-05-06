package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}
	arr3 := arr2[1:]

	d := arr1[:]

	fmt.Println(arr2 == arr1)

	fmt.Println(d, arr3[:])

}
