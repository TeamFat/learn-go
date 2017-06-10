package main

import (
	"database/sql"
	"log"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
CREATE TABLE `posts` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `title` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
 `author` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
 `content` text COLLATE utf8_unicode_ci NOT NULL,
 `created_at` timestamp NULL DEFAULT NULL,
 `updated_at` timestamp NULL DEFAULT NULL,
 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15335 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci
*/

//访问数据库是通过sql包中的DB对象实现的。
//sql.DB并不是一个数据库连接！它也不对应于数据库理论中的任一概念，其只是一个接口的抽象以及数据库的表示。 数据库可以是一个本地文件、一个网络链接或者内存、进程数据程。
//通过数据库驱动打开或者关闭与底层数据库的连接
//若有需要，它会去管理维护一个数据库连接池

// 1.直接使用查询语句进行查询
// 2.使用预声明进行多次查询
// 3.使用预声明进行单次查询
// 4.只获取单行数据的查询

func main() {
	db, err := sql.Open("mysql", "root:mysql123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		// do something here
		fmt.Println("db not ping")
	}

	//3.使用预声明进行单次查询 以下是注意点：
	stmt, err := db.Prepare("select title from posts where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var title string
	err = stmt.QueryRow(1).Scan(&title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(title)
}
