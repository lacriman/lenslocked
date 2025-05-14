package main

import (
	"fmt"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "Welcome to my awesome site!")
	fmt.Println("First request handled:", time.Now())

}

func contactHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("Build time:", time.Now())

	http.HandleFunc("/", homeHandler)
	fmt.Println("Starting the server on port 3000...")
	http.ListenAndServe(":3000", nil)
}
