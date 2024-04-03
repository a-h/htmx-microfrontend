package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/htmx-microfrontend/routes/home"
	"github.com/a-h/htmx-microfrontend/routes/quoteevents"
)

func main() {
	mux := http.NewServeMux()

	homeHandler := home.NewHandler()
	mux.Handle("/", homeHandler)

	quoteeventsHandler := quoteevents.NewHandler()
	mux.Handle("/quoteevents", quoteeventsHandler)
	mux.Handle("/quoteevents/", quoteeventsHandler)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Listening on :8080")
	http.ListenAndServe("localhost:8080", mux)
}
