package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	t.Run("Retangulo", func(t *testing.T) {
		retangulo := Retangulo{
			Largura: 12,
			Altura:  10,
		}
		areaEsperada := float64(120)
		areaRecebida := retangulo.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Area diferente da esperada")
		}
	})

	t.Run("circulo", func(t *testing.T) {
		circulo := Circulo{
			Raio: 10,
		}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circulo.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Area diferente da esperada")
		}
	})
}
