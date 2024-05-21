package main

import "fmt"

func main() {

	// IFs

	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// No caso abaixo, outroNumero recebe valor de numero
	// E é validado se é maior que 0

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número maior que 0")
	}

	// Depois que o IF é finalizado, a variável tem o seu tempo de vida
	// Finalizado!

	// Else IF

	if numero > 0 {
		fmt.Println("Maior que 0")
	} else if numero < -10 {
		fmt.Println("Menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

	// SWITCH (NÃO PRECISA DE BREAK)

	fmt.Println("Switch")

	fmt.Println("Switch1\n---------------------")

	var dia1 string = diaDaSemana(3)
	fmt.Println(dia1)

	var dia2 string = diaDaSemana(5)
	fmt.Println(dia2)

	var dia3 string = diaDaSemana(10)
	fmt.Println(dia3)

	fmt.Println()

	// NOVO SWITCH

	fmt.Println("Switch2\n---------------------")

	var dia4 string = diaDaSemana2(1)
	fmt.Println(dia4)

	var dia5 string = diaDaSemana2(5)
	fmt.Println(dia5)

	var dia6 string = diaDaSemana2(10)
	fmt.Println(dia6)

	fmt.Println()

	// NOVO SWITCH COM FALLTHROUGH => CAI NA PROXIMA CONDIÇÃO

	fmt.Println("Switch3\n---------------------")

	var dia7 string = diaDaSemana3(1)
	fmt.Println(dia7)

}

func diaDaSemana(num int) string {
	switch num {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "NaN"
	}

}

func diaDaSemana2(numero int) string {

	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "NaN"
	}

	return diaDaSemana
}

func diaDaSemana3(numero int) string {

	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "NaN"
	}

	return diaDaSemana
}
