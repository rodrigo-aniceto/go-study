package servidor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"modulo/crud21/banco"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(rw http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		rw.Write([]byte("Falha ao ler o corpo da requisicao"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		rw.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		rw.Write([]byte("Erro ao conetar com o Banco de Dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email ) values (?, ?)")
	if erro != nil {
		rw.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		rw.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		rw.Write([]byte("Erro ao obter o ID inserido"))
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso. Id: %d", idInserido)))

}

func BuscarUsuarios(rw http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		rw.Write([]byte("Erro ao conetar com o Banco de Dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("Select * from usuarios")
	if erro != nil {
		rw.Write([]byte("Erro ao buscar usuários"))
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			rw.Write([]byte("Erro ao escanear usuários"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	rw.WriteHeader(http.StatusOK)

	//nova maneira de converter o slice para json
	if erro := json.NewEncoder(rw).Encode(usuarios); erro != nil {
		rw.Write([]byte("Erro ao converter os usuários para json"))
		return
	}

}

func BuscarUsuario(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32) //o numero está em base 10, 32 o tamanho em bits
	if erro != nil {
		rw.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		rw.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		rw.Write([]byte("Erro ao buscar o usuario no banco"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			rw.Write([]byte("Erro ao escanear o usuario"))
			return
		}
	}

	rw.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(rw).Encode(usuario); erro != nil {
		rw.Write([]byte("Erro ao converter o usuarios para json"))
		return
	}

}

func AtualizarUsuario(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		rw.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		rw.Write([]byte("Falha ao ler o corpo da requisicao"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		rw.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		rw.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		rw.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	//tem q ser o ID q vei como parametro
	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		rw.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	rw.WriteHeader(http.StatusNoContent)

}

func DeletarUsuario(rw http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		rw.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		rw.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		rw.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		rw.Write([]byte("Erro ao deletar o usuário"))
		return
	}

	rw.WriteHeader(http.StatusNoContent)

}
