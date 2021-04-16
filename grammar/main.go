package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("13")
	}()

	{
		defer func() {
			fmt.Println("14")
		}()

		fmt.Println("15")
	}

	defer func() {
		fmt.Println("16")
	}()
}
