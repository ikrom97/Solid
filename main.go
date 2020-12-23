package main

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	_"github.com/mattn/go-sqlite3"
	"log"
	"solidArchitecture/cmd/app"
	"solidArchitecture/cmd/pkg/core/services"
)
func main() {
	DB, err := sql.Open("sqlite3", "bank")
	if err != nil {
		log.Fatal("Can't find sql connection", err)
	}
	router := httprouter.New()
	svc := services.NewUserSvc(DB)
	server := app.NewMainServer(DB, router, svc)



	server.Start(DB)
}


