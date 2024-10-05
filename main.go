package main

// the package that our code is part of
// the way to separate code into logical parts
// main package and main function is the starting point

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>
	FAQ Page
	</h1>
	<p>
	<b> Q: Is there a free version? </b>
	<br />
	A: Yes! We offer a free trial for 30 days on any paid plans.
	</p>
	<p>
	<b> Q: What are your support hours? </b>
	<br />
	A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.
	</p>
	<p>
	<b> Q: How do I contact support? </b>
	<br />
	A: Email us - <a href="mailto:sultansakyp77@gmail.com">sultansakyp77@gmail.com</a>
	</p>
	
	`)
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/faq":
// 		faqHandler(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

func main() {
	r := chi.NewRouter()

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contact", contactHandler)
	//var router http.HandlerFunc = pathHandler

	fmt.Println("Starting the server on :3000... ")
	http.ListenAndServe(":3000", r)

}
