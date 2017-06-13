package main

import (
	"net/http"
	"log"
	"io"
)

type MyHandler struct {}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &MyHandler{})
	mux.HandleFunc("/home", Home)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("listen falded")
	}
}
func (this *MyHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world " + r.URL.Path)
}
func Home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "home " + r.URL.Path)
}