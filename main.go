package main

import (
	"URLShortener/view"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	yamlFile := flag.String("YamlFile", "resources/textYaml.yaml", "File with yaml data")
	flag.Parse()

	//////
	file, err := os.Open(*yamlFile)
	if err != nil {
		log.Fatalf("Open file error: %v", err)
	}
	defer file.Close()

	//////
	mux := defaultMux()

	pathsToUrl := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := view.MapHandler(pathsToUrl, mux)

	//////
	yamlHandler, err := view.YAMLHandler(*file, mapHandler)
	if err != nil {
		log.Fatalf("yaml error: %v", err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
