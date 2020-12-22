package app

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"solidArchitecture/cmd/db"
	"solidArchitecture/cmd/pkg/core/services"
)

type MainServer struct {
	 Db  *sql.DB
	 router *httprouter.Router
	 usvc *services.UserSvc
}



func NewMainServer(Db *sql.DB, router *httprouter.Router, usvc *services.UserSvc) *MainServer {
	return &MainServer{Db:Db, router:router, usvc: usvc}
}
func (server *MainServer) Start(Db *sql.DB) {
	err := db.DatabaseInit(Db)
	if err != nil {
		log.Fatal("Cannot create tables, error is:", err)
	}
	server.InitRoutes()
}
func (server *MainServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}
