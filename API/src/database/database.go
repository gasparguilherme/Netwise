package database

import (
	"database/sql"

	"github.com/gasparguilherme/Netwise/api/src/config"
)

// connection abre a conexao com o banco de dados e a retorna
func Connection() (*sql.DB, error) {
	db, err := sql.Open("psql", config.StringConnectionBase)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
