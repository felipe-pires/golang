package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"modulo/19-crud/banco"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	request, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("falha ao ler o corpo da requisicao"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(request, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios(nome, email) values (?,?)")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insert, erro := statement.Exec(usuario.Nome, usuario.Email)

	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	id, erro := insert.LastInsertId()

	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso Id: %d", id)))

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	defer db.Close()

	query, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuarios"))
		return
	}

	defer query.Close()

	var usuarios []usuario

	for query.Next() {
		var usuario usuario

		if erro := query.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear usuario"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuarios para json"))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	ID, erro := strconv.ParseInt(params["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter parametro para int"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	defer db.Close()

	query, erro := db.Query("select * from usuarios where id = ?", ID)

	if erro != nil {
		w.Write([]byte("Erro ao buscar usuario"))
		return
	}

	defer query.Close()

	var usuario usuario
	if query.Next() {
		if erro := query.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear usuario"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario para json"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	ID, erro := strconv.ParseInt(params["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter parametro para int"))
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler body da requisicao"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(body, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter json para struct de usuario"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar com banco de dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, erro := strconv.ParseInt(params["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter parametro para int"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")

	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao excluir usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
