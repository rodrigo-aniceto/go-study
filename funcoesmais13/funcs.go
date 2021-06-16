package funcoesmais13

import "fmt"

func funcao1() {
	fmt.Println("Executnado funcao 1- fim do programa")
}

func funcao2() {
	fmt.Println("Executnado funcao 2")
}

func recuperaBagunca() {
	if r := recover(); r != nil { //se estiver em panic ele entra aqui e encerra o panic retornando o fluxo pra funcção externa a do panic
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperaBagunca()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA EH SEIS!") //isso faz o programa entrar em panico e para tudo, os defers são executados primeiro
	// se tiver um defer que chama a função recover o programa ta salvo, e essa função vai retornar o valor padrão do tipo da saida
}

func closure() func() { //função que retorna uma função
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto) // o valor da variavel texto vai ser o daqui, independente de la fora ter outra variável com esse nome
	}

	return funcao

}

func Funcao() {

	defer funcao1()
	//o defer faz ele adiar a chamada a essa função até o ultimo momento, só será chamada antes do retorno
	// o valor de parametos se mantem no momento q chamou
	// isso eh ultil tipo para fechar a conexão do BD que vc não sabe quando
	funcao2()
	fmt.Println("Lalalala")
	fmt.Println("123")

	fmt.Println(alunoAprovado(6, 6)) // se a função entrar em panico, o defer vai ser executado antes dela

	funcaonova := closure() //funcao closure, ela guarda as variveis do contexto em que foi criada
	funcaonova()

}
