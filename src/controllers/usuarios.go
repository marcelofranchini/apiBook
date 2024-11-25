package controllers

import (
	"apiBook/src/db"
	"apiBook/src/models"
	"apiBook/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(body, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := db.ConectarDb()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositories.UsuarioRepository(db)
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioID)))

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar Usuários"))

}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar Usuário"))

}
func AlterarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alterar Usuário"))

}
func RemoverUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Remover Usuário"))

}
