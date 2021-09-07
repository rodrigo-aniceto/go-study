package main

// 'go mod init modulo' cria um executavel de nome 'modulo' referente a todo o projeto
// 'go build' recompila todo mundo
// para eecução de um unico arquivo (que seja do pacote main e tenha função main) 'go run arquivo.go'
// go install faz o mesmo que go build mas cria o executavel na raiz do go
// adicionar pacoe externo: 'go get github.com/badoux/checkmail'
// remover do go mod as referencias a pacotes que não são mais usados: 'go mod tidy'

import (
	"fmt"
	"modulo/crud21"
	//"modulo/bancodados20"
	//"modulo/html19"
	//"modulo/http18"
	//"modulo/pacote1"
	//"modulo/variaveis2"
	//"modulo/tiposdados3"
	//"modulo/funcoes4"
	//"modulo/operadores5"
	//"modulo/estruturas6"
	//"modulo/heranca7"
	//"modulo/ponteiros8"
	//"modulo/arrayslices9"
	//"modulo/mapas10"
	//"modulo/estcontrole11"
	//"modulo/loops12"
	//"modulo/funcoesmais13"
	//"modulo/metodos14"
	//"modulo/concorrencia15"
	//"modulo/testes16"
)

func main() {
	//pacote1.Escreve()
	//variaveis2.Vars()
	//tiposdados3.Funcao()
	//funcoes4.Funcao()
	//operadores5.Funcao()
	//estruturas6.Funcao()
	//heranca7.Funcao()
	//ponteiros8.Funcao()
	//arrayslices9.Funcao()
	//mapas10.Funcao()
	//estcontrole11.Funcao()
	//loops12.Funcao()
	//funcoesmais13.Funcao()
	//metodos14.Funcao1()
	//metodos14.Funcao2()
	//metodos14.Funcao3()
	//concorrencia15.Funcao()
	//concorrencia15.FuncaoCanais()
	//concorrencia15.FuncaoSelect()
	//concorrencia15.FuncaoPadroes()
	//testes16.Funcao()
	//json17.Funcao()
	//http18.Funcao()
	//html19.Funcao()
	//bancodados20.Funcao()
	crud21.Funcao()
}

// a init eh chamada antes da main
// todos os arquivos podem ter uma init, ela eh chamada assim que os pacotes são importados
// o ultimo init a ser chamado eh esse aqui:
func init() {
	fmt.Println("Hello word!")
}
