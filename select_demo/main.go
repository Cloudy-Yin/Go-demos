package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var once sync.Once

func randNum(jobChan chan<- int64) {
	defer wg.Done()
	for i := 1; i < 100; i++ {
		jobChan <- rand.Int63()
		time.Sleep(time.Millisecond * 10)
	}
	close(jobChan)

}

func sumNum(jobChan <-chan int64, resultChan chan<- int64) {
	defer wg.Done()

	for {
		num, ok := <-jobChan
		if !ok {
			once.Do(func() {
				close(resultChan)
			})
			break
		}
		var t int64
		for num > 0 {
			left := num % 10
			num = num / 10
			t += left
		}
		resultChan <- t
	}

	// for v := range jobChan {
	// 	var s int64
	// 	for v > 0 {
	// 		s += v % 10
	// 		v = v / 10
	// 	}
	// 	resultChan <- s
	// }
}

func main() {
	jobChan := make(chan int64, 1)
	resultChan := make(chan int64, 24)

	wg.Add(1)
	go randNum(jobChan)

	for i := 1; i < 24; i++ {
		wg.Add(1)
		go sumNum(jobChan, resultChan)
	}

	for v := range resultChan {
		fmt.Println(v)
	}

	wg.Wait()

}
