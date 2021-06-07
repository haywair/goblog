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

func aboutHanlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系")
}

func main() {
	http.HandleFunc("/", hanleFunc)
	http.HandleFunc("/how", aboutHanlerFunc)
	http.ListenAndServe(":3000", nil)
}
