// Package repository contém todos os repositórios do módulo de usuário.
package repository

/**
 * TODO: Este arquivo é a fábrica de repositórios do módulo de usuário.
 * Ele recebe a conexão de uma ou mais fontes de dados e retorna um container
 * de repositórios que podem ser usados pelo aplicativo.
 */

import (
	"database/sql"
	"errors"
	"fmt"

	userinterfaces "project/internal/modules/user/interface"
	postgresrepo "project/internal/modules/user/repository/postgre_sql"
)

// DBConnections agrupa as conexões de banco de dados
type DBConnections struct {
	Postgres *sql.DB
}

// NewUserRepository atua como uma fábrica que implementa alguma conexão de banco de dados
func NewUserRepository(repoType string, conns DBConnections) (userinterfaces.UserRepository, error) {
	switch repoType {
	case "postgres":
		if conns.Postgres == nil {
			return nil, errors.New("repository factory: conexão Postgres (*sql.DB) não fornecida no container para o tipo 'postgres'")
		}

		return postgresrepo.NewUserRepositoryPostgre(conns.Postgres), nil

	default:
		return nil, fmt.Errorf("repository factory: tipo de repositório desconhecido ou não suportado: %s", repoType)
	}
}
