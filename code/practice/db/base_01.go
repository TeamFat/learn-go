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

// https://github.com/go-sql-driver/mysql 支持database/sql，全部采用go写。
// https://github.com/ziutek/mymysql 支持database/sql，也支持自定义的接口，全部采用go写。

//在Go语言中，读取操作一般使用db.Query这个接口，而操作数据一般使用db.Exec这个接口。 因为在Go语言中，Exec方法返回的是sql.Result对象，而Query返回的是Rows对象。
//使用Scan方法时，你需要预先知道所取出来的数据有多少行。假如取出来的行数不定时，可以使用Columns。
//cols, err := rows.Columns()

/*通用方法 既不知道列数也不知道列类型
cols, err := rows.Columns() // Remember to check err afterwards
vals := make([]interface{}, len(cols))
for i, _ := range cols {
	vals[i] = new(sql.RawBytes)
}
for rows.Next() {
	err = rows.Scan(vals...)
	// Now you can check each element of vals for nil-ness,
	// and you can use type introspection and type assertions
	// to fetch the column into a typed variable.
}
*/

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

	//1.直接查询 以下是注意点：
	//使用完查询结果后，要注意关闭rows.Close()
	//读取查询结果时使用的是Scan方法
	//使用完rows之后要注意检查是否有错误rows.Err()
	//rows.Close()这个方法是一个无害的方法，可以重复调用。 但是，在使用这个方法前，必须先检查前面操作有没有错误产生。
	rows, err := db.Query("select id,title from posts where id < ?", 10)
	if err != nil {
		log.Fatal(err)
	}
	var (
		id    int
		title string
	)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, title)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
