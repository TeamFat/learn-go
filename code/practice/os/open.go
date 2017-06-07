package main

import (
	"bufio"
	"fmt"
	"os"
)

//从多个文件中查找重复的行
//go run open.go demo.txt demo1.txt
//结果：4 hello world

//os.Open返回的第二个值是一个Go语言内置的error类型。如果这个error和内置值的nil（译注：相当于其它语言里的NULL）相等的话，说明文件被成功的打开了
//之后文件被读取，一直到文件的最后，文件的Close方法关闭该文件，并释放占用的一切资源。
//如果err的值不是nil的话，那说明在打开文件的时候出了某种错误。这种情况下，error类型的值会描述具体的问题。这里简单错误处理会在标准错误流中用Fprintf和%v来格式化该错误字符串。
//在Go语言里，函数和包级别的变量可以以任意的顺序被声明，并不影响其被调用。
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// fmt.Println(arg)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
