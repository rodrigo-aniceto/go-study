package metodos14

import (
	"fmt"
	"math"
)

// eh uma estrutura que define o método area mas não o implementa

type forma interface {
	area() float64
}

// funcaão que recebe qualquer implementação do tipo forma
// não precisava criar ela
func calculaArea(f forma) {
	valor := f.area()
	fmt.Printf("A area da forma eh %0.2f\n", valor)

}

// se eu criar estruturas que tem um metodo de nome area ele vai automaticamente
// considerar ela uma implementação de forma

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
	return math.Pi * c.raio * c.raio
}

func Funcao2() {
	r := retangulo{10, 15}
	calculaArea(r)
	//fmt.Println(r.area())
	c := circulo{6}
	calculaArea(c)
	//fmt.Println(c.area())

}
