package main

import (
	"fmt"
	"net/http"
)

func samplehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, from Docker container!")
}

func main() {
	http.HandleFunc("/", samplehandler)
	http.ListenAndServe(":8080", nil)
}
