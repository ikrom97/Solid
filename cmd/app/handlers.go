package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"solidArchitecture/cmd/models"
	"solidArchitecture/cmd/token"
)

func (server *MainServer) AddUserHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	//TODO:Something
	server.usvc.AddUser("ikrom", "ikrom")
	w.Write([]byte("hello"))
}
func (server *MainServer) LoginHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	w.Header().Set("Content-Type", "application/json: charset=utf-8")
	var requestBody models.Login
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json_invalid")
		if err != nil {
			log.Println("Can't find connection")
			return
		}
		return
	}
	user, id := server.usvc.CheckHasUser(requestBody)
	createToken := token.CreateToken(id, user)
	responseBody := models.ResponseLogin{
		Description: "Description",
		Token:       createToken,
		Role:        "Admin",
	}

	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		log.Println("Can't find connection")
		return
	}
	//TODO: func (models.Login) - check has a user --> r.login & r.password
	//TODO: func () - create token
}