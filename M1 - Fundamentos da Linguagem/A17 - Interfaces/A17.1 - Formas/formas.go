package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
	tipo() string
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma: %s é %0.2f\n", f.tipo(), f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (retang retangulo) area() float64 {
	return retang.altura * retang.largura
}

func (retang retangulo) tipo() string {
	return "retangulo"
}

type circulo struct {
	raio float64
}

func (circ circulo) area() float64 {
	return math.Pi * math.Pow(circ.raio, 2)
}

func (circ circulo) tipo() string {
	return "circulo"
}

func main() {

	// As interfaces são usadas para abstração, neste caso se eu não tivesse usado uma interface, não existiria a possibilidade de usar
	// o escreverArea(r) para circulo e para retangulo visto que são formas diferentes, mas pela abstração foi possível, visto que implementamos
	// o método área em ambos

	// Interface = Flexibilidade

	r := retangulo{
		20.0,
		35.0,
	}

	c := circulo{
		10.0,
	}

	escreverArea(r)
	fmt.Println(r.area())
	escreverArea(c)
	fmt.Println(c.area())
}
