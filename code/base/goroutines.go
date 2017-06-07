package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步
//会启动一个新的 Go 程并执行
func main() {
	go say("world")
	say("hello")
}
