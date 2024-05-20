package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 10
	fmt.Println(array1)

	array2 := [5]string{"Posicao1", "Posicao2"}
	array2[2] = "Posicao3"

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Tamanho baseado na quantidade de valores
	array3[0] = 5

	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice1)
	slice1 = append(slice1, 20)
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array1))

	slice2 := array2[1:3] // 1 Inclusivo, 3 exclusivo
	fmt.Println(slice2)

	array2[1] = "Posic Alterada"

	fmt.Println(slice2)

	// Arrays Internos

	slice3 := make([]float32, 10, 11) //TIPO, TAMANHO, CAPACIDADE TOTAL
	// MAKE CRIA UM ARRAY de 11 posic e ela retorna um slice com as 10 primeiras
	// posicoes do array.
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 8)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	// SE FOR ESTOURADO ESSE LIMITE, O GO PEGA UM ARRAY E
	// DOBRA A CAP E REFERENCIA O NOVO ARRAY

	slice4 := make([]float32, 5)

	fmt.Println(slice4)

	slice4 = append(slice4, 20.0)

	fmt.Println(slice4)

}
