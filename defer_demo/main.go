package main

import (
	"fmt"
)

//https://blog.csdn.net/anlz729/article/details/105940142?csdn_share_tail=%7B%22type%22%3A%22blog%22%2C%22rType%22%3A%22article%22%2C%22rId%22%3A%22105940142%22%2C%22source%22%3A%22unlogin%22%7D&ctrtid=dHvbD
//Go语言中defer和return执行顺序解析： 先为返回值赋值---》然后执行defer---》 return到函数调用处

func test() string { //无名返回值
	a := "a"
	defer func() {
		a += "c"
	}()

	defer func() {
		a += "d"
	}()

	a = a + "b"

	return a
}

func test1() (a string) { //具名返回值
	a = "a"
	defer func() {
		a += "c"
	}()

	defer func() {
		a += "d"
	}() //defer函数执行顺序遵循先进后出原则

	a = a + "b"

	return a
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func test3() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func test4() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

func main() {
	fmt.Println(test())
	fmt.Println(test1())
	test3()
	test4()
}
