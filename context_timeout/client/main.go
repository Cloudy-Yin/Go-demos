package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

type respData struct {
	resp *http.Response
	err  error
}

func docall(ctx context.Context) {
	transport := http.Transport{
		DisableKeepAlives: true,
	}

	client := http.Client{
		Transport: &transport,
	}

	respChan := make(chan *respData, 1)
	req, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1:8000/", nil)
	if err != nil {
		fmt.Printf("new request failed, err : %v\n", err)
		return
	}

	// var wg sync.WaitGroup
	// wg.Add(1)
	go func() {
		resp, err := client.Do(req) //传入带超时context的req
		time.Sleep(time.Second * 2)
		fmt.Printf("client.do resp:%v, err:%v\n", resp, err)
		rd := &respData{
			resp: resp,
			err:  err,
		}

		respChan <- rd
		//wg.Done()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
		return
	case result := <-respChan:
		fmt.Println("api call success")
		if result.err != nil {
			fmt.Printf("call api fail , err:%v\n", result.err)
			return
		}

		defer result.resp.Body.Close()
		data, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("rep:%v\n", string(data))
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)

	for i := 0; i < 3; i++ {
		go docall(ctx)
	}

	time.Sleep(time.Second * 3)
	cancel()
	fmt.Println(runtime.NumGoroutine())
}
