package heranca7

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint16
}

// diferenca sutil para o exemplo de estruturas
// não eh uma 'pessoa' dentro de 'estudante', são os mesmos atributos herdados
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func Funcao() {
	fmt.Println("Simulacao de heranca em Go")

	p1 := pessoa{"João", "Silva", 22, 180}
	fmt.Println(p1)

	e1 := estudante{p1, "Direito", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)

}
