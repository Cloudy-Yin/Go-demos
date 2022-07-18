package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var Counter int = 0

func Routine() {
	for count := 0; count < 2; count++ {
		value := Counter
		//time.Sleep(time.Second * 1)
		value++
		Counter = value

	}
	wg.Done()
}

func main() {

	for routine := 1; routine < 2; routine++ {
		wg.Add(1)
		go Routine()
	}

	wg.Wait()
	fmt.Printf("fina counter is :%d\n", Counter)
}
