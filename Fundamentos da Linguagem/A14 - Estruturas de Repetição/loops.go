package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i =", i)
		time.Sleep(time.Second)

	}

	for j := 1; j <= 10; j++ {
		fmt.Println("Incrementando J", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{
		"Davi", "João", "Lucas",
	}

	// ITERANDO SOB FOR, OBRIGATÓRIO 2 VARIÁVEIS
	for indice, nome := range nomes {
		fmt.Println("índice=", indice, " Nome:", nome)
	}

	// ITERAR SOB STRING

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "davi",
		"sobrenome": "marques",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// LOOP INFINITO
	v := 0
	for {
		fmt.Println(v)
		v++
		if v > 10 {
			break
		}
	}
}
