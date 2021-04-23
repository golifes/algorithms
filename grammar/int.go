package main

import (
	"fmt"
	"math"
)

func main() {
	reverse(123)

}

func reverse(x int) int {

	reverseX := 0
	if x != 0 {
		reverseX = reverseX*10 + x%10
		x = x / 10
	}

	if reverseX > math.MaxInt32 || reverseX < math.MinInt32 {
		return 0
	}

	return reverseX
}

func bitMap() {
	/**
	 // int is a signed integer type that is at least 32 bits in size. It is a
	// distinct type, however, and not an alias for, say, int32.
	type int int


	// int32 is the set of all signed 32-bit integers.
	// Range: -2147483648 through 2147483647.
	type int32 int32
	*/
	bit := make(map[int32]bool)
	fmt.Println(bit)
}
