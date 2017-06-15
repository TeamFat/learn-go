package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//使用mysql数据库的增删改查（CURD）功能
type Post struct {
	Id      int
	Content string
	Author  string
}

//插入
func (post *Post) Create() (err error) {
	rs, err := Db.Exec("INSERT INTO posts (content, author) Values (?, ?)", post.Content, post.Author)
	if err != nil {
		log.Fatalln(err)
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id)
	return
}

//删除
func (post *Post) Delete() (err error) {
	rs, err := Db.Exec("DELETE FROM posts WHERE id=?", post.Id)
	if err != nil {
		log.Fatalln(err)
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rows)
	return
}

//修改
func (post *Post) Update() (err error) {
	rs, err := Db.Exec("UPDATE posts SET author=? WHERE id=?", post.Author, post.Id)
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(rows)
	return
}

//查询
func RetrievePost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("SELECT id, content, author FROM posts WHERE id=?", id).Scan(
		&post.Id, &post.Content, &post.Author)
	return
}

var Db *sql.DB

func main() {
	var err error
	Db, err = sql.Open("mysql", "root:mysql123456@tcp(127.0.0.1:3306)/beego?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	post := Post{
		Content: "hello world",
		Author:  "vanyarpy",
	}

	post.Create()
	query, err := RetrievePost(1)
	fmt.Println(query)
	defer Db.Close()
}
