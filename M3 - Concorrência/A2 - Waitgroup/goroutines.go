package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(4) // Quantidade de GoRoutines

	go func() {
		escrever("Olá mundo")
		waitgroup.Done() // -1 na qtd passada no Add
	}()

	go func() {
		escrever("Programando em Go!")
		waitgroup.Done() // -1 na qtd passada no Add
	}()

	go func() {
		escrever("Go Routine 3!")
		waitgroup.Done() // -1 na qtd passada no Add
	}()

	go func() {
		escrever("Go Routine 4!")
		waitgroup.Done() // -1 na qtd passada no Add
	}()

	waitgroup.Wait() // Esperar a contagem chegar em 0
	// Se não tivesse o Wait, elas nem chegariam a executar, pois ignorariam o fluxo do código e por chegarem no final, já finalizariam o código.

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)

	}
}
