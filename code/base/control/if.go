package control

import "fmt"

func main() {
	//可省略条件表达式括号
	//支持初始化语句，可定义代码块局部变量
	//代码块左大括号必须在条件表达式尾部
	x := 0
	if n := "abc"; x>10 {
		fmt.Println(n[1])
	} else if x <0 {
		fmt.Println(n[2])
	} else {
		fmt.Printf("%c",n[0])
	}

	//
}