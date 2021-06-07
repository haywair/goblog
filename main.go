package main

import (
	"fmt"
	"net/http"
)

func hanleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprintln(w, "<h1>this ia welcome page celebrate</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprintln(w, "<h2>yeah, this is a information page</h2>")
	} else if r.URL.Path == "/company" {
		fmt.Fprintln(w, "this is company information page")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "the page is not exist")
	}
}

func main() {
	http.HandleFunc("/", hanleFunc)
	http.ListenAndServe(":3000", nil)
}
