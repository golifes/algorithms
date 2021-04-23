package main

import "fmt"

//func f() (result int) {
//	defer func() {
//		result++
//	}()
//	return 0
//}

//func f() (r int) {
//	t := 5
//	defer func() {
//		t = t + 5
//	}()
//	return t
//}

func f() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

//defer是在return之前执行的
/*
	先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。
*/
func main() {
	result := f()
	fmt.Println(result)
}
