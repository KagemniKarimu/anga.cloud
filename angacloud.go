package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var mu sync.Mutex
var totalRequests int
var startTime = time.Now()
var default404 http.Handler = http.HandlerFunc(errorPage)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/debug", debugInfo)
	r.NotFoundHandler = default404
	fmt.Print("anga.cloud is active!\n")
	fmt.Print("http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	increaseCount()
	content, err := os.ReadFile("error.html")
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	fmt.Fprint(w, string(content))
}

func increaseCount() {
	mu.Lock()
	totalRequests++
	mu.Unlock()
}

func home(w http.ResponseWriter, r *http.Request) {
	increaseCount()
	w.Header().Set("Content-Type", "text/html")
	content, err := os.ReadFile("index.html")
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	fmt.Fprintf(w, string(content))
}

func debugInfo(w http.ResponseWriter, r *http.Request) {
	increaseCount()
	upTime := time.Since(startTime)
	fmt.Fprintf(w, "~ANGA.CLOUD~\n")
	fmt.Fprintf(w, "*************\n")
	fmt.Fprintf(w, "Requests: %d\n", totalRequests)
	fmt.Fprintf(w, "Start Time: %q\n Uptime: %s\n", startTime, upTime)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// fmt.Fprintf(w, "Host OS:%s\n", )
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
