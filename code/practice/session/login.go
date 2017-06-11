package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
)

//session 用户登录
var globalSessions *session.Manager

//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func login(res http.ResponseWriter, req *http.Request) {
	sess := globalSessions.SessionStart(res, req)
	res.Header().Set("Content-Type", "text/html")
	req.ParseForm()
	if req.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		res.Header().Set("Content-Type", "text/html")
		t.Execute(res, sess.Get("username"))
	} else {
		fmt.Println(req.Form["username"])
		sess.Set("username", req.Form["username"])
		sessionval := sess.Get("username")
		if sessionval != nil {
			fmt.Println(sessionval)
		}
		http.Redirect(res, req, "/", 302)
	}
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello, world")
}

func main() {
	http.HandleFunc("/", sayHelloName)                // each request calls handler
	http.HandleFunc("/login", login)                  //设置访问的路由
	err := http.ListenAndServe("127.0.0.1:9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
