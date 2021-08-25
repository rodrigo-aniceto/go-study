package html19

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//se convenciona manter os templates como variável global
var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func telaHome(rw http.ResponseWriter, r *http.Request) {

	u := usuario{"Felipe da Silva", "felipinho19@hotmail.com"}
	templates.ExecuteTemplate(rw, "home.html", u) //se não for passar nenhum parametro mantenha nil no terceiro
}

func Funcao() {

	templates = template.Must(template.ParseGlob("html19/*.html"))

	http.HandleFunc("/home", telaHome)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
