package main

import "fmt"

//数组是值类型，赋值和传参会复制整个数组，而不是指针
//数组长度必须是常量，且是类型的组成部分。[2]int和[3]int是不同类型
//支持”==”，”！=”操作符，因为内存总是被初始化过
//指针数组[n]T,数组指针[n]T

func test(x [2]int) {
	fmt.Printf("x:%p", &x)  //x:0xc42000e280    指针位置发生了改变，说明是重新复制了值
	x[1] = 1000
}

func main()  {
	a := [3]int{1,2}    //定义了数组有三个元素，未初始化元素值为0
	fmt.Println(a)      //[1 2 0]

	b := [...]int{1, 2, 3, 4, 5}    //通过初始化值确定数组长度
	fmt.Println(b)      //[1 2 3 4 5]

	c := [5]int{2:100, 4:200}   //指定了数组对应的键，索引从0开始以0的值进行补齐到指定的数组长度
	fmt.Println(c)      //[0 0 100 0 200]

	//多维数组
	d := [2][3]int{{1,2,3}, {4,5,6}}    //指定每一层的元素个数
	fmt.Println(d)      //[[1 2 3] [4 5 6]]

	e := [...][3]int{{1,2,3}, {4,5,6}}  //省略第一维的元素个数，第二维不能省略

	fmt.Println(e)      //[[1 2 3] [4 5 6]]

	//数组长度
	fmt.Println(len(a),cap(a))

	f := [2]int{1,2}
	fmt.Printf("f:%p", &f)  //f:0xc42000e270x
	test(f)
	fmt.Println(f)  //[1 2] 由于是值复制，因此当前的数组还是原来的值
}