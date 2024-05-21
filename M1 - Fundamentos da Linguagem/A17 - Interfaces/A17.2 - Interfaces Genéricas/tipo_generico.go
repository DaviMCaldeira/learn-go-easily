package main

import "fmt"

func generica(interf interface{}) {
	{
		fmt.Println(interf) // Tipo Genérico, ou seja, qualquer coisa pode ser colocada, INT, STRING, FLOAT...
	}
}

func main() {
	generica("Teste")
	generica(true)
	generica(20.5)

	mapGeneric := map[interface{}]interface{}{
		"teste":    1232123123,
		3:          40.5,
		"adsasdas": 1132312,
	}

	fmt.Println(mapGeneric)

	// No caso acima fica explicito o quão bagunçado pode ficar dependendo do mal uso do tipo genérico
}
