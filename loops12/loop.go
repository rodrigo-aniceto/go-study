package loops12

import (
	"fmt"
	"time"
)

// não existe while, apenas for

func Funcao() {

	i := 0

	for i < 10 { // semelhante a um while
		fmt.Println("Valor de i:", i)
		time.Sleep(time.Second)
		i++
	}

	fmt.Println("Valor final de i:", i)

	for j := 0; j < 5; j++ { //for tradicional
		fmt.Println("Valor de j:", j)
		time.Sleep(time.Second)
	}

	nomes := []string{"Pedro", "Thiago", "João"}
	for indice, nome := range nomes { // a primeira variavel eh sempre o valor a segunda sempre o nome
		fmt.Println(indice) // se não quiser usar os dois, coloue um _
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" { // o mesmo q o anterior mas percorerendo uma string
		fmt.Println(indice, string(letra)) // se não colocar strng antes, ele vai exibir o codigo asc do character
	}

	//agora com um MAP
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario { // a chave e o valor vao ser a chave e o valor mesmo
		fmt.Println(chave, valor)
	}

	//assim vc faz um loop infinito
	// for {
	//
	//}
}
