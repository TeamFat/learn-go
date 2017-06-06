package main

import (
	"fmt"
)

//没有do和while的循环，可以用for来替换
func main() {
	//for循环的2种方式
	s := "hello"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i]) //byte方式
	}

	//当使用 for 循环遍历切片时，每次迭代都会返回两个值。 第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
	//可以将下标或值赋予 _ 来忽略它
	for _, r := range s {
		fmt.Printf("%c", r) //rune方式
	}

	//while的for循环方式
	n := len(s)
	for n > 0 {
		println(n)
		n--
	}
}
