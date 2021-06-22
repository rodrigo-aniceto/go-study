package metodos14

import "fmt"

// função recebe uma interface sem nenhum campo
// ou seja qualquer dado vai ser considerado implementacao dele
// a função aceita qualquer tipo de dado
func generica(inter interface{}) {
	fmt.Println(inter)
}

func Funcao3() {
	generica("abc")
	generica(2)
	generica(true)

	// isso daki eh um mapa que pode colocar qualquer tipo dos dois lados
	// não faça isso!
	mapa := map[interface{}]interface{}{
		1:            "texto",
		float32(100): true,
		"string":     "string 2",
		true:         float64(-12.888),
	}

	fmt.Println(mapa)

	// a função Println recebe interface para permitir imprimir qualquer dado
}
