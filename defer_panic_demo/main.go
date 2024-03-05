package main

import "fmt"

func main() {
	fmt.Println("main func start")
	defer func() {
		fmt.Println("main defer func 1")
	}()
	s := test()
	fmt.Println("main get test() return:", s)

}

func test() (str string) {
	defer func() {
		//捕获panic
		if msg := recover(); msg != nil {
			fmt.Println("test defer func1 捕获到错误:", msg)
		}
		str = "bbb"
	}()

	defer func() {
		fmt.Println("test defer func2")
	}()

	defer func() {
		fmt.Println("test defer func3")
	}()

	str = "aaa"

	fmt.Println("panic抛出前")
	panic("test painc")
	fmt.Println("panic抛出后")

	return str
}
