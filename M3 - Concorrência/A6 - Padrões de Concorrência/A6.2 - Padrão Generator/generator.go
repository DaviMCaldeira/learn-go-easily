package main

/*

O padrão generator deixa as GoRoutines de Forma encapsulada a medida
que é necessário

// Abstraimos da função main a necessidade de declarar as Routines, etc.

*/

import "fmt"

func main() {
	canal := escrever("Olá mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto) // Sprintf porque ele printa e retorna
		}

	}()

	return canal
}
