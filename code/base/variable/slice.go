package main

import "fmt"

//slice不是数组或数组指针，它是通过内部指针和相关属性引用数组片段，以实现边长方案
//slice引用类型，但自身是结构体，值拷贝传递
/*
在runtime.h中的源码
struct slice
{
	byte* array;
	uintgo len;     //表示可用元素数量，读写操作不能超过该限制
	uintgo cap;     //表示最大扩张容量，不能超出数组限制
}
*/
func main() {
	//slice使用方式[a:b:c]，a表示从a位置开始，b表示到b位置结束，c表示最大位置是c，len=b-a,cap=c-a
	data := [...]int{0, 1, 2, 3, 4, 5, 6}
	slice := data[1:4:5]
	fmt.Println(slice)      //[1 2 3]

	//使用内建函数make初始化一个长度很大的值
	ips := make([]string,100)
	fmt.Println(ips)


}