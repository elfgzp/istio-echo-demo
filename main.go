package main

import (
	"os"
    "fmt"
    "net/http"
)

func index(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, os.Getenv("ECHO"))
}

func indexV2(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, os.Getenv("ECHO2"))
}


func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/v2", indexV2)

	fmt.Println("http://127.0.0.1:8088")
    http.ListenAndServe(":8088", nil)
}