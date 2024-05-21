package main

import "fmt"

/*

As Workers Pools são filas de tarefas e os Workers/processos
Vão pegando as tarefas de maneira independente

No caso de fibonacci, seriam vários processos fazendo o cálculo
de Fibonacci

*/

func main() {
	tarefas := make(chan int, 45) // Canal com troca de dados INT e buffer de 45
	resultados := make(chan int, 45)

	go worker(tarefas, resultados) // Pode ser chamadas quantas vezes você quiser
	go worker(tarefas, resultados) // Para assim melhorar o processo
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	/// ...

	for i := 0; i < 45; i++ { // Preenchendo o canal de tarefas
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	// No posicionamento da <- antes do chan, significa que o canal é de entrada, Read-Only
	// Já depois da palavra reservada chan<-, significa que é apenas saída, ou seja, Write-Only

	// Conforme as tarefas vão chegando, ele vai executando Fibonacci

	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}

}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
