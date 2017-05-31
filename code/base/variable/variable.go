package variable

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
	//"_" 是一个特殊的变量名，任何赋予它的值都会被丢弃
	var _,sex = getPerson();
	fmt.Println(sex);

	//可以使用"+"操作符来链接两个字符串
	s := "hello"
	s = "c" + s[1:] //字符串切片操作
	fmt.Printf("%s\n", s)

	//要修改字符串，必须先将其转换成[]rune或[]byte,完成后再转换为string，这两种方式都会重新分配内存，并复制字节数组
	m := "hello";
	ms := []byte(m)
	ms[2] = 'a';
	fmt.Println(ms)
	fmt.Println(string(ms))
}