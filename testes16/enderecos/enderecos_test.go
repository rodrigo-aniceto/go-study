package enderecos_test

import (
	. "modulo/testes16/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

//TESTE DE UNIDADE
// unica situação em que o pacote pode ter nome diferente _test, use um . na hora de importar o pacote orignal
//para executar basta ir na pasta e executar go test
func TestTipoDeEndereco(t *testing.T) {
	//t.Parallel() //coloque isso para rodar testes em paralelo

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "rua"},
		{"Avenida jacaré", "avenida"},
		{"Estrada da perdição", "estrada"},
		{"Rodovia dos imigrantes", "rodovia"},
		{"AVENIDA ENGARAFADA", "avenida"},
		{"SQN 103", "Tipo Inválido"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("Tipo incorreto. Esperado: %s, recebido: %s", cenario.retornoEsperado, retornoRecebido)
		}
	}

}

func TestQualquer(t *testing.T) {
	//t.Parallel()

	if 1 > 2 {
		t.Error("Algo está muito errado..")
	}
}
