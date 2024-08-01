package main

import (
	"URLShortener/view"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const DEFAULFILE = "resources/textYAML.yaml"

func main() {

	// flags
	customFile := flag.String("fileName", DEFAULFILE, "File with YAML data")
	flag.Parse()

	// Default handler(fallback)
	mux := defaultMux()
	pathsToUrl := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := view.MapHandler(pathsToUrl, mux)

	//Work with all types of files(yaml, json, yml)
	fileHandler := WorkWithFile(*customFile, mapHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", fileHandler)
}

func WorkWithFile(fileName string, fallback http.HandlerFunc) http.HandlerFunc {

	customFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Open file error: %v", err)
	}
	defer customFile.Close()

	FileHandler, err := view.FileHandler(*customFile, fallback)
	if err != nil {
		log.Fatalf("Work with files error:", err)
	}
	return FileHandler
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
