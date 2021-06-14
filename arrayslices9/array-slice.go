package arrayslices9

import "fmt"

func Funcao() {

	var arr1 [5]string // todos os elementos tem que ter o mesmo tipo

	arr1[0] = "posicao 1"

	fmt.Println(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5} // o tamanho eh fixo, sem choro

	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5, 6} // assim quando vc não sabe o tamanho, mas ele continua sendo fixo

	fmt.Print(arr3, "\n\n")

	// slices são arrays com tamanho variavel, mas ainda do mesmo tipo

	sli := []int{12, 13, 14, 15, 16, 17, 18}
	fmt.Println(sli)

	sli = append(sli, 10)
	sli = append(sli, 15)

	fmt.Println(sli)

	// slices tbm são ponteiros de array, se atribuir para ele um subarray, eh possivel editar o array principal assim
	// da mesma maneira tbm eh possivel editar o slice pelo array principal

	sli2 := arr2[1:3]
	fmt.Println(sli2)

	arr2[1] = 90

	fmt.Println(sli2)

	sli2[0] = 900

	fmt.Println(arr2)

	fmt.Print("\nARRAY INTERNOS:\n")
	//ARRAY INTERNO

	//funcão make aloca espaço na memória
	sli3 := make([]float32, 10, 15) // tipo, tamanho do slice, tamanho maximo q pode crescer
	fmt.Println(sli3)
	fmt.Println(len(sli3)) // length
	fmt.Println(cap(sli3)) // capacidade

	sli3 = append(sli3, 11)
	sli3 = append(sli3, 12)
	sli3 = append(sli3, 13)
	sli3 = append(sli3, 14)
	sli3 = append(sli3, 15)
	sli3 = append(sli3, 16)

	//estourei o tamanho de propósito
	fmt.Println(sli3)
	fmt.Println(len(sli3)) // length
	fmt.Println(cap(sli3)) // capacidade
	// ele dobrou a capacidade máxima para suportar o crescimento desse array de memória

	//atras de todo slice existe um array em memória sendo manipulado
}
