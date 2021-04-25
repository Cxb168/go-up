package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-up/dao/srv"
	"go-up/utils"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:qqby666@tcp(localhost:3306)/test")
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
