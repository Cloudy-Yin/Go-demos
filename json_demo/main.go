package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty"`
	Hobby    []string `json:"hobby,omitempty"`
	*Profile `json:"profile,omitempty"`
	//想要在嵌套的结构体为空值时，忽略该字段，仅添加omitempty是不够的, 还需要使用嵌套的结构体指针
}

type Profile struct {
	Website string `json:"site"`
	Slogan  string `json:"slogan"`
}

func omitemptyDemo() {
	u1 := User{
		Name: "七米",
		Hobby: []string{
			"pingpang",
			"basketball",
		},
	}
	// struct -> json string
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
}

//Go语言中json.Marshal（序列化）与json.Unmarshal（反序列化）的基本用法。
type Person struct {
	Name   string `json:"name"`
	Age    int64  `json:"age,string"`
	Weight float64
}

func test1() {
	p1 := Person{
		Name:   "七米",
		Age:    18,
		Weight: 71.5,
	}
	// struct -> json string
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%v\n", string(b))
	// json string -> struct
	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}

	p2.Age = 28
	p2.Name = "zhangsan"
	fmt.Printf("p2:%#v\n", p2)

	s, err := json.Marshal(p2)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%v\n", string(s))
}

//优雅处理字符串格式的数字
//有时候，前端在传递来的json数据中可能会使用字符串类型的数字，这个时候可以在结构体tag中添加string来告诉json包从字符串中解析相应字段的数据：
type Card struct {
	ID    int64   `json:"id,string"`    // 添加string tag
	Score float64 `json:"score,string"` // 添加string tag
}

func intAndStringDemo() {
	jsonStr1 := `{"id": "1234567","score": "88.50"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("json.Unmarsha jsonStr1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("c1:%#v\n", c1) // c1:main.Card{ID:1234567, Score:88.5}

	b, _ := json.Marshal(c1)
	fmt.Println(string(b))

}

func main() {
	intAndStringDemo()
	fmt.Println("------------------")
	test1()
	fmt.Println("------------------")
	omitemptyDemo()
}
