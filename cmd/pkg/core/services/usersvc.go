package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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