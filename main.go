package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	fmt.Println("oldman is ready to Go!")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":1203", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	values, _ := query["t"]

	var sleepTime int = 0
	if len(values) > 0 {
		i, err := strconv.Atoi(values[0])
		if err != nil {
			// handle error
			fmt.Println(err)
			http.Error(w, err.Error(), 400)
			return
		}
		sleepTime = i
	}

	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Fprintf(w, fmt.Sprintf("hello, oldman is ready to Go! sleep %d seconds", sleepTime))
}
