package main

// the package that our code is part of
// the way to separate code into logical parts
// main package and main function is the starting point

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!<h1>")

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>Contact Page</h1><p>To get in touch, email me at <a 
	href="mailto:sultansakyp77@gmail.com">sultansakyp77@gmail.com</a>.</p>`)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on :3000... ")
	http.ListenAndServe(":3000", nil)

}
