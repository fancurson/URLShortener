package view

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func makingMap(pathsToUrls []PathUrl) map[string]string {
	pathUrlMap := make(map[string]string)
	for _, d := range pathsToUrls {
		pathUrlMap[d.Path] = d.Url
	}
	return pathUrlMap
}

type PathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

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

func FileHandler(file os.File, fallback http.Handler) (http.HandlerFunc, error) {
	ext := filepath.Ext(file.Name())
	pathsToUrls, err := FileDeserializing(file, ext)
	if err != nil {
		return nil, err
	}
	pathUrlMap := makingMap(pathsToUrls)
	return MapHandler(pathUrlMap, fallback), nil
}

func FileDeserializing(file os.File, ext string) ([]PathUrl, error) {
	var PathsUrls []PathUrl
	if ext == ".yaml" || ext == ".yml" {
		decoder := yaml.NewDecoder(&file)
		if err := decoder.Decode(&PathsUrls); err != nil {
			return nil, fmt.Errorf("Decode yamlFile error(invalid data): %v", err)
		}
	} else if ext == ".json" {
		decoder := json.NewDecoder(&file)
		if err := decoder.Decode(&PathsUrls); err != nil {
			return nil, fmt.Errorf("Decode jsonFile error(invalid data): %v", err)
		}
	}
	return PathsUrls, nil
}
