package main

import "fmt"

type usuario struct {
	nome  string
	email string
	idade int
}

func (u usuario) salvar() usuario {
	return u
}

func main() {
	user := usuario{
		nome:  "felipe",
		email: "felipe@gmail.com",
		idade: 20,
	}

	fmt.Println(user.salvar())
}
