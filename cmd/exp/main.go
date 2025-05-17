package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

type User struct {
	Name   string
	Number string
	Email  string
	Tasks  []string
}

func executeTemplate(w http.ResponseWriter, filepath string, user interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//  Used filepath.Join(...) to avoid platform-specific path issues — super important for cross-platform deployments
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, user)
	if err != nil {
		log.Printf("executing template %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:   "Yaroslav Proskurin",
		Number: "0141224123",
		Email:  "y.proskurin@rto.de",
		Tasks:  []string{"buy milk", "read 30 pages", "wash clothes"},
	}
	executeTemplate(w, "user.gohtml", user)
}

func main() {

	r := chi.NewRouter()
	r.Get("/", userHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	http.ListenAndServe(":3000", r)
}
