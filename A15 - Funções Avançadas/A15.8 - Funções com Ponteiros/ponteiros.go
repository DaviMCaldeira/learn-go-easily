package main

import "fmt"

func inverterSinal(num int) int {
	return num * -1
}

func inverterSinalComPonteiros(num *int) {
	*num = *num * -1
}

func main() {

	// No caso abaixo, está sendo feito uma cópia do número e depois inverte a cópia, tanto que número e número invertido serão diferentes

	numero := 20
	numeroInvertido := inverterSinal(numero)

	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	// Agora neste caso será usando ponteiros

	novoNumero := 40
	fmt.Println(novoNumero)

	inverterSinalComPonteiros(&novoNumero) // & Para referenciar que é um endereço de memória, chegando na func e usando o * ele vai desreferenciar e alterar o valor

	fmt.Println(novoNumero)

}
