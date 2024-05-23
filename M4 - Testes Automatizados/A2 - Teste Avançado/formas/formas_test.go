package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	/*

		A função t.Run em testes unitários em Go permite criar subtestes dentro de um teste principal.

		Cada subteste é executado de forma independente, facilitando a organização e a identificação de falhas
		específicas em diferentes cenários.

	*/

	// TDD - Test Driven Development (Desenvolvimento Guiado a Teste) - Fazer os testes antes de implementar

	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{20.0, 30.0}

		areaEsperada := float64(600.0)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("[ERROR] Esperado:%f; Recebido:%f", areaEsperada, areaRecebida)
			// Fatal != Error, o Fatal ele quebra a execução dos testes, enquanto o Error não.
		}

	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10.0}

		areaEsperada := float64(math.Pi * math.Pow(10.0, 2))
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("[ERROR] Esperado:%f; Recebido:%f", areaEsperada, areaRecebida)
			// Fatal != Error, o Fatal ele quebra a execução dos testes, enquanto o Error não.
		}
	})
}
