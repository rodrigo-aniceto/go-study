package crud21

import (
	"fmt"
	"log"
	"modulo/crud21/servidor"
	"net/http"

	"github.com/gorilla/mux"
)

// go get github.com/gorilla/mux

func Funcao() {
	//CRUD - CREATE, READ, UPDATE, DELETE

	// CREATE - POST
	// READ - GET
	// UPDATE - PUT
	// DELETE - DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
