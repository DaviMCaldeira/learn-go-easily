package main

import (
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
	cachorroEmJson := `{"nome":"Teste","raca":"Golden","idade":5}`

	var c cachorro

	// Tem que ser no formato Slice de Byte
	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorroEmJson2 := `{"nome":"Tob","raca":"Bulldog"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorroEmJson2), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
}
