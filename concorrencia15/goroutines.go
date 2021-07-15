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

	//sicronizar por canais

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

	fmt.Println("Fim do programa")
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
