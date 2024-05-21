package main

import (
	"fmt"
)

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma1 := somar(20, 30)
	var soma2 int8 = somar(40, 50)
	fmt.Println(soma1, soma2)

	var funcao = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := funcao("texto")
	fmt.Println(result)

	resultadoCalculosSoma, ResultadoCalculosSubtracao := calculosMatematicos(5, 6)
	fmt.Println(resultadoCalculosSoma, ResultadoCalculosSubtracao)

	resultadoCalculosSoma1, _ := calculosMatematicos(5, 6) // IGNORAR MAIS RETORNOS
	fmt.Println(resultadoCalculosSoma1)
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}
