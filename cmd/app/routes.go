package app

import (
	"fmt"
	"log"
	"net/http"
)

func (server *MainServer) InitRoutes() {
	fmt.Println("Routes are init in localhost:8888")

	server.router.POST("/api/humo/login", server.LoginHandler)

	server.router.GET("/", server.AddUserHandler)

	err := http.ListenAndServe("localhost:8888", server)
	if err != nil {
		log.Fatal("Cannot listen and serve")
	}
}
