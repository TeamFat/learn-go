package main

import "fmt"

//常量可以定义为数值、布尔值或字符串等类型
const Pi float32 = 3.1415926

const a int = 1
const b = "b"
const c = false

func main() {
	fmt.Println(Pi);

	//在常量数组中，如不提供类型和初始值，则视作与上一个常量相同
	const (
		d = 2
		x
	)
	println(x)  //2
}
