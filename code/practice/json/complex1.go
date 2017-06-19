package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name    string
	Age     int
	Roles   []string
	Skill   map[string]float64
	Account Account
	Extra   []interface{}
}
type Account struct {
	Email    string  `json:"email"`
	Password string  `json:"pass_word"`
	Money    float64 `json:"money"`
}

//run:go run complex1.go
//result:{"Name":"rsj217","Age":27,"Roles":["Owner","Master"],"Skill":{"elixir":90,"python":99.5,"ruby":80},"Account":{"email":"rsj217@gmail.com","pass_word":"123456","money":100.5},"Extra":null}
//通过定义嵌套的结构体Account，实现了key与value不一样的结构。
//使用空接口，也可以定义像结构体实现那种不同value类型的字典结构。当空接口没有初始化其值的时候，零值是 nil。编码成json就是 null
func main() {
	skill := make(map[string]float64)

	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0

	account := Account{
		Email:    "rsj217@gmail.com",
		Password: "123456",
		Money:    100.5,
	}

	user := User{
		Name:    "rsj217",
		Age:     27,
		Roles:   []string{"Owner", "Master"},
		Skill:   skill,
		Account: account,
	}

	rs, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rs))

}
