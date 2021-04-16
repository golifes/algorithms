package main

import "fmt"

func main() {
	s := "12334"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
