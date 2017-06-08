package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//net/http和io/ioutil包，http.Get函数是创建HTTP请求的函数，如果获取过程没有出错，那么会在resp这个结构体中得到访问的请求结果。
//resp的Body字段包括一个可读的服务器响应流。这之后ioutil.ReadAll函数从response中读取到全部内容；其结果保存在变量b中。
//resp.Body.Close这一句会关闭resp的Body流，防止资源泄露，Printf函数会将结果b写出到标准输出流中。
//res.Status输出状态码
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		fmt.Println("状态码：" + resp.Status)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
