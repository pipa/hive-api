package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Im the API!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8888", nil)
	fmt.Println("GoLang server started on: 8888")
}
