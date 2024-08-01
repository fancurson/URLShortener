package view

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrl map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if val, ok := pathsToUrl[p]; ok {
			http.Redirect(w, r, val, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(file os.File, fallback http.Handler) (http.HandlerFunc, error) {
	pathsToUrls, err := DeserializingYaml(file)
	if err != nil {
		return nil, err
	}
	pathUrlMap := makingMap(pathsToUrls)
	return MapHandler(pathUrlMap, fallback), nil
}

type PathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func DeserializingYaml(yamlFile os.File) ([]PathUrl, error) {
	var pathsToUrls []PathUrl
	decoder := yaml.NewDecoder(&yamlFile)
	err := decoder.Decode(&pathsToUrls)
	if err != nil {
		return nil, fmt.Errorf("Decode yamlFile error(invalid data): %v", err)
	}
	return pathsToUrls, nil
}

func makingMap(pathsToUrls []PathUrl) map[string]string {
	pathUrlMap := make(map[string]string)
	for _, d := range pathsToUrls {
		pathUrlMap[d.Path] = d.Url
	}
	return pathUrlMap
}
