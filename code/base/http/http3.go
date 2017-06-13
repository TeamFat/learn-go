package main


import (
	"net/http"
	"log"
	"io"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)
type MyHandler struct{}

func main() {
	server := http.Server{
		Addr: ":8080",
		// 定义server的默认handler
		Handler: &MyHandler{},
		ReadTimeout: 2 * time.Second,
	}
	// 对全局变量赋值
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/home"] = Home
	mux["/user"] = User
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("listen faild")
	}
}
// handler默认执行的方法
func (*MyHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // 根据map转发路由
	if v, ok := mux[r.URL.String()]; ok {
		v(w, r)
		return
	}
	io.WriteString(w, "hello world " + r.URL.Path)
}
func Home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "home " + r.URL.Path)
}
func User(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "user " + r.URL.Path)
}