package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA

	var variavel3 int = 100
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // PEGAR ENDEREÇO DE MEMORIA

	fmt.Println(variavel3, *ponteiro) // desreferenciar * (PEGA OQ TEM NO ENDEREÇO)
	fmt.Println(variavel3, ponteiro)

}
