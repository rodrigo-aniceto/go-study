package http18

import (
	"log"
	"net/http"
)

func telaRaiz(rw http.ResponseWriter, r *http.Request) {
	//quando usuário acessar /home, executar:
	rw.Write([]byte("Pagina Raiz")) // tem q ser slice de bytes
}

func telaHome(rw http.ResponseWriter, r *http.Request) {
	//quando usuário acessar /home, executar:
	rw.Write([]byte("Olá mundo!")) // tem q ser slice de bytes
}

func telaUsuarios(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Estou na pagina de usuários"))
}

func Funcao() {
	// HTTP - cliente faz requisições (requests), servidor manda respostas (responses)
	// rotas são compostas por: URI - idntificador do recurso
	// e metodos - GET, POST, PUT, DELETE

	http.HandleFunc("/", telaRaiz)
	http.HandleFunc("/home", telaHome) //criar funções segundo o padrão de cima

	http.HandleFunc("/usuarios", telaUsuarios)

	// servidor rodando no localhost porta 5000 se declara as rotas antes disso, pq ele vai parar aqui
	log.Fatal(http.ListenAndServe(":5000", nil))
}
