package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex
var totalRequests int
var startTime = time.Now()

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	increaseCount()
	echo(w, r)
}

func increaseCount() {
	mu.Lock()
	totalRequests++
	mu.Unlock()
}

func echo(w http.ResponseWriter, r *http.Request) {
	upTime := time.Since(startTime)
	fmt.Fprintf(w, "~ANGA.CLOUD~\n")
	fmt.Fprintf(w, "*************\n")
	fmt.Fprintf(w, "Requests: %d\n", totalRequests)
	fmt.Fprintf(w, "Start Time: %q\n Uptime: %s\n", startTime, upTime)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
