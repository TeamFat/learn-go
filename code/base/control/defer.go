package main

import (
	"fmt"
)

//延迟函数调用的执行顺序，会与其所属的defer语句执行顺序相反
func printNumbers() {
	for i:=0;i<5;i++{
		defer func(n int){
			fmt.Printf("%d",n)
		}(i)
	}
}

func main() {
	printNumbers();
}
