package repositories

import (
	"apiBook/src/models"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

// repo de usuários apontando para o db
func UsuarioRepository(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repository Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?,?,?,?)",
	)

	if erro != nil {
		return 0, nil
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}
