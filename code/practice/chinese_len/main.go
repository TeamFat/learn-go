package main

import (
	"fmt"
	"unicode/utf8"
)

//中英文的混合长度
// go run main.go
//result:7
func main() {
	var str = "Hello世界"
	fmt.Println(utf8.RuneCountInString(str)) // 7
}
