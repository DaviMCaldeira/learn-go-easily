package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` // Chave que vão se tornar dentro do Json, convenção no JSON é tudo ser feito em letras minúsculas
	Raca  string `json:"raca"` // se quiser ignorar, você pode usar um "-"
	Idade uint   `json:"idade"`
}

func main() {
	// json.Marshal() // Converte Map => Json, Struct => Json

	// json.Unmarshal() // Converte Json => Map, Json => Struct

	c := cachorro{
		"Teste",
		"Golden",
		5,
	}

	cachorroEmJson, erro := json.Marshal(c) // retorna slice de bytes e um erro

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJson)

	fmt.Println(bytes.NewBuffer(cachorroEmJson)) // Traduz o slice de Bytes.

	c2 := map[string]string{
		"nome": "Tob",
		"raca": "Bulldog",
	}

	fmt.Println()

	cachorroEmJson, erro = json.Marshal(c2) // retorna slice de bytes e um erro

	fmt.Println(cachorroEmJson)

	fmt.Println(bytes.NewBuffer(cachorroEmJson)) // Traduz o slice de Bytes.

}
