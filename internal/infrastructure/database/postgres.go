// package database é responsável por criar uma conexão com o banco de dados.
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Config é a estrutura que armazena as configurações do banco de dados.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
} 

// NewPostgresConnection cria uma nova conexão com o banco de dados.
func NewPostgresConnection(config *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
	)

	// Cria uma nova conexão com o banco de dados
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Verifica se a conexão com o banco de dados está ok
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}