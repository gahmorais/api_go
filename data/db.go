package data

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=morais host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}
