package main

import "fmt"

// Retornos nomeados é que o desenvolvedor irá retornar
// variáveis

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2 // Não precisa declarar visto que está no retorno
	subtracao = n1 - n2
	return
}

func main() {
	soma, sub := calculosMatematicos(10, 20)
	fmt.Println(soma, sub)
}
