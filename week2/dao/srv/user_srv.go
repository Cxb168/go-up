package srv

import (
	"database/sql"
	models2 "go-up/week2/dao/models"
	"log"
)

type UserSrv struct {
	db *sql.DB
}

func NewUserSrv(db *sql.DB) *UserSrv {
	return &UserSrv{db: db}
}

func (srv UserSrv) Get(id int) (*models2.User, bool) {
	rtn := &models2.User{}
	row := srv.db.QueryRow(`SELECT * FROM user WHERE id=?`, id)

	err := row.Scan(&rtn.Id, &rtn.Name, &rtn.Age)
	if err == sql.ErrNoRows {
		log.Printf("Get User id=[%d] no exist", id)
		return nil, false
	} else if err != nil {
		log.Printf("Get User id=[%d] Error: %v", id, err)
		return nil, false
	}
	return rtn, true
}
