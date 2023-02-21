package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverArea(f forma) {
	fmt.Println("area = ", f.area())
}

func main() {
	retangulo := retangulo{
		altura:  10,
		largura: 15,
	}
	escreverArea(retangulo)

	circulo := circulo{
		raio: 10,
	}
	escreverArea(circulo)
}
