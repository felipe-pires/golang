package enderecos

import "testing"

type cenario struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// t.Parallel() executar testes em paralelo

	cenarioDeTeste := []cenario{
		{"Rua teste", "Rua"},
		{"Avenida teste", "Avenida"},
		{"Rodovia teste", "Rodovia"},
		{"Estrada teste", "Estrada"},
		{"Praca teste", "Invalido"},
		{"", "Invalido"},
	}

	for _, cenario := range cenarioDeTeste {
		TipoDeEndereco := TipoDeEndereco(cenario.enderecoInserido)

		if TipoDeEndereco != cenario.retornoEsperado {
			t.Error("O tipo recebido é diferente do esperado!")
		}
	}
}
