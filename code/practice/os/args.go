package main

import (
	"fmt"
	"os"
	"strings"
)

//命令行参数
//运行 go run args.go hello world
//输出 hello world
func main() {
	var s string
	var sep = ""
	for i := 1; i < len(os.Args); i++ {
		// fmt.Println(os.Args[i])
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//
	sep = ""
	var s1 string
	for _, arg := range os.Args[1:] {
		s1 += sep + arg
		sep = " "
	}
	fmt.Println(s1)

	//如果不断连接的字符串数量众多，那么上面这种操作就是成本非常高的操作。更简单并且有效的一种方式是使用strings包提供的Join函数
	fmt.Println(strings.Join(os.Args[1:], " "))
}
