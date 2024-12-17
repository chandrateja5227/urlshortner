package main

import (
	"fmt"
	"net/http"
	"urlshortner/urlshort"
)

func main() {
	mux := defaultMux()


	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/chandrateja5227/urlshortner
- path: /urlshort-readme
  url: https://github.com/chandrateja5227/urlshortner/blob/main/README.md
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mux)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	//mux.HandleFunc("/", hello)
	mux.HandleFunc("/NotFound", NotFound)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "404! Page Not Found")
}
