package app

import (
	"fmt"
	"net/http"
)

func (server *MainServer) InitRoutes() {
	fmt.Println("Routes are init in localhost:8888")

	server.router.GET("/", server.AddUserHandler)

	http.ListenAndServe("localhost:8888", server)
}
