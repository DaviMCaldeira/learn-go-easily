package main

import "fmt"

var n uint8

func main() {
	fmt.Println("Função main sendo executada!") // A Função INIT pelo compilador do GO será sempre executada primeiramente, independente da ordem delas no arquivo
	fmt.Println(n)
}

func init() {
	fmt.Println("Função init sendo executada!")
	n = 12
}

// A diferença é que você pode ter uma por arquivo e não 1 por pacote (igual a main), usado normalmente usado pra inicializar variaveis, etc...
