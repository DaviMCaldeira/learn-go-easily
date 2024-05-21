package main

import "fmt"

func main() {
	var retorno string = func(texto string) string {
		return texto
	}("Passando Parâmetro")

	fmt.Println(retorno)
}
