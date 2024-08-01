package view

import (
	"fmt"
	"net/http"

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

func YAMLHandler(yamlBytes string, fallback http.Handler) (http.HandlerFunc, error) {

	pathsToUrls := mustParsing(DeserializingYaml(yamlBytes))
	pathUrlMap := makingMap(pathsToUrls)
	return MapHandler(pathUrlMap, fallback), nil
}

type PathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func DeserializingYaml(yamlString string) ([]PathUrl, error) {
	var pathsToUrls []PathUrl
	err := yaml.Unmarshal([]byte(yamlString), &pathsToUrls)
	if err != nil {
		return nil, fmt.Errorf("yaml parsing error: %v", err)
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

func mustParsing(res []PathUrl, err error) []PathUrl {
	if err != nil {
		panic(err)
	}
	return res
}
