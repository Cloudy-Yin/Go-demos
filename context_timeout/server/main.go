package main

import (
	"fmt"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	time.Sleep(time.Second * 10)
	fmt.Fprintf(w, "slow response")

	// number := rand.Intn(2)
	// if number == 0 {
	// 	time.Sleep(time.Second * 10)
	// 	fmt.Fprintf(w, "slow response")
	// 	return
	// }

	// fmt.Fprintf(w, "quick response")
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}

}
