package main

// 'go mod init modulo' cria um executavel de nome 'modulo' referente a todo o projeto
// 'go build' recompila todo mundo
// para eecução de um unico arquivo (que seja do pacote main e tenha função main) 'go run arquivo.go'
// go install faz o mesmo que go build mas cria o executavel na raiz do go
// adicionar pacoe externo: 'go get github.com/badoux/checkmail'
// remover do go mod as referencias a pacotes que não são mais usados: 'go mod tidy'

import (
	"fmt"
	"modulo/funcoes4"
	//"modulo/loops12"
	//"modulo/estcontrole11"
	//"modulo/mapas10"
	//"modulo/arrayslices9"
	//"modulo/ponteiros8"
	//"modulo/heranca7"
	//"modulo/pacote1"
	//"modulo/variaveis2"
	//"modulo/tiposdados3"
	//"modulo/funcoes4"
	//"modulo/operadores5"
	//"modulo/estruturas6"
)

func main() {
	fmt.Println("Ola mundo!")
	//pacote1.Escreve()
	//variaveis2.Vars()
	//tiposdados3.Funcao()
	funcoes4.Funcao()
	//operadores5.Funcao()
	//estruturas6.Funcao()
	//heranca7.Funcao()
	//ponteiros8.Funcao()
	//arrayslices9.Funcao()
	//mapas10.Funcao()
	//estcontrole11.Funcao()
	//loops12.Funcao()
}
