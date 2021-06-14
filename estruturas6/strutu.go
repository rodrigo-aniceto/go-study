package estruturas6

import "fmt"

type endereco struct {
	cidade  string
	lugar   string
	morador usuario
}

type usuario struct {
	nome  string
	idade uint8
}

func Funcao() {

	var u usuario

	fmt.Println(u) //valores default

	u.nome = "João"
	u.idade = 19

	fmt.Println(u) //eh possivel imprimir os conteudos dele assim

	u2 := usuario{"Davi", 24}

	fmt.Println(u2)

	u3 := usuario{nome: "Pedro"} // é possivel setar apenas alguns dos campos, os outros ficaram com o valor default

	fmt.Println(u3)

	local := endereco{"Brasília", "asa sul", u3} //struct dentro do outro
	fmt.Println(local.morador.idade)
}
