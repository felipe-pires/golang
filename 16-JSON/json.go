package main

import (
	"bytes"
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

	user := usuario{
		Nome:  "Felipe",
		Email: "felipe@gmail.com",
		Idade: 20,
		Senha: "123456",
	}

	json1, erro := json.Marshal(user)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(json1)
	fmt.Println(bytes.NewBuffer((json1)))

	user2 := map[string]string{
		"nome":  "felipe",
		"email": "felipe@gamil.com",
	}

	json2, erro2 := json.Marshal(user2)

	if erro2 != nil {
		log.Fatal(erro2)
	}

	fmt.Println(json2)
	fmt.Println(bytes.NewBuffer(json2))

}
