package main

import "fmt"

func closure() func() {
	texto := "Dentro da função CLOSURE"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao

	// Closure vai ser essa função que retorna outra função com uma memória interna
	// Durante a execução deste arquivo será possível ver que a variável texto na func closure será diferente da que está na main

}

func main() {
	texto := "Dentro da função MAIN"
	fmt.Println(texto)
	funcao := closure()
	funcao()
}
