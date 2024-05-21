package main

import "fmt"

func main() {
	canal := make(chan string, 2) // Com a adição do buffer de 2
	// Funcionou, visto que o canal só será bloqueado quando o buffer
	// Atingir o limite, nesse caso é 2

	canal <- "Olá mundo" // Ele espera o programa receber, ou seja, bloqueia, visto que ações de entrada e saída de dados são bloqueantes
	// Basicamente está enviado mas nesse momento x não tem ninguém para receber

	canal <- "Programando em Go" // Podemos passar mais um valor visto que
	// Ainda não chegou na capacidade total

	// canal <- "Programando em Go2" // Aqui daria um erro, visto que estourou o buffer

	msg1 := <-canal // Nesse caso aqui seria X+Y miliseconds
	msg2 := <-canal

	fmt.Printf("%s\n%s", msg1, msg2)
}
