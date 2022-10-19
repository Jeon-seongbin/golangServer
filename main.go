package main

import (
	"fmt"
	"log"
	"net/http"
)

// http server
func main() {
	http.HandleFunc("/hello", helloHander)

	fmt.Println("http server started")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	// http : request -> response -> cut
}

func helloHander(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Method not surpported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello!")
}
