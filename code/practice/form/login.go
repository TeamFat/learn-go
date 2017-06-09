package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

//go run login.go
//chrome open http://127.0.0.1:9090/login
//username: <script>fsfs</script>  passport:123456
//可以看到防xss的方式

//login.gtpl 登录页面
//login.gtpl表单会提交到服务器的/login，当用户输入信息点击登陆之后，会跳转到服务器的路由login里面
//补充：如req.Form["username"]也可写成req.FormValue("username")。 调用req.FormValue时会自动调用req.ParseForm，所以不必提前调用。
//req.FormValue只会返回同名参数中的第一个，若参数不存在则返回空字符串。

func sayhelloName(res http.ResponseWriter, req *http.Request) {

	//解析url传递的参数，对于POST则解析响应包的主体（request body）
	req.ParseForm()

	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(req.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", req.URL.Path)
	fmt.Println("scheme", req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(res, "Hello astaxie!") //这个写入到w的是输出到客户端的

}

func login(res http.ResponseWriter, req *http.Request) {
	fmt.Println("method:", req.Method) //获取请求的方法
	if req.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(res, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		//req.Form里面包含了所有请求的参数，比如URL中query-string、POST的数据、PUT的数据，
		//所有当你在URL的query-string字段和POST冲突时，会保存成一个slice，里面存储了多个值
		req.ParseForm()
		fmt.Println("username:", req.Form["username"])
		fmt.Println("password:", req.Form["password"])
		//防止xss攻击，能转义
		fmt.Println("username:", template.HTMLEscapeString(req.Form.Get("username")))
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
