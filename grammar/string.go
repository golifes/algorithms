package main

import "fmt"

func main() {
	str := "hello"
	b := []byte(str)
	b[1] = 'l'
	fmt.Println(string(b))
}
