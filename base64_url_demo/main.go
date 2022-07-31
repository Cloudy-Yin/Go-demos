package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
)

func Encrypt(encrypt_key string, plainText string) string {
	key, _ := base64.StdEncoding.DecodeString(encrypt_key)
	encryptBlock, _ := aes.NewCipher(key)

	aesGcm, _ := cipher.NewGCM(encryptBlock)
	nonce := make([]byte, 12)

	_, _ = io.ReadFull(rand.Reader, nonce)

	seal := aesGcm.Seal(nil, nonce, []byte(plainText), nil)
	// 这里打印的字符串是乱码
	fmt.Println("seal:", string(seal))

	cipherText := base64.StdEncoding.EncodeToString(nonce) + ":" + base64.StdEncoding.EncodeToString(seal)
	//这里打印出的字符串包含 / 字符，该字符是url中用来分隔路径的
	fmt.Println("iv:content:", cipherText)

	// 这里对base64进行编码，转换成web安全的字符串
	cipherText = url.QueryEscape(cipherText)
	fmt.Println("query escape:", cipherText)

	return cipherText
}

func test1() {
	// 进行标准base64加密，将二进制字节转换成文本
	e := base64.StdEncoding.EncodeToString([]byte("Man"))
	fmt.Println(e)
	teststr := e + "+/?,;&@:"
	fmt.Println(url.QueryEscape(teststr))

	// 进行解密
	str, err := base64.StdEncoding.DecodeString(e)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(str))
	fmt.Println("-------------------------")

	// 进行URL和文件的base64加密
	e1 := base64.URLEncoding.EncodeToString([]byte("http://localhost?redirect=http://www.baidu.com/search?query=go学堂&name=渔&夫?,:;+子"))
	fmt.Println(e1)
	fmt.Println(url.QueryEscape(e1))
	// 进行解密
	str1, err := base64.URLEncoding.DecodeString(e1)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(str1))

}

func main() {
	encrypt_key := "jSUYHjkt7WTNx/XjLduwiD+xwJNN97dNgVE1M0y6Nk8="
	plainText := "10"
	cipherText := Encrypt(encrypt_key, plainText)
	fmt.Println(cipherText)
	fmt.Println("-------------------------")
	test1()
}
