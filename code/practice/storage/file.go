package main

import (
	"fmt"
	"io/ioutil"
)

//golang提供了ioutil库用于读写文件，也提供了os相关的文件创建，写入，保存的工具函数。

func main() {
	data := []byte("Hello World!\n")
	fmt.Println(data)
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}

	read1, _ := ioutil.ReadFile("data1")
	fmt.Println(string(read1))
}
