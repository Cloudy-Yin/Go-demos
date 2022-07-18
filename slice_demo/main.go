package main

import "fmt"

func main() {
	var num1 []int
	var num2 = []int{1, 2, 3}
	num3 := append(num1, num2...)
	fmt.Println(len(num3), num3)

	var num4 []interface{}
	num5 := []int{1, 2, 3}
	num6 := append(num4, num5)
	num6 = append(num6, 1, 2, 3, 4, 5)
	fmt.Println(len(num6), num6)

	fmt.Println("this is branch bugfix111")

}
