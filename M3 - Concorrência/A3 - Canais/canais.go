package main

import (
	"fmt"
	"time"
)

// O objetivo dos canais é sincronizar, mas sem se preocupar
// com contador e nada do tipo, deixando o código mais organizado

func main() {
	// Canal = Canal de comunicação
	// Esse envio e recebimento de dados são usados para sincronizar
	// as Routines

	canal := make(chan string) // Só vai poder enviar e receber Strings
	canal2 := make(chan string)
	go escrever("Olá mundo", canal)
	go escrever("Olá mundo2", canal2)

	fmt.Println("Depois da Routine do escrever e antes da espera do canal.")

	// mensagem := <-canal // Nesse ponto que sincroniza escrever com o canal | Espera o canal receber o valor
	// fmt.Println(mensagem)

	// E se quisermos percorrer por todo o for da função escrever?

	for {
		mensagem, aberto := <-canal // variável aberto para ver se o canal foi fechado ou não
		if !aberto {
			break
		}
		fmt.Println(mensagem) // Isso vai dar deadlock, porque o programa vai esperar receber algum valor, mas pelo for já ter terminado, isso nunca vai ser finalizado
		// Deadlock é em runtime e não em compilação

		// Para resolver isso, iremos sinalizar se o canal estará fechado ou aberto.
	}

	// Outra maneira de ser reescrita a função acima é

	for mensagem := range canal2 {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa")

}

func escrever(texto string, canal chan string) {

	for i := 0; i < 5; i++ {
		canal <- texto // Enviando o dado para o canal, após recebimento, acontece a sincronização
		time.Sleep(time.Second)
	}

	close(canal) // Para evitar deadlocks
}
