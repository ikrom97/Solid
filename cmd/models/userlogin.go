package models

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type ResponseLogin struct {
	Description string `json:"description"`
	Token       string `json:"token"`
	Role        string `json:"role"`
}
