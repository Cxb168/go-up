package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:qqby666@tcp(localhost:3306)/test")
	checkErr(err)
	err = db.Ping()
	checkErr(err)
	row := db.QueryRow("SELECT * FROM `user` WHERE age=?", 18)
	var id int
	var name string
	var age int
	err = row.Scan(&id, &name, &age)
	checkErr(err)
	fmt.Println("---------", id, name, age)
}

func checkErr(err error) {
	if err != nil {
		//log.Println(err)
		panic(err)
	}
}
