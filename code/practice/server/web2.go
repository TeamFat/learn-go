package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

//对/count这个url的请求会调用到count这个函数，其它所有的url都会调用默认的处理函数。
//服务器每一次接收请求处理时都会另起一个goroutine，这样服务器就可以同一时间处理多数请求
//注意并发
//这也就是代码里的mu.Lock()和mu.Unlock()调用将修改count的所有行为包在中间的目的
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
