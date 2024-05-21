package main

import (
	"fmt"
	"time"
)

func main() {

	// CONCORRENCIA != PARALELISMO
	// PARARELISMO EXECUTANDO AO MESMO TEMPO

	// CONCORRENCIA PODE ESTAR AO MESMO TEMPO,
	// MAS NÃO NECESSÁRIAMENTE,

	// Pontos chave

	// Concorrência:
	// Pode ser realizada em um único núcleo de CPU através de
	// alternância rápida entre tarefas.

	// Paralelismo:
	// Requer múltiplos núcleos ou processadores para
	// executar tarefas simultaneamente.

	// Objetivo:

	// Concorrência:
	// Melhorar a responsividade e utilização do
	// processador, permitindo que várias tarefas progridam.

	// Paralelismo:
	// Melhorar a performance reduzindo o tempo de
	// execução total de tarefas, executando-as ao mesmo tempo.

	// escrever("Olá Mundo") // Nesse caso, a linha 40 nunca vai ser executada
	// Vai ficar esperando ele terminar, para contornamos,
	// usaremos go routines
	go escrever("Olá mundo")       // goroutine
	escrever("Programando em Go!") // Você deve estar se perguntando
	// Porque não adicionar uma goroutine nesse segundo escrever?

	// Simples, pois depois não existirá mais nada para ser feito no programa
	// e ele sera finalizado!

	// O que a goroutine significa, que ela pode executar o restante do código
	// sem necessariamente terminar a execução da primeira parte.
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)

		// No próximo arquivo de goroutines será demonstrado como sincronizar
		// Bem parecido com o await, pois para dar prosseguimento no código, todas as goroutines daquele escopo devem ser finalizadas
	}
}
