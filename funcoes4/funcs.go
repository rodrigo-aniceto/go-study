package funcoes4

import "fmt"

func somar(n1 int8, n2 int8) int8 {

	return n1 + n2
}

func calculos(n1, n2, n3 int32) (int32, int32, int32) {
	r1 := n1 + n2 + n3
	r2 := (n1 + n2 + n3) / 3
	r3 := n1 * n2 * n3
	return r1, r2, r3
}

func calculos2(n1, n2 int) (soma int, subtracao int) { //aqui ele já declara as variaveis de retorno
	soma = n1 + n2
	subtracao = n1 - n2
	return //não precisa completar o return pq ele já sabe o q retorna
}

func Funcao() {
	soma := somar(2, 4)
	fmt.Println(soma)

	// declaracao de variavel do tipo funcao
	var f = func() string {
		txt := "Funcao F"
		return txt
	}

	retorno := f()
	fmt.Println(retorno)

	//funcao com multiplos retornos
	// o _ eh pra ignorar um valor q não vou usar
	v1, _, v3 := calculos(2, 4, 7)
	fmt.Println(v1, v3)

	fmt.Println(calculos2(4, 2))

}
