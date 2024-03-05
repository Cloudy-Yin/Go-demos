package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 3.1415
	y := 3
	z := "sheena"
	u := true

	fmt.Println("x修改前的value为:", reflect.ValueOf(x))
	fmt.Println("y修改前的value为:", reflect.ValueOf(y))
	fmt.Println("z修改前的value为:", reflect.ValueOf(z))
	fmt.Println("u修改前的value为：", reflect.ValueOf(u))

	//通过反射传入变量x的地址，并且通过Ele
	rex := reflect.ValueOf(&x).Elem()
	rey := reflect.ValueOf(&y).Elem()
	rez := reflect.ValueOf(&z).Elem()
	reu := reflect.ValueOf(&u).Elem()

	//判断是否可以修改变量x的值，若可以，则用Set()方法进行修改
	if rex.CanSet() {
		ax := reflect.ValueOf(61.23466) // 使用Set方法修改值，Set方法接收的是ValueOf的返回值
		rex.Set(ax)
		fmt.Println("x修改后的value为：", reflect.ValueOf(x))
	} else {
		fmt.Println("该变量不能修改")
	}

	if rey.CanSet() {
		ay := reflect.ValueOf(10000) // 使用Set方法修改值，Set方法接收的是ValueOf的返回值
		rey.Set(ay)
		fmt.Println("y修改后的value为：", reflect.ValueOf(y))
	} else {
		fmt.Println("该变量不能修改")
	}

	if rez.CanSet() {
		az := reflect.ValueOf("hello world") // 使用Set方法修改值，Set方法接收的是ValueOf的返回值
		rez.Set(az)
		fmt.Println("z修改后的value为：", reflect.ValueOf(z))
	} else {
		fmt.Println("该变量不能修改")
	}

	if reu.CanSet() {
		au := reflect.ValueOf(false) // 使用Set方法修改值，Set方法接收的是ValueOf的返回值
		reu.Set(au)
		fmt.Println("u修改后的value为：", reflect.ValueOf(u))
	} else {
		fmt.Println("该变量不能修改")
	}
}
