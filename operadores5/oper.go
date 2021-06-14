package operadores5

import (
	"fmt"
)

func Funcao() {
	soma := 1 + 2
	subtracao := 3 - 1
	divisao := 10 / 4
	mult := 3 * 4
	resto := 5 % 2

	fmt.Println(soma, subtracao, divisao, mult, resto)

	// ele não permite operações com tipo de diferentes

	var str1 string = "Olá!"
	str2 := "Olá!!"

	fmt.Println(str1, str2)

	// eh possivel comparar com > >= == <= < !=
	// os operadores logicos são &&, || e !

	// eh possível fazer numero++ e numero+=1
	// não existe ++numero aqui, mas ningue usa mesmo

	//não existe operador ternário:
	//variavel := numero > 5 ? "Maior que 5" : "Menor que 5"
	//mas isso tbm não eh importante
}
