package main

const y = 0x200

//https://gfw.go101.org/article/memory-layout.html#size-and-padding
//https://gfw.go101.org/article/unofficial-faq.html#print-builtin-fmt-log
/*
	常量会被编译器在预处理阶段做处理,不会分配内存空间
	var x= 0x10
	var y = 0x20
	fmt.Println(&x,x)
	fmt.Println(&y,y)  //connot take the address of y


	// go build -gcflags "-N"  xxx.go
	// go tool objdump -s "main\.main" xxx
	func main(){
		const y = 0x200
		println(y)
		//fmt.Println(y)
	}
*/
func main() {
	println(y)

	//const(
	//	_= iota
	//	a
	//	c
	//	d
	//	iota // 作用类似于_,表示可以跳过的值
	//	e
	//	f
	//	g = 20 //显式中断iota,后面的值与这一行相同
	//	h
	//	i
	//	j = iota  //显式恢复
	//	m
	//	n
	//)
	//
	//fmt.Println(a,c,d,e,f,g,h,i,j,m,n)

	//const(
	//	_= iota
	//	a
	//	c
	//	d
	//	e
	//	f
	//	g = 20 //显式中断iota,后面的值与这一行相同
	//	h
	//	i
	//	j = iota  //显式恢复
	//	m
	//	n
	//)
	//
	//fmt.Println(a,c,d,e,f,g,h,i,j,m,n)

}
