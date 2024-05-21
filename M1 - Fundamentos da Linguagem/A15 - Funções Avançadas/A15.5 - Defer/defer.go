package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func AlunoEstaAprovado(n1, n2 int) bool {
	defer fmt.Println("Média Calculada. Resultado será retornado!")
	fmt.Println("Verificando se o aluno está aprovado...")
	media := (n1 + n2) / 2

	if media > 7 {
		return true
	} else {
		return false
	}
}

func main() {
	defer funcao1()
	// DEFER = ADIAR, ou seja, adiar a funcao até o último momento, sempre executado antes de um return, como é mostrado na funcao AlunoEstaAprovado
	// Útil com banco de dados, fechar conexão com banco de dados, etc.
	funcao2()

	fmt.Println(AlunoEstaAprovado(10, 8))
}
