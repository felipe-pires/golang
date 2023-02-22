package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type usuario struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Idade int    `json:"idade"`
	Senha string `json:"senha"`
}

func main() {
	userJson := `{"nome":"Felipe","email":"felipe@gmail.com","idade":20,"senha":"123456"}`

	var user usuario

	if erro := json.Unmarshal([]byte(userJson), &user); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(user)
}
