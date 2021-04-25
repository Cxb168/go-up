package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-up/dao/srv"
	"go-up/utils"
	"log"
)

// MySQL配置
const (
	mysqlUrl = "localhost:3306"
	username = "root"
	pwd      = "qqby666"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/go_up_test", username, pwd, mysqlUrl))
	utils.CheckErr(err)
	err = db.Ping()
	utils.CheckErr(err)

	userSrv := srv.NewUserSrv(db)

	for i := 1; i < 4; i++ {
		user, has := userSrv.Get(i)
		if has {
			log.Printf("User id=%d, name=%s, age=%d", user.Id, user.Name, user.Age)
		}
	}
}
