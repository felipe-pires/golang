package main

import "fmt"

func main() {
	//maps
	//map[string] - tipo da chave map[]string - tipo dos valores
	usuario := map[string]string{
		"nome":      "felipe",
		"sobrenome": "pires",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	// deletar chave

	delete(usuario, "nome")

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"user": {
			"nome":      "felipe",
			"sobrenome": "pires",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["user"]["nome"])

	//adicionar nova chave

	usuario2["user2"] = map[string]string{
		"nome":      "teste",
		"sobrenome": "teste",
	}

	fmt.Println(usuario2)
}
