package main

import "fmt"

func main() {
	var variavel1 string = "variavel1" // Explicitos
	variavel2 := "Variavel2"           // Implicito

	fmt.Println(variavel1)

	fmt.Println(variavel2)

	var (
		variavel3 string = "teste1"
		variavel4 string = "teste2"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "teste3", "teste4"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "const1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
}
