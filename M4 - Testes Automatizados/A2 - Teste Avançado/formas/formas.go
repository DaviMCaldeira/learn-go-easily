package formas

import (
	"math"
)

type Forma interface {
	area() float64
	tipo() string
}

type Retangulo struct {
	altura  float64
	largura float64
}

func (retang Retangulo) Area() float64 {
	return retang.altura * retang.largura
}

func (retang Retangulo) Tipo() string {
	return "retangulo"
}

type Circulo struct {
	raio float64
}

func (circ Circulo) Area() float64 {
	return math.Pi * math.Pow(circ.raio, 2)
}

func (circ Circulo) Tipo() string {
	return "circulo"
}
