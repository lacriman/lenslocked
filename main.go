package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "Welcome to my awesome site!")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:jon@calhoun.io\">jon@calhoun.io</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<h1>FAQ page</h1>

		<div style="margin-bottom: 40px;">
			<p><strong>Q:</strong> Is there a free version?</p>
			<p><strong>A:</strong> Yes, we offer a free trial for 30 days on any paid plans.</p>
		</div>

		<div style="margin-bottom: 40px;">
			<p><strong>Q:</strong> What are your support hours?</p>
			<p><strong>A:</strong> We have support staff answering emails 24/7, though response times may be slower on weekends.</p>
		</div>

		<div style="margin-bottom: 40px;">
			<p><strong>Q:</strong> How do I contact support?</p>
			<p><strong>A:</strong> Email us at <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>.</p>
		</div>
	`)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/faq/{faqID}", func (w http.ResponseWriter, r *http.Request)  {
		faqID:= chi.URLParam(r, "faqID")
		fmt.Fprintf(w, "Faq ID: %s", faqID)
	})
	r.NotFound(func (w http.ResponseWriter, r *http.Request)  {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on port 3000...")
	http.ListenAndServe(":3000", r)
}
