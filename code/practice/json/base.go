package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// json源于javascript的对象结构，golang中直接对应的数据结构，可是golang的map也是key-value结构，同时struct结构体也可以描述json。
//当然，对于json的数据类型，go也会有对象的结构所匹配。关系如下：

// 数据类型	JSON	Golang
// 字串	string	string
// 整数	number	int64
// 浮点数	number	flaot64
// 数组	arrary	slice
// 对象	object	struct
// 布尔	bool	bool
// 空值	null	nil

// golang提供了encoding/json的标准库用于编码json。大致需要两步：

// 首先定义json结构体。
// 使用 Marshal方法序列化。

type Account struct {
	Email    string
	password string
	Money    float64
}

//小写命名的password字段没有被编码到json当中，生成的json结构字段和Account结构一致
//run:go run base.go
//result:{"Email":"rsj217@gmail.com","Money":100.5}
func main() {
	account := Account{
		Email:    "rsj217@gmail.com",
		password: "123456",
		Money:    100.5,
	}

	rs, err := json.Marshal(account)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(rs)
	fmt.Println(string(rs))
}
