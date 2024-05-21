package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa // Nos não colocamos o pessoa pessoa para não funcionar
	// da seguinte maneira estudante.pessoa.nome e sim estudante.nome
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Pessoa", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "BCC", "PUC"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
}
