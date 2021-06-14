package variaveis2

import "fmt"

func Vars() {
	fmt.Println("oi")

	var str string = "string que criei"

	str2 := "segunda string que criei" //declaração implicita

	fmt.Println(str)
	fmt.Println(str2)

	var (
		str3 string = "lalala"
		str4 string
	)

	str4 = "lellele" // se declarou tem que usar, ou erro

	fmt.Println(str3)
	fmt.Println(str4)

	str5, str6 := "string 5", "string 6"

	println("aqui: " + str5 + " e aqui: " + str6)

	const constante1 = "constante 1 criada"

	fmt.Println(constante1)
}
