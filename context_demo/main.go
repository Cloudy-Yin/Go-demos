package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

//withdeadline返回父上下文的副本，并将deadline调整为不迟于d。如果父上下文的deadline已经早于d，则WithDeadline(parent, d)在语义上等同于父上下文。
//当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。
/*
func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	select {
	case <-time.After(10 * time.Millisecond):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}*/

//WithTimeout   取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel，通常用于数据库或者网络连接的超时控制
type workername string

func worker(ctx context.Context) {
	key := workername("name")
	val := ctx.Value(key).(string)
	ctx = context.WithValue(ctx, workername("name2"), "lisi")
	go worker2(ctx)

	fmt.Printf("%v worker start to work......\n", val)
	for {
		time.Sleep(time.Millisecond * 100)
		select {
		case <-ctx.Done():
			fmt.Printf("%v worker over\n", val)
			return
		default:
			fmt.Printf("%v worker is running \n", val)
		}
	}
}

func worker2(ctx context.Context) {
	key := workername("name2")
	val := ctx.Value(key).(string)

	fmt.Printf("%v worker2 start to work......\n", val)
	for {
		time.Sleep(time.Millisecond * 100)
		select {
		case <-ctx.Done():
			fmt.Printf("%v worker2 over\n", val)
			return
		default:
			fmt.Printf("%v worker2 is running \n", val)
		}
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, workername("name"), "zhangsan")
	defer cancel()

	go worker(ctx)
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Second * 3)

	fmt.Println("main over")
	fmt.Println(runtime.NumGoroutine())
}

//WithCancel返回带有新Done通道的父节点的副本。当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论先发生什么情况。
//取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。
/*
func worker(ctx context.Context) {
	for {
		fmt.Println("worker running")
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			fmt.Println("worker over")
			return
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second)
	fmt.Println("main over")
}

*/

/*  WithValue函数能够将请求作用域的数据与 Context 对象建立关系。

type TraceCode string

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}

	//LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			fmt.Println("worker done!")
			return
		default:
		}
	}
	//fmt.Println("worker done!")

}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")

	go worker(ctx)
	fmt.Printf("goroutine num : %v\n", runtime.NumGoroutine())
	time.Sleep(time.Second * 3)
	cancel()
	fmt.Println("over")
	fmt.Printf("goroutine num : %v\n", runtime.NumGoroutine())
}

*/
