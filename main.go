package main

import (
	"fmt"
	"net/http"
	"os"
)

func PrintRequestMiddleware(handlerfunc func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("%v\n", req)
		handlerfunc(w, req)
	}
}

func index(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, os.Getenv("ECHO"))
}

func indexV2(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, os.Getenv("ECHO2"))
}

func main() {
	http.HandleFunc("/", PrintRequestMiddleware(index))
	http.HandleFunc("/v2", PrintRequestMiddleware(indexV2))

	fmt.Println("http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
