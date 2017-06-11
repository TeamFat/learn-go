package main

import (
	"net/http"
)

//Cookie是由浏览器维持的，存储在客户端的一小段文本信息，伴随着用户请求和页面在Web服务器和浏览器之间传递。
//用户每次访问站点时，Web应用程序都可以读取cookie包含的信息。
//go run cookie.go
//read write delete cookie

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func ReadCookieServer(w http.ResponseWriter, req *http.Request) {
	// read cookie
	cookie, err := req.Cookie("testcookiename")

	//另一种读取方式
	//for _, cookie := range req.Cookies() {
	// fmt.Fprint(w, cookie.Name)
	// }
	if err == nil {
		cookievalue := cookie.Value
		w.Write([]byte("<b>cookie的值是：" + cookievalue + "</b>\n"))
	} else {
		w.Write([]byte("<b>读取出现错误：" + err.Error() + "</b>\n"))
	}
}

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/", MaxAge: 86400}
	http.SetCookie(w, &cookie)

	w.Write([]byte("<b>设置cookie成功。</b>\n"))
}

func DeleteCookieServer(w http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{Name: "testcookiename", Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)

	w.Write([]byte("<b>删除cookie成功。</b>\n"))
}

func main() {
	http.HandleFunc("/", SayHello)
	http.HandleFunc("/readcookie", ReadCookieServer)
	http.HandleFunc("/writecookie", WriteCookieServer)
	http.HandleFunc("/deletecookie", DeleteCookieServer)

	http.ListenAndServe("127.0.0.1:8000", nil)
}
