package main

import (
	"fmt"
	"unicode/utf8"
)

// rune -> []byte
func encodeRune() {
	var r rune = 0x4E2D
	fmt.Printf("the unicode charactor is %c\n", r) // 中
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, r) // 对rune进行utf-8编码
	for i := 0; i < len(buf); i++ {
		fmt.Printf("utf-8 representation is 0x%x\n", buf[i]) // 0xE4B8AD
	}
}

// []byte -> rune
func decodeRune() {
	var buf = []byte{0xE4, 0xB8, 0xAD}
	r, _ := utf8.DecodeRune(buf)                                                             // 对buf进行utf-8解码
	fmt.Printf("the unicode charactor after decoding [0xE4, 0xB8, 0xAD] is %v\n", string(r)) // 中
}

func modifystrval() {
	var byteString string = "hello, world"
	byteArr := []byte(byteString)
	fmt.Println(byteArr, []byte("go"))
	//修改切片中某个值
	byteArr[6] = []byte("g")[0]
	//打应输出
	fmt.Println(string(byteArr))
}

func change(s *[]string) {
	//修改string切片中的某一个元素，再通过指针赋值回原切片，会影响到外部的切片
	tmp := []byte((*s)[0])
	fmt.Println(tmp)
	tmp[0] = []byte("g")[0]
	tmp[5] = []byte("9")[0]
	(*s)[0] = string(tmp)
	(*s)[1] = "Go"
	*s = append(*s, "yes")
}

func main() {
	modifystrval()
	var s = "hello 世界"

	var a = []byte(s)
	for i := 0; i < len(a); i++ {
		fmt.Printf("0x%x  ", a[i])

	}

	fmt.Println()

	var b = []rune(s)
	for i, v := range b {
		fmt.Printf("the %v value is : %v ", i, string(v))
	}
	fmt.Println()

	encodeRune()
	decodeRune()
}
