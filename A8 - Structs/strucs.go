package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	fmt.Println("Arquivo STRUCTS")

	enderecoExemplos := endereco{
		"Campo Azul",
		203,
	}

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	user2 := usuario{
		"Davi",
		21,
		enderecoExemplos,
	}

	fmt.Println(user2)

	user3 := usuario{

		nome: "Davi"}

	fmt.Println(user3)
}
