package main

import "fmt"


//不支持嵌套，重载和默认参数
//无需声明原型
//支持不定长变参
//支持多返回值
//支持命名返回参数
//支持匿名函数和闭包
//使用func定义函数，左大括号依旧不能另起一行

func test (x, y int, s string) (int, string) {  //需要定义返回参数的类型，可以返回多个数值
	n := x + y
	return n, s
}

//函数是第一类对象，可作为参数传递
func test1(fn func() int) int {  //参数可以传入函数
	return fn()
}

type FormatFunc func(s string, x, y int) string //定义函数类型

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

//变参，有且只有一个 只能在最后
func test2(s string,n ... int) string {
	var x int
	for _,i := range n {
		x += i
	}
	return fmt.Sprintf(s,x)
}

//命名返回参数允许defer延迟调用通过闭包读取和修改
//命名返回参数可看做与形参类似的局部变量，最后由return隐式返回
func add(x, y int) (z int) {
	defer func() {
		println(z)  //输出3
		z += 100
	}()
	z = x + y
	return
}

func main() {
	x,y := test(1, 2, "a")  //接收返回的值
	fmt.Println(x,y)

	s1 := test1(func() int { return 100 })

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)
	fmt.Println(s1,s2)

	println(test2("sum:%d", 1, 2, 3))

	println(add(1, 2))

	//匿名函数,做为变量
	fn := func() { println("hello world")}
	fn()
}
