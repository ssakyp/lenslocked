package main

// the package that our code is part of
// the way to separate code into logical parts
// main package and main function is the starting point

import (
	"fmt"
	"net/http"
	"time"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!<h1>")

}

func main() {
	http.HandleFunc("/", handlerFunc)
	time.Sleep(3 * time.Second)

	fmt.Println("Starting the server on :3000... ")
	http.ListenAndServe(":3000", nil)

}
