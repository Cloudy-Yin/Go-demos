package main

import (
	"fmt"
	"strings"
)

func countSentence(testString string) map[string]int {
	stringAfterop := strings.Split(testString, " ")
	mapDemo := make(map[string]int, len(stringAfterop))
	for k, v := range stringAfterop {
		fmt.Printf("%v, %v\n", k, v)
		mapDemo[v]++
	}
	fmt.Println(mapDemo)
	return mapDemo
}

func main() {
	var testString = "how do you do"
	result := countSentence(testString)
	for stringWord, mapResult := range result {
		fmt.Printf("the word:%v appeared %v times in :%v\n", stringWord, mapResult, testString)
	}
}
