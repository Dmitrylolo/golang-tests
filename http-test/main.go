package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
	fmt.Println("Endpoint hit: HomePage")
}

func Setup() {
	http.HandleFunc("/", HomePage)
}

func main() {
	Setup()
	// http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
