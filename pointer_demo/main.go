package main

import (
	"fmt"
	"unsafe"
)

func main() {
	num := 12345
	numptr := &num
	fmt.Printf("address is :%v, %v, %v, %v\n", numptr, &numptr, *numptr, &num)

	numptr2 := (*int64)(unsafe.Pointer(numptr))
	fmt.Printf("address is :%v, %v, %v\n", numptr2, &numptr2, *numptr2)

}
