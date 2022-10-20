package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHander(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method not surpported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// http server
func main() {

	//http
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	// static 밑 html파일을 읽어옴

	http.HandleFunc("/hello", helloHander)
	http.HandleFunc("/form", formHandler)

	fmt.Println("http server started")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	// http : request -> response -> cut
}
