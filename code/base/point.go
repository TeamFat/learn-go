package main

import "fmt"

func main() {
	//默认值为nil，没有NULL常量
	//操作符”&”取变量地址，”*”透过指针访问目标对象
	//不支持指针运算，不支持”->”运算符，直接用”.”访问目标成员
	type data struct {
		a int
	}
	s := data{1}
	var p *data
	p = &s
	fmt.Printf("%p,%d\n",p,p.a)
}