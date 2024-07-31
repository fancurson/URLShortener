package main

import (
	"URLShortener/view"
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()
	pathsToUrl := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := view.MapHandeler(pathsToUrl, mux)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/hello", anotherHello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func anotherHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello All")
}
