package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de Conexão com o Mysql
)

// Conectar CONECTA AO BANCO DE DADOS
func Conectar() (*sql.DB, error) {
	stringConexao := "root:260803@/estac?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
