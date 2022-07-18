package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func getReq(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "query is %v\n", values)
}

func getheader(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.Header)
	fmt.Fprintln(w, data)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL) //获取整个URL
	fmt.Fprintln(w, string(data))
	data, _ = json.Marshal(r.URL.Query()) //获取URL中的请求参数
	fmt.Fprintln(w, string(data))

}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/getreq", getReq)
	http.HandleFunc("/wholeurl", wholeUrl)
	http.HandleFunc("/getheader", getheader)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
