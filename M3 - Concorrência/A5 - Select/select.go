package main

import (
	"fmt"
	"time"
)

func main() {

	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {

		// 	mensagemCanal1 := <-canal1 // Isso ocasiona a perda de tempo devido ao bloqueio
		// 	// Por conta do tempo de execução ser diferente em ambos os casos
		// 	// Para isso usamos o select, que é o switch da concorrência

		// 	fmt.Println(mensagemCanal1)

		// 	mensagemCanal2 := <-canal2
		// 	fmt.Println(mensagemCanal2)
		// }

		// Agora com o uso do Select o canal1 não precisa
		// aguardar o canal2 e vice versa

		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
