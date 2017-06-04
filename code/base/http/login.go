package http

import (

	"fmt"
	"html/template"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r* http.Request) {

	fmt.Fprintf(w, "hello, world")
}

func login(w http.ResponseWriter, r* http.Request) {

	if r.Method == "GET" {

		t, _ := template.ParseFiles("login.gtpl");
		t.Execute(w, nil)
	} else {

		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password", r.Form["password"])

	}

}

func main() {

	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9090", nil)
}