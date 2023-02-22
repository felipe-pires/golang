package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavra := strings.Split(strings.ToLower(endereco), " ")[0]

	verifica := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			verifica = true
		}
	}

	if verifica {
		return strings.Title(primeiraPalavra)
	}
	return "Invalido"
}
