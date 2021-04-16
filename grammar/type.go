package main

type X int

func main() {
	var a int = 100
	var b X = X(a)
	println(b)

	println(X(a) == b)
	println(a == int(b))
}
