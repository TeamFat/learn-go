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
	//& 操作符会生成一个指向其操作数的指针
	//* 操作符表示指针指向的底层值
	var i = 1;
	//类型 *T 是指向 T 类型值的指针。其零值为 nil 
	var p1 *int;
	fmt.Println(p1);
	p1 = &i
	fmt.Println(p1)
	fmt.Println(*p1)	
	
	var p *data
	p = &s
	//结构体字段可以通过结构体指针来访问
	//语言也允许我们使用隐式间接引用，直接写 p.a 就可以
	fmt.Printf("%p,%d\n",p,p.a)
}