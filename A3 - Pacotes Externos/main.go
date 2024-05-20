package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main!")
	auxiliar.Escrever()
	// auxiliar.escrever2 - Não consigo usar por causa da letra minuscula ser considerada como private
	msg := checkmail.ValidateFormat("teste@gmail.com")

	fmt.Println(msg)
}
