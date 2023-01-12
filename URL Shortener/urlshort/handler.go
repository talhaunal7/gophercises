package urlshort

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"net/http"
)

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func buildMap(urls []pathUrl) map[string]string {
	m := make(map[string]string)
	for _, pu := range urls {
		m[pu.Path] = pu.URL
	}
	return m

}
func parseYAML(yml []byte) (map[string]string, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		fmt.Println(err)
	}
	pathUrlMap := buildMap(pathUrls)

	return pathUrlMap, err
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		for key := range pathsToUrls {
			if r.URL.Path == key {
				http.Redirect(w, r, pathsToUrls[key], http.StatusSeeOther)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}

}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	return MapHandler(parsedYaml, fallback), nil

}
