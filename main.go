package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	fmt.Println("oldman is ready to Go!")

	// HTTP2 server
	srv := &http.Server{Addr: ":1204", Handler: http.HandlerFunc(http2Hello)}
	fmt.Println("HTTP2 server is serving on https://x.x.x.x:1204")
	go func() {
		log.Fatal(srv.ListenAndServeTLS("key/server.crt", "key/server.key"))
	}()

	// HTTP1.1 server
	http.HandleFunc("/", http1Hello)
	fmt.Println("HTTP1.1 server is serving on http://x.x.x.x:1203")
	http.ListenAndServe(":1203", nil)
}

func http1Hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP1 server: %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	hello(w, r)
}

func http2Hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("HTTP2 server: %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	hello(w, r)
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
