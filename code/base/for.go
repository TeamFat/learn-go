package main

import (
	"fmt"
)

func main() {
	//for循环的2种方式
	s := "hello"
	for i := 0; i<len(s);i++  {
		fmt.Printf("%c",s[i]) //byte方式
	}

	for _, r := range s {
		fmt.Printf("%c",r) //rune方式
	}

	//while的for循环方式
	n := len(s)
	for n > 0 {
		println(n)
		n--
	}
}

