package json17

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct { // marca usando crase mesmo..
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

type gato struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func Funcao() {

	//gerando json
	fmt.Println("EXPORTANTO JSON")
	//a partir de struct
	c := cachorro{"Rex", "Golden", 3}

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)                  // assim ele vi exibir como um slice de bytes
	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) // converte para um json normal

	// agora usando map

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))

	//processo de importar json
	fmt.Println("IMPORTANDO JSON")
	//para uma struct
	gatoEmJSON := `{"nome":"Ping","Raca":"Sírio","Idade":3}`

	var g gato

	//precisa converter a string para slice de bytes
	// ira escrever em &g, funcao retorna um erro
	if erro = json.Unmarshal([]byte(gatoEmJSON), &g); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(g)

	//agora para um map
	gato2EmJSON := `{"Nome":"Tom","Raca":"Persa"}`
	g2 := make(map[string]string) //precisa definir os tipos corretamente

	if erro = json.Unmarshal([]byte(gato2EmJSON), &g2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(g2)

}
