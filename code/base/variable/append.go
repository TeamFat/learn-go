package main

import "fmt"

//向slice尾部添加数据，返回新的slice对象
func main() {
	s := make([]int, 0, 5)
	fmt.Printf("%p\n", &s)      //0xc42007c080

	s2 := append(s, 1)
	fmt.Printf("%p\n", &s2)     //0xc42007c0a0

	fmt.Println(s, s2)          //[] [1]

}