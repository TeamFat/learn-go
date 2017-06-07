package main

import (
	"bufio"
	"fmt"
	"os"
)

//终端情况下请使用 ctrl+d
//go run stdin 换行输入 aa bb aa cc aa
//结果 3 aa

////程序每次读取输入的一行，这一行的内容会被当做一个map的key，而其value值会被+1
//counts[input.Text()]++这个语句和下面的两句是等价的
//line := input.Text()
//counts[line] = counts[line] + 1

/*格式化
%d          int变量
%x, %o, %b  分别为16进制，8进制，2进制形式的int
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔变量：true 或 false
%c          rune (Unicode码点)，Go语言里特有的Unicode字符类型
%s          string
%q          带双引号的字符串 "abc" 或 带单引号的 rune 'c'
%v          会将任意变量以易读的形式打印出来
%T          打印变量的类型
%%          字符型百分比标志（%符号本身，没有其他操作）
*/
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// fmt.Println(input.Text())
		counts[input.Text()]++
	}
	for line, n := range counts {
		fmt.Println(line)
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
