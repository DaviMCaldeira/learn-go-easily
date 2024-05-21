package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*

A ideia do Multiplexador é fazer a junção dos canais, no caso que usamos no último arquivo
Se eu fizer a chamada da função escrever 2x, vamos juntar e unificar tudo em um canal só para
a main poder tratar.

*/

func main() {
	// O Multiplexador basicamente olha dois canais e unifica isso em um único canal

	canal := multiplexar(escrever("Olá mundo!"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

func multiplexar(canal1 <-chan string, canal2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalDeSaida <- mensagem
			case mensagem := <-canal2:
				canalDeSaida <- mensagem
			}

		}
	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto) // Sprintf porque ele printa e retorna
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}

	}()

	return canal
}
