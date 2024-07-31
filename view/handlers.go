package view

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandeler(pathsToUrl map[string]string, fallback http.Handler) http.HandlerFunc {
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
	var pathsToUrls []PathUrl
	err := yaml.Unmarshal([]byte(yamlBytes), &pathsToUrls)
	if err != nil {
		return nil, err
	}

	pathUrlMap := make(map[string]string)
	for _, d := range pathsToUrls {
		pathUrlMap[d.Path] = d.Url
	}

	return MapHandeler(pathUrlMap, fallback), nil
}

type PathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
