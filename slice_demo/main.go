package main

import (
	"fmt"
)

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

}

func test2() {
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := append(s, 11)
	x = append(x, 12)
	x[0] = 1
	y := append(x, 12)
	y[0] = 2
	fmt.Println(s, x, y, cap(s), cap(x), cap(y))
}

func main() {
	//test1()
	test2()
}
