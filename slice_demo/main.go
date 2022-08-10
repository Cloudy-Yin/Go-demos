package main

import (
	"fmt"
	"unsafe"
)

/*
func test() {
	var num1 []int
	var num2 = []int{1, 2, 3}
	num3 := append(num1, num2...)
	fmt.Println(len(num3), num3)

	var num4 []interface{}
	num5 := []int{1, 2, 3}
	num6 := append(num4, num5)
	num6 = append(num6, 1, 2, 3, 4, 5)
	fmt.Println(len(num6), num6)
}

func test1() {

	s := make([]int, 0)

	oldCap := cap(s)

	for i := 0; i < 2048; i++ {
		s = append(s, i)

		newCap := cap(s)

		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] oldcap = %-4d  |  after append %-4d  newcap = %-4d\n", 0, i-1, oldCap, i, newCap)
			oldCap = newCap
		}
	}

}*/

func test2() {
	s := []int{5}
	fmt.Printf("%v, %p, %v\n", &s[0], &s, unsafe.Sizeof(s))
	s = append(s, 7)
	fmt.Printf("%v, %p\n", &s[0], &s)
	s = append(s, 9)
	fmt.Printf("%v, %p\n", &s[0], &s)
	x := append(s, 11)
	fmt.Printf("%v, %p\n", &s[0], &s)
	x = append(x, 12)
	fmt.Printf("%v, %p\n", &x[0], &x)
	x[0] = 1
	y := append(x, 12)
	y[0] = 2
	fmt.Println(s, x, y, cap(s), cap(x), cap(y))
}

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
}

func test3() {
	s := []int{1, 1, 1}
	newS := myAppend(s)

	fmt.Println(s)
	fmt.Println(newS)

	s = newS

	myAppendPtr(&s)
	fmt.Println(s)
}

func main() {
	//test1()
	test2()
	fmt.Println("--------------")
	test3()

	// var a = [5]int{1, 2, 3, 4, 5}
	// var r [5]int
	// for i, v := range &a {
	// 	if i == 0 {
	// 		a[1] = 12
	// 		a[2] = 13

	// 	}

	// 	r[i] = v

	// }
	// fmt.Println(r)
	// fmt.Println(a)
}
