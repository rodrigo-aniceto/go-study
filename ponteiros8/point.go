package ponteiros8

import "fmt"

func inverteSinal(num int) int {
	num = num * -1
	return num
}

func inverteSinalPonteiro(num *int) {
	*num = *num * -1
}

func Funcao() {

	var1 := 10
	var2 := &var1

	// var var3 *int // outra maneira de declarar
	// ponteiro eh um tipo de dados que gaurda referencias de memoria

	fmt.Println(var2)

	fmt.Println(&var2)
	fmt.Println(*var2)

	var1++

	fmt.Println(var1, *var2)

	//funções com ponteiros

	numero := 20
	numeroInvertido := inverteSinal(numero)
	fmt.Println("numero Invertido:", numeroInvertido)
	fmt.Println("numero Normal:", numero)

	numero2 := 35
	inverteSinalPonteiro(&numero2)
	fmt.Println("invertido por referencia:", numero2)
}
