package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// 数据库驱动。下划线为命名，因为不会直接将驱动拿来写代码，所以命名为下划线
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	user     = "root"
	password = "mysql"
	database = "go_db"
)

/**
* docker run --name go-web -e MYSQL_ROOT_PASSWORD=mysql -p 3306:3306 -d mysql
* docker exec -it go-web bash
 */

func main() {
	connectString := fmt.Sprintf("%s:%s@/%s", user, password, database)

	var _error error
	// Open 的第一个参数名为数据库驱动的名称
	db, _error = sql.Open("mysql", connectString)

	if _error != nil {
		log.Fatalln(_error.Error())
		return
	}

	_context := context.Background()

	// PingContext 用来验证与数据库的连接是否仍然有效
	_error = db.PingContext(_context)

	if _error != nil {
		log.Fatalln(_error.Error())
		return
	}

	fmt.Println("Connected!")

	// 查询复数条数据
	// list, _error := getMany(1)

	// 查询一条数据
	user, _ := getOne(2)

	user.name = "shamare"
	// _error = user.Update()

	user03 := customUser{
		id:   3,
		name: "suzuran",
	}

	_error = user03.Insert()

	if _error != nil {
		log.Fatalln(_error.Error())
		return
	}

	resutl, _ := getOne(3)
	fmt.Println(resutl)
}
