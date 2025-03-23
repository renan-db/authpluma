# Define todos os targets
.PHONY: app

# Importa as definições de cores
include makefiles/colors.mk

# Nome do binário
BINARY_NAME=myapp

# Diretório de origem
SRC_DIR=./cmd/api

# Target app
app:
	@echo
	@echo "=== APP ==="
	@echo
	@echo "  make setup       - Instala e configura um ambiente de desenvolvimento em linux"
	@echo "  make dev         - Inicia os containers"
	@echo "  make stop        - Para os containers"
	@echo "  make logs        - Exibe logs dos containers"
	@echo "  make remove       - Remove todos os containers e volumes"
	@echo "  make sqlc        - Gera código SQLC"
	@echo "  make go-mod      - Gerencia as dependências"
	@echo "  make sw-docs     - Gera documentação do Swagger"
	@echo "  make vim         - Comandos Vim"
# Inicia os containers
dev: verifyEnv dev-sqlc dev-sw-docs dev-go-mod dev-docker-run

verifyEnv:
	@if [ ! -f .env ]; then \
		echo "$(RED)Erro$(RESET): O arquivo .env não foi encontrado na raiz do projeto."; \
		echo "      O arquivo .env-example está disponível para você criar o arquivo .env."; \
		exit 1; \
	fi

dev-sqlc:
	@echo "$(YELLOW)Gerando código SQLC...$(RESET)"
	sqlc generate

dev-sw-docs:
	@echo "$(YELLOW)Gerando documentação do Swagger...$(RESET)"
	swag init -g $(SRC_DIR)/main.go --output ./docs/swagger

dev-go-mod:
	@echo "$(YELLOW)Baixando dependências...$(RESET)"
	go mod download
	@echo "$(YELLOW)Gerenciando dependências...$(RESET)"
	go mod tidy

dev-build:
	@echo "$(YELLOW)Gerando binário...$(RESET)"
	go build -o $(BINARY_NAME) $(SRC_DIR)

dev-docker-run:
	@echo "$(YELLOW)Gerando imagem Docker...$(RESET)"
	docker-compose up -d

# Para os containers
stop:
	docker-compose down

# Exibe logs dos containers
logs:
	docker-compose logs -f

# Remove todos os containers e volumes
remove:
	docker-compose down
	docker system prune --all --volumes

# Gera código do SQLC
sqlc:
	sqlc generate