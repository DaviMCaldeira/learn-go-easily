package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	// 1 1 2 3 5 8 13 -> FIBONACCI // SE NÃO PARAR PODE ESTOURAR A PILHA
	// PORQUE CADA CHAMADA VAI ESTAR ESPERANDO SEU RETORNO => STACK OVERFLOW

	var posicao uint = 12

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
