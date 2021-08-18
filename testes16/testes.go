package testes16

import (
	"fmt"
	"modulo/testes16/enderecos"
)

//apenas executa funcao de endereco.go, os testes estão em enderecos_test.go
// go test executa todos testes dentro da pasta que estou
// para executar todos os testes do projeto, executar go test ./...
// verbose: go test -v
// exibe quanto % das linhas de código foram cobertas: go test --cover

// para identificar quais linhas foram cobertas e quais não:
// go test --coverprofile cobertura.txt //gera um arquivo com a cobertura dos testes
// go tool cover --func=cobertura.txt // le o arquivo e exibe os % por função - isso eh pouco util
// go tool cover --html=cobertura.txt // gera relatorio html das linhas cobertas no teste - use esse

func Funcao() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)

	tipoEndereco = enderecos.TipoDeEndereco("SQN 202")
	fmt.Println(tipoEndereco)
}
