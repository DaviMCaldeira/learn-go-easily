package main

import (
	"errors"
	"fmt"
)

func main() {

	// INT

	// int8, int16, int32, int64
	// todos esses valores podem ser u(int) uint é sem negativo
	// int => baseado na arquitetura, sistemas 32 bits int32, 64 bits int64

	var numero int16 = 100
	fmt.Println(numero)

	//alias

	var numero3 rune = 123456 // rune são numeros que representam caracteres
	// rune = int32
	fmt.Println(numero3)

	// BYTE = UINT8

	var numero4 byte = 123
	fmt.Println(numero4)

	// FIM INT

	// REAIS

	var numeroReal1 float32 = 122.3123
	fmt.Println(numeroReal1)

	var numeroReal2 float32 = 1231231231312.3123
	fmt.Println(numeroReal2)

	numeroReal3 := 12312312312.21312 // S.O Bits
	fmt.Println(numeroReal3)

	// FIM REAIS

	// STRINGS

	var string1 string = "OlaPessoal!"
	fmt.Println(string1)

	string2 := "OlaPessoal!"
	fmt.Println(string2)

	char := 'B'
	fmt.Println(char)

	// FIM STRINGS

	// VALORES DEFAULT

	var texto string // VAZIO É UMA STRING VAZIA
	fmt.Println(texto)

	var inteiro int // VAZIO É O ZERO
	fmt.Println(inteiro)

	// FIM VALORES DEFAULT

	// BOOLEAN

	var booleano1 bool = true
	fmt.Println(booleano1)

	// FIM BOOLEAN

	// ERRO TYPE

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
