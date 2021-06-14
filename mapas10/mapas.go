package mapas10

import "fmt"

func Funcao() {

	// todas as chaves devem ser do mesmo tipo, o mesmo pra valores
	usuario := map[string]string{ // estrutura de chave/valor
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"segundo":  "Felipe",
		},
		"curso": {
			"nome":   "engenharia",
			"campus": "Darcy Ribeiro",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["curso"]["nome"])

	//adicionar chave
	usuario2["endereço"] = map[string]string{
		"bairro": "asa sul",
	}

	fmt.Println(usuario2)

	//deletar chave
	delete(usuario2, "nome")
	fmt.Println(usuario2)

}
