package app

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (server *MainServer) AddUserHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	//TODO:Something
	server.usvc.AddUser("ikrom", "ikrom")
	w.Write([]byte("hello"))
}
