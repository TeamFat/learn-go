package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//一次性的把整个输入内容全部读到内存中，然后再把其分割为多行，然后再去处理这些行内的数据
//ReadFile函数返回byte类型的slice，这个slice必须被转换为string，之后才能够用strings.Split方法来进行处理
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
