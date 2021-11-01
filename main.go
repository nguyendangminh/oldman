package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var sleepTime int

func main() {
	flag.IntVar(&sleepTime, "s", 0, "sleep time (default 0 second)")
	flag.Parse()

	fmt.Println("oldman is ready to Go!")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":1203", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Fprintf(w, fmt.Sprintf("hello, oldman is ready to Go! sleep %d seconds", sleepTime))
}
