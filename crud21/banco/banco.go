package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //driver de conexao com mysql
)

// Conectar abre a conxao com o BD
func Conectar() (*sql.DB, error) { //ou vai retornar a conexao ou o erro
	stringConexao := "go_user:Go_pass123@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil { //verificar conexao
		return nil, erro
	}

	return db, nil
}
