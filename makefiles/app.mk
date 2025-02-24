# Define todos os targets
.PHONY: app app-help

# Target app
app:
	@echo
	@echo "=== APP ==="
	@echo
	@echo "  make help   - Exibe esta mensagem de ajuda"
	@echo "  make dev    - Inicia o ambiente de desenvolvimento"
	@echo "  make prod   - Inicia o ambiente de produção"
	@echo "  make test   - Inicia o ambiente de teste"
	@echo "  make stop   - Para e remove todos os containers"
	@echo

# Inicia o ambiente de desenvolvimento
# 1. Inicia os containers Docker em background (-d)
# 2. Executa a aplicação em modo desenvolvimento
dev:
	GO_ENV=dev docker-compose up -d
	go run cmd/api/main.go dev

# Inicia o ambiente de produção
# 1. Inicia os containers Docker em background (-d)
# 2. Executa a aplicação em modo produção
prod:
	GO_ENV=prod docker-compose up -d
	go run cmd/api/main.go prod

# Inicia o ambiente de teste
# 1. Inicia os containers Docker em background (-d)
# 2. Executa a aplicação em modo teste
test:
	GO_ENV=test docker-compose up -d
	go run cmd/api/main.go test

# Para e remove todos os containers
# Útil para limpar o ambiente ou reiniciar do zero
stop:
	docker-compose down
