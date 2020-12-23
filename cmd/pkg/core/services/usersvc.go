package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"solidArchitecture/cmd/models"
)

type UserSvc struct {
	Db *sql.DB
}

func NewUserSvc(Db *sql.DB) *UserSvc {
	if Db == nil {
		log.Println(errors.New("Db can't be nil"))
	}
	return &UserSvc{Db: Db}
}

func (receiver *UserSvc) AddUser(name, surname string) {
	fmt.Println(name, surname)
}
//TODO: func () - check has a user --> r.login & r.password
func (receiver *UserSvc) CheckHasUser(usr models.Login) (user, id string) {
	//TODO: get from db info by -- user
	usr.Login = user
	id = "2"
	return
}
//TODO: func () - create token