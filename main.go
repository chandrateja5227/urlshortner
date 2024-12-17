package main

import (
	"fmt"
	"net/http"
	"os"
	"urlshortner/urlshort"
)

func main() {
	mux := defaultMux()
	ymlfile, err := os.ReadFile("Mapping.yml")
	if err != nil {
		panic(err)
	}

	yamlHandler, err := urlshort.YAMLHandler(ymlfile, mux)
	if err != nil {
		panic(err)
	}

	/*jsonfile , err :=os.ReadFile("Mapping.json")
	if err != nil {
		panic(err)
	}

	jsonHandler, err := urlshort.JSONHandler(jsonfile, mux)
	if err != nil {
		panic(err)
	}*/

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
