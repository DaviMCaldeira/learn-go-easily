package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (user usuario) salvar() { // user é a variável que referencia a struct user, para podermos chamar user.nome, user.idade... user é apenas um nome, poderia ser qualquer coisa
	fmt.Printf("Salvando os dados do usuário %s no DB\n", user.nome) // sobre o Printf é para podermos formatar as strings, https://gobyexample.com/string-formatting
}

func (user usuario) maiorDeIdade() bool {
	return user.idade >= 18
}

func (user *usuario) fazerAniversario() { // Caso queira alterar alguma propriedade da própria struct, o recomendado é o uso de ponteiros.
	user.idade += 1
}

func escrever() {
	fmt.Println("Escrevendo...")
}

func main() {
	escrever() // Não é um método, pois não está associada a nada, um método seria por exemplo, a struct usuario, e em pessoa vou ter a função (método) pessoa.salvar()

	var userOfDB usuario = usuario{
		"Davi",
		17,
	}

	userOfDB.salvar() // Agora aqui é um método pois está sendo referenciado uma função a partir de uma struct

	fmt.Println(userOfDB.maiorDeIdade())

	userOfDB.fazerAniversario()

	fmt.Println(userOfDB.maiorDeIdade())
}
