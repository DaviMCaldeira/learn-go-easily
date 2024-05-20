package auxiliar

import "fmt"

func Escrever() { // Public Escrever \ Private escrever
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2() // Consigo chamar pois estão dentro do mesmo package
}
