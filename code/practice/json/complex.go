package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name  string
	Age   int
	Roles []string
	Skill map[string]float64
}

//map的key必须是字串，而value必须是同一类型的数据。
//run:go run complex.go
//result:{"Name":"rsj217","Age":27,"Roles":["Owner","Master"],"Skill":{"elixir":90,"python":99.5,"ruby":80}}

//总结
//1.首先定义需要编码的结构，然后调用encoding/json标准库的Marshal方法生成json byte数组，转换成string类型即可。
//2.golang和json的大部分数据结构匹配，对于复合结构，go可以借助结构体和空接口实现json的数组和对象结构。通过struct tag可以灵活的修改json编码的字段名和输出控制。
func main() {

	skill := make(map[string]float64)

	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0

	user := User{
		Name:  "rsj217",
		Age:   27,
		Roles: []string{"Owner", "Master"},
		Skill: skill,
	}

	rs, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rs))
}
