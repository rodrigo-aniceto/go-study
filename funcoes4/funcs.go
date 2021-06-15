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

func calculos3(nome string, numeros ...int) int { //função variática - recebe um numero variado de parametros
	// numeros eh um slice, ele deve ser o último parametro
	total := 0
	fmt.Println(nome)
	for _, numero := range numeros {
		total += numero
	}
	return total
}

//função recursiva eh normal
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
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

	fmt.Println(calculos3("abc", 1, 2, 3, 4, 5, 6))

	//funcao anonima -  não tem nome e fica aqui dentro, eh estranho
	retorno2 := func(a, b int) int {
		return a + b
	}(2, 4)

	fmt.Println("função anonima:", retorno2)

	fmt.Println("Função recursiva:")
	posicao := uint(12)
	for i := uint(1); i <= posicao; i++ {
		fmt.Print(fibonacci(i), " ") //chamada recursiva
	}
	fmt.Println()

}
