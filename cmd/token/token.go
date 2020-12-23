package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)
type MyCustomClaims struct {
	Login string
	jwt.StandardClaims
}
func CreateToken(id, login string) string {
	mySigningKey := []byte("AllYourBase")
	claims := MyCustomClaims{
		Login:          login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:time.Now().Unix() +300,
			Issuer: "test",
			Id:id,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println("Devnai")

	}
	fmt.Printf("I am a token = %v\n", ss)
	return ss
}