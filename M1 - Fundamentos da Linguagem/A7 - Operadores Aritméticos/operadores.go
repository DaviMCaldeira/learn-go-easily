package main

import "fmt"

func main() {
	// ARITMÉTICOS

	// +, -, /, * e %

	// ATRIBUIÇÃO

	var variavel string = "texto"
	variavel2 := "texto2"

	fmt.Println(variavel, variavel2)

	// Relacionais

	fmt.Println(1 == 1)
	fmt.Println(2 <= 3)
	fmt.Println(4 >= 2)
	fmt.Println(5 > 1)
	fmt.Println(1 > 3)
	fmt.Println(1 != 2)

	// Lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!(verdadeiro || falso))
	fmt.Println(!falso)
	fmt.Println(!verdadeiro)

	// Operadores Unários

	numero := 1
	numero++
	numero += 15

	numero--
	numero -= 20

	numero *= 3

	numero /= 3

	numero %= 3

	// Operadores Ternários

	// Não existe no go, teriamos que usar if-else

	if numero > 5 {
		variavel = "Maior que 5"
	} else {
		variavel = "Menor que 5"
	}

}
