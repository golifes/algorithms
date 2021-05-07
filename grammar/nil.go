package main

import (
	"fmt"
)

func main() {
	fmt.Println("echo--->", echo())
}

func echo() int {
	var i int

	defer func() {
		i++
		fmt.Println(i)
	}()

	defer func() {
		i++
		fmt.Println(i)
	}()

	/*  todo ???
	defer fmt.Println(i)
	fmt.Println(i)

	*/
	defer fmt.Println(i)

	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}

	fmt.Println("main ends")

	return i
}
