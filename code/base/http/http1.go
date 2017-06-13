package main

import "net/http"

func main() {

	h := http.FileServer(http.Dir("./home"))
	http.ListenAndServe(":8888", h)

}