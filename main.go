package main

import (
	"fmt"
	"net/http"
)

func hanleFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintln(w, "this ia welcome page celebrate")
	} else if r.URL.Path == "/about" {
		fmt.Fprintln(w, "yeah, this is a information page")
	} else if r.URL.Path == "/company" {
		fmt.Fprintln(w, "this is company information page")
	}
}

func main() {
	http.HandleFunc("/", hanleFunc)
	http.ListenAndServe(":3000", nil)
}
