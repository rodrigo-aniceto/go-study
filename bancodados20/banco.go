package bancodados20

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//precisa de: go get github.com/go-sql-driver/mysql
// importa com _ na frente pq sera o pacote sql que vai usar e não esse
func Funcao() {
	//string: usuario:senha@/banco...etc
	stringConexao := "go_user:Go_pass123@/devbook?charset=utf8&parseTime=True&loc=Local"

	//criar conexão
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close() //defer executa no final ou se der erro

	if erro = db.Ping(); erro != nil { //verificar conexao
		log.Fatal(erro)
	}

	fmt.Println("Conexao estabelecida")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()

	fmt.Println(linhas)
}
