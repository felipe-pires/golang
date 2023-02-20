package main

import (
	"fmt"
)

// são as "classes" em go
type endereco struct {
	logradouro string
	numero     int
}

type usuario struct {
	nome  string
	email string
	idade int
	senha string
	endereco
}

var u usuario

func main() {

	user := usuario{
		nome:  "felipe",
		email: "felipe@gmail.copm",
		idade: 20,
		senha: "123456",
		endereco: endereco{
			logradouro: "rua",
			numero:     5,
		},
	}
	fmt.Println(user)

	u = usuario{
		nome:  "teste",
		email: "teste@gmail.copm",
		idade: 20,
		senha: "123456",
	}
	fmt.Println(u)

	var user1 usuario

	user1.nome = "user"
	user1.email = "user@user.com"
	user1.idade = 100
	user1.senha = "123456"

	fmt.Println(user1)
}
