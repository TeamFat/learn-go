package main

import (
	"fmt"
	"strconv"
)

// 示例
func main() {
	fmt.Println(strconv.ParseInt("FF", 16, 0))
	// 255
	fmt.Println(strconv.ParseInt("0xFF", 16, 0))
	// 0 strconv.ParseInt: parsing "0xFF": invalid syntax
	fmt.Println(strconv.ParseInt("0xFF", 0, 0))
	// 255
	fmt.Println(strconv.ParseInt("9", 10, 4))
	// 7 strconv.ParseInt: parsing "9": value out of range

	s := "Hello\t世界！\n"
	fmt.Println(s)                         // Hello	世界！（换行）
	fmt.Println(strconv.Quote(s))          // "Hello\t世界！\n"
	fmt.Println(strconv.QuoteToASCII(s))   // "Hello\t\u4e16\u754c\uff01\n"
	fmt.Println(strconv.QuoteToGraphic(s)) // "Hello\t世界！\n"
}
