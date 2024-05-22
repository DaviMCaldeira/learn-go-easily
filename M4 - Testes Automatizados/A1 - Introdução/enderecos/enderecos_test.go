package enderecos

import (
	"testing"
)

func TestTipoDeEndereco(t *testing.T) {
	// Assinatura deve começar com TestXxx...
	enderecoParaTeste := "Avenida Paulista de São Paulo"

	resultadoEsperado := "Avenida"
	resultadoRecebido := TipoDeEndereco(enderecoParaTeste)

	if resultadoEsperado != resultadoRecebido {
		t.Error("O tipo recebido é diferente do esperado!")
	}
}

// Para rodar usar o "go test"

// TESTE DE UNIDADE

// OBS: Os arquivos que serão testados devem ter (qualquer coisa)_test.go
