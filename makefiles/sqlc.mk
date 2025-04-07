# Define todos os targets
.PHONY: sqlc new-migration

# Importa as definições de cores
include makefiles/colors.mk

# Nome do binário
BINARY_NAME=myapp

# Diretório de origem
SRC_DIR=./cmd/api

# Target app
sqlc:
	@echo
	@echo "=== SQLC ==="
	@echo
	@echo "  make sqlc generate            - Gera código SQLC"
	@echo "  new migration                 - Cria interativamente novos arquivos de migração"
	@echo

# Para os containers
sqlc-generate:
	sqlc generate

## new-migration: Cria interativamente novos arquivos de migração .up.sql e .down.sql
new migration:
	@echo "Qual o nome da migração? (use snake_case): "; \
	read migration_name; \
	if [ -z "$$migration_name" ]; then \
		echo "Erro: O nome da migração não pode ser vazio."; \
		exit 1; \
	fi; \
	echo ">> Criando migração: $$migration_name"; \
	migrate create -ext sql -dir internal/infrastructure/database/migration $$migration_name; \
	echo ">> Arquivos de migração criados em internal/infrastructure/database/migration/"
