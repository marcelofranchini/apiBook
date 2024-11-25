package db

import (
	"apiBook/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectarDb() (*sql.DB, error) {

	db, erro := sql.Open("mysql", config.DbString)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
