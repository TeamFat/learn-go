package main

import (
	"fmt"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello, world")
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.ListenAndServe(":9090", nil)
}
