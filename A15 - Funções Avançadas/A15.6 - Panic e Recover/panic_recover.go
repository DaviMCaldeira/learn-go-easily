package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 int) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6!") // Mata a execução do programa
	// E para recuperarmos usaremos o recover
}

func alunoEstaAprovadoSemRecover(n1, n2 int) bool {
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6!") // Mata a execução do programa
	// E para recuperarmos usaremos o recover
}

func main() {
	// APROVADO
	fmt.Println(alunoEstaAprovado(7, 8))

	// REPROVADO
	fmt.Println(alunoEstaAprovado(5, 4))

	// PANIC SEM RECOVER
	fmt.Println(alunoEstaAprovadoSemRecover(6, 6))

	// PANIC COM RECOVER
	fmt.Println(alunoEstaAprovado(6, 6))
}
