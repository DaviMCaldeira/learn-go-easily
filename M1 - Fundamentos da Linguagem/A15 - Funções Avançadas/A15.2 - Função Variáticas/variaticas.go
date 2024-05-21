package main

import "fmt"

// PASSAR N PARAMETROS

func soma(numeros ...int) int {

	soma := 0

	for _, num := range numeros {
		soma += num
	}

	return soma
}

func somaForSlices(numeros []int) int {

	soma := 0

	for _, num := range numeros {
		soma += num
	}

	return soma
}

// PASSAR 1 PARAMETRO + N PARAMETROS

func escreverSoma(texto string, numeros ...int) {
	total := somaForSlices(numeros)

	fmt.Println(texto, "; Soma=", total)

}

func main() {
	totalSoma1 := soma()
	totalSoma2 := soma(5, 5, 5)
	totalSoma3 := soma(6234, 432423, 243423, 4234)
	totalSoma4 := soma(432423, 4234232, 4324234324324, 324123123)

	fmt.Println(totalSoma1, totalSoma2, totalSoma3, totalSoma4)

	escreverSoma("Oi mundo", 20, 30, 40, 50, 60)
}
