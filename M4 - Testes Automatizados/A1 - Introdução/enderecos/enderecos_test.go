package enderecos

// package enderecos_test

// Normalmente podemos apenas ter 1 pacote por pasta, a exceção são os testes, ao invés de usar o mesmo nome, você pode usar enderecos_test, tem que ter o prefixo da pasta

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// Assinatura deve começar com TestXxx...
	enderecoParaTeste := "Avenida Paulista de São Paulo"

	resultadoEsperado := "Avenida"
	resultadoRecebido := TipoDeEndereco(enderecoParaTeste)

	t.Parallel() // Para rodar em Paralelo

	if resultadoEsperado != resultadoRecebido {
		t.Errorf("O tipo recebido é diferente do esperado!; Esperado = %s; Recebido = %s",
			resultadoEsperado,
			resultadoRecebido)
	}
}

func TestTipoDeEnderecoComVariosCenarios(t *testing.T) {

	t.Parallel() // Para rodar em Paralelo

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia X", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"null", "Tipo Inválido"},
		{"Estrada 69", "Estrada"},
		{"Rua dos Matrizes", "Rua"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado!; Esperado = %s; Recebido = %s",
				cenario.retornoEsperado,
				tipoDeEnderecoRecebido)
		}
	}

}

// Para rodar usar o "go test"

// Para rodar testes de um projeto, pasta raiz e executar go test ./...

// Para executar os testes no modo verboso, pode ser passado da seguinte maneira: go test -v, irá mostrar quais funções estão sendo executadas

// Execução de testes em Paralelo (t.Parallel())

// Ver a cobertura dos teste: go test --cover (Basicamente é ver se o teste cobre tudo que é possível da função pode retornar)
// Se removessemos os casos que retornam "Tipo Inválido" isso faria que ao invés de ser 100% fosse para uma porcentagem menor

// Para ver com mais detalhes a cobertura, primeiro vamos gerar um relatório, utilize: go test --coverprofile cobertura.txt (qualquer nome)
// Para fazer a leitura, utilize: go tool cover --func=cobertura.txt

// Outra ferramenta para visualizarmos: go tool cover --html=cobertura.html

// TESTE DE UNIDADE

// OBS: Os arquivos que serão testados devem ter (qualquer coisa)_test.go
