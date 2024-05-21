package main

import (
	"fmt"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main!")
	auxiliar.Escrever()
	// auxiliar.escrever2 - Não consigo usar por causa da letra minuscula ser considerada como private
}
