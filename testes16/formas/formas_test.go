package formas_test

import (
	"math"
	. "modulo/testes16/formas"
	"testing"
)

func TestAarea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEperada := float64(120)
		areaRecebida := ret.Area()

		if areaRecebida != areaEperada {
			// se o erro for fatal ele não continua
			t.Fatalf("Area recebida %f diferente area esperada %f", areaRecebida, areaEperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{10}

		areaEperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaRecebida != areaEperada {
			t.Fatalf("Area recebida %f diferente area esperada %f", areaRecebida, areaEperada)
		}
	})

}
