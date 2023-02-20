package main

import (
	"fmt"
)

// utilizar a "heranca" em go

type pessoa struct {
	nome  string
	idade int
}

type estudante struct {
	pessoa
	matricula string
}

func main() {

	// pessoa := pessoa {
	// 	nome: "felipe",
	// 	idade: 20,
	// }

	p := pessoa{
		nome:  "teste",
		idade: 200,
	}

	e := estudante{
		p,
		"123456",
	}
	fmt.Println(e, e.nome, e.matricula)

	aluno := estudante{
		pessoa{
			nome:  "felipe",
			idade: 20,
		},
		"123456",
	}
	fmt.Println(aluno, aluno.nome)
}
