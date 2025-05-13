package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my great site!")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on port 3000...")
	http.ListenAndServe(":3000", nil)
}
