package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

	mux.Handle(
    "/static/",
    http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))),
  )

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/app/index.html"))
	tmpl.Execute(w, "")
}
