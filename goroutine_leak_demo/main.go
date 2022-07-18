package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/*泄露的原因大多集中在：

1. Goroutine 内正在进行 channel/mutex 等读写操作，但由于逻辑问题，某些情况下会被一直阻塞。
2. Goroutine 内的业务逻辑进入死循环，资源一直无法释放。
3. Goroutine 内的业务逻辑进入长时间等待，有不断新增的 Goroutine 进入等待。*/

/* 发送不接收 */
func query() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}

func queryAll() int {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func() {
			ch <- query()
		}()
	}
	return <-ch
}

/* 在这个例子中，我们调用了多次 queryAll 方法，并在 for 循环中利用 Goroutine 调用了 query 方法。其重点在于调用 query 方法后的结果会写入 ch 变量中，接收成功后再返回 ch 变量。
最后可看到输出的 goroutines 数量是在不断增加的，每次多 2 个。也就是每调用一次，都会泄露 Goroutine。
原因在于 channel 均已经发送了（每次发送 3 个），但是在接收端并没有接收完全（只返回 1 个 ch），所诱发的 Goroutine 泄露。 */

func test1() {
	for i := 0; i < 3; i++ {
		queryAll()
		fmt.Printf("goroutine : %v; Numcpu : %v; Numcgocall: %v \n", runtime.NumGoroutine(), runtime.NumCPU(), runtime.NumCgoCall())
	}
}

/* 接收不发送 */
/* 在这个例子中，与 “发送不接收” 两者是相对的，channel 接收了值，但是不发送的话，同样会造成阻塞。
但在实际业务场景中，一般更复杂。基本是一大堆业务逻辑里，有一个 channel 的读写操作出现了问题，自然就阻塞了。 */
func test2() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var ch chan int = make(chan int) /*channel 如果忘记初始化，那么无论你是读，还是写操作，都会造成阻塞。调用make函数进行初始化*/
	go func() {
		ch <- 800
	}()

	n := <-ch
	time.Sleep(time.Second)
	fmt.Printf("chan result: %v ", n)
}

/*互斥锁忘记解锁*/
func test3() {
	total := 0
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("total: ", total)
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock() /* 在这个例子中，第一个互斥锁 sync.Mutex 加锁了，但是他可能在处理业务逻辑，又或是忘记 Unlock 了。因此导致后面的所有 sync.Mutex 想加锁，却因未释放又都阻塞住了。一般在 Go 工程中，我们建议如下写法： */
			defer mutex.Unlock()
			total += 1
		}()
	}
}

/*同步锁使用不当*/
/* 在这个例子中，我们调用了同步编排 sync.WaitGroup，模拟了一遍我们会从外部传入循环遍历的控制变量。
但由于 wg.Add 的数量与 wg.Done 数量并不匹配，因此在调用 wg.Wait 方法后一直阻塞等待。 */
func handle(v int) {
	var wg sync.WaitGroup
	//wg.Add(5)
	for i := 0; i < v; i++ {
		wg.Add(1) //在 Go 工程中推荐这样使用
		fmt.Println("something happens")
		wg.Done()
	}
	wg.Wait()
}

func test4() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	go handle(3)
	time.Sleep(time.Second)
}

func test5() {
	// 模拟单核 CPU
	runtime.GOMAXPROCS(1)

	// 模拟 Goroutine 死循环
	go func() {
		for {

		}
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("脑子进煎鱼了")
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
}
