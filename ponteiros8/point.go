package ponteiros8

import "fmt"

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
}
