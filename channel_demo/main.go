package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

//goroutine超时控制使用 context实现
func worker(ctx context.Context, ch chan struct{}) {

	for {
		select {
		case ch <- struct{}{}:
			fmt.Println("do some staff.....")
			time.Sleep(time.Second * 1)

		case <-ctx.Done():
			close(ch)
			fmt.Println("goroutine over")
			return
		}
	}
	//fmt.Println("goroutine over")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	ch := make(chan struct{})
	go worker(ctx, ch)
	fmt.Println(runtime.NumGoroutine())
	i := 0
	for n := range ch {
		fmt.Println(n)
		i++
		if i == 6 {
			cancel()
		}
	}

	time.Sleep(time.Second * 2)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("main over")

}

/*  goroutine超时控制使用
var wg sync.WaitGroup

func worker(ch chan string) {
	for {
		select {
		case res := <-ch:
			fmt.Println(res)
		case <-time.After(time.Second * 1):
			wg.Done()
			return
		}
	}

}
func main() {

	wg.Add(2)
	ch := make(chan string, 2)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- "job res"
		close(ch)

		wg.Done()
	}()

	go worker(ch)

	wg.Wait()
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("main over")
}

/*
var wg sync.WaitGroup

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	for j := 0; j < 3; j++ {
		wg.Add(1)
		go func() {

			for val := range ch {
				fmt.Println(val)
			}
			// LOOP:
			// 	for {

			// 		select {
			// 		case task, ok := <-ch:
			// 			if !ok {
			// 				ch = nil
			// 				break LOOP
			// 			}
			// 			fmt.Println(task)
			// 		default:
			// 			break LOOP
			// 		}
			// 		// task, ok := <-ch
			// 		// if !ok {
			// 		// 	break
			// 		// }
			// 		// fmt.Println(task)
			// 	}
			wg.Done()
		}()
	}

	wg.Wait()
}

/*
func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- 5
		close(ch1)
	}()

	go func() {
		time.Sleep(time.Second * 3)
		ch2 <- 7
		close(ch2)
	}()

	for {
		select {
		case x, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				fmt.Println(x)
			}
		case x, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				fmt.Println(x)
			}
		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}
	fmt.Println("program end")
}

//channel 是在 produce 函数中被关闭的，这也是 channel 的一个使用惯例，那就是发送端负责关闭 channel。
//这是因为发送端没有像接受端那样的、可以安全判断 channel 是否被关闭了的方法。同时，一旦向一个已经关闭的 channel 执行发送操作，这个操作就会引发 panic
/*
func produce(ch chan<- string) {
	for i := 0; i < 10; i++ {
		ch <- "hello"
		time.Sleep(time.Microsecond)
	}
	close(ch)
}

func consume(ch <-chan string) {
	for n := range ch {
		fmt.Println(n)
	}
}

func test1() {
	ch := make(chan string, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()

	wg.Wait()
}

//无缓冲channel用作信号传递
type signal struct{}

func worker() {
	println("worker is working...")
	time.Sleep(1 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		c <- signal{} //无缓冲channel，使用struct{}{}作为信号
	}()
	return c
}

func test2() {
	println("start a worker...")
	c := spawn(worker)
	<-c
	fmt.Println("worker work done!")
}

func main() {
	test1()
	test2()
}
*/
