package concorrencia15

import (
	"fmt"
	"sync"
	"time"
)

func Funcao() {

	//concorrencia != parelelismo
	// nem toda concorrencia é necessariamente paralela
	go escrever("chamada 1") //executa em paralelo e segue
	escrever("chamada 2")

	//sicronizar com wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) //define o tamanho da fila de espera

	go func() {
		escrever("Chamada 3")
		waitGroup.Done() //subtrai 1 da fila

	}()

	go func() {
		escrever("Chamada 4")
		waitGroup.Done()
	}()

	fmt.Println("Antes do wait")

	waitGroup.Wait() // só irá passar esse ponto quando a fila estiver fazia
	fmt.Println("Depois do wait - sicronizado")

	fmt.Println("Fim do programa")
}

func FuncaoCanais() {

	fmt.Println("Sincronização por canais")

	canal := make(chan string) // cria um canal
	go escreverUsandoCanal("Chamada 10", canal)

	for {
		mensagem, aberto := <-canal //só avança quando um valor for enviado pro canal
		if !aberto {                //boleanico que informa se o canal ainda está aberto
			break
		}
		fmt.Println(mensagem)
	}

	fmt.Println("Canal usando range")

	// outra forma de usar canais
	canal2 := make(chan string) // cria canal
	go escreverUsandoCanal("Chamada 20", canal2)

	for mensagem := range canal2 { //repete enquanto o canal não tenha sido fechado
		fmt.Println(mensagem)
	}

	// canal com buffer
	canal3 := make(chan string, 2) // o buffer tem q ser maior ou igual ao tamanho do conteudo
	canal3 <- "Chamada 30"
	canal3 <- "Chamada 31"

	// se não houvesse buffer isso não ia funcionar, pq ele ia esperar a entrada e a saida ao mesmo tempo
	fmt.Println(<-canal3)
	fmt.Println(<-canal3)

}

func FuncaoSelect() { // isso eh um switch case de canais
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500) //meio segundo
			canal1 <- "canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2) //dois segundos
			canal2 <- "canal 2"
		}
	}()

	for {

		select { // Vai executar o que o canal receber
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}

}

func FuncaoPadroes() { // padroes de concorrencia usados

	// worker pools: varias funcoes workers dividindo as tarefas
	entrada := make(chan int, 40)
	saida := make(chan int, 40)

	go worker(entrada, saida) //aumenta os workers melhora a performace
	go worker(entrada, saida)
	go worker(entrada, saida)
	go worker(entrada, saida)

	for i := 0; i < 40; i++ {
		entrada <- i // cada worker vai ir pegando um numero conforme aparece
	}
	close(entrada)

	for i := 0; i < 40; i++ {
		resultado := <-saida
		fmt.Println(resultado)
	}

	// padrão Generator: funcao q encapsula chamada a goroutine e restorna um canal

	canalX := escreverGenerator("Ola mundo") // essa funcao esconde da main o trabalho de go routine

	for i := 0; i < 10; i++ {
		fmt.Println(<-canalX)

	}

	// padrão multiplexador: juntar mais de um canal
	canalZ := multiplexar(escreverGenerator("lalala"), escreverGenerator("lelele"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-canalZ)
	}
}

func worker(entrada <-chan int, saida chan<- int) { //já especifica se eh um canal de entrada ou saida
	for numero := range entrada {
		saida <- fibonacci((numero))
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func escrever(texto string) {
	for i := 0; i < 3; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func escreverUsandoCanal(texto string, canal chan string) {
	for i := 0; i < 3; i++ {
		canal <- texto // adiciona o texto ao canal
		time.Sleep(time.Second)
	}
	close(canal)
}

func escreverGenerator(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}

func multiplexar(canalA, canalB <-chan string) <-chan string {
	canalSaida := make(chan string)

	// escreve no canal de saida o conteudo dos dois canais
	go func() {
		for {
			select {
			case mensagem := <-canalA:
				canalSaida <- mensagem
			case mensagem := <-canalB:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}
