package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker(ch1, ch2 <-chan int, stopch chan struct{}) {
	for {
		select {
		case <-stopch:
			fmt.Println("The worker is over")
			return
		case job1 := <-ch1:
			fmt.Println("job1 is : ", job1)
		case job2 := <-ch2:
		priority:
			for {
				select {
				case job1 := <-ch1:
					fmt.Println("job1 is : ", job1)
				default:
					break priority
				}
			}

			fmt.Println("job2 is : ", job2)
		}
	}
}

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	stopch := make(chan struct{})

	go worker(ch1, ch2, stopch)

	go func() {
		for i := 1; i < 5; i++ {
			ch2 <- i
		}
	}()

	go func() {
		for i := 1; i < 5; i++ {
			ch1 <- i
		}
	}()

	time.Sleep(time.Second * 2)
	stopch <- struct{}{}
	time.Sleep(time.Second * 2)
	fmt.Println(runtime.NumGoroutine())
}
