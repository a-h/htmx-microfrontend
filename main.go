package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/htmx-microfrontend/routes/home"
	"github.com/a-h/htmx-microfrontend/routes/result"
)

func main() {
	mux := http.NewServeMux()

	homeHandler := home.NewHandler()
	mux.Handle("/", homeHandler)

	resultHandler := result.NewHandler()
	mux.Handle("/result", resultHandler)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Listening on :8080")
	http.ListenAndServe("localhost:8080", mux)
}
