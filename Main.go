package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("GoLang server started on: 8888")
	http.ListenAndServe(":8888", nil)
}
