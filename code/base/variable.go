package main

import "fmt"

func getPerson()(int,string) {
	return 20,"girl"
}

func main() {
	//定义没有初始值的变量，必须带类型
	var n int
	n = 1
	fmt.Println(n)

	//有初始值不必须定义类型 且必须用双引号
	var a = "a";
	fmt.Println(a);

	//函数内部可以使用:=赋值
	b := "b";
	fmt.Println(b);

	//一次性可以赋值多个变量
	var c,d = "c","d";
	fmt.Println(c,d);

	//特殊占用符_,请忽略。使用场景，局部变量如果没有被使用会当做错误
	var _,sex = getPerson();
	fmt.Println(sex);
}