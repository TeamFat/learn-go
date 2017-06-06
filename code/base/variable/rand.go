package main

import (
	"fmt"
	"math/rand"
	"time"
)

//不设置时间种子的话，每次生成的rand值相同
// 设置时间种子使用time包 
// 生成随机数需要math/rand包 
// 打印输出使用fmt包
func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("My favorite number is", rand.Intn(10))
}