# Define todos os targets
.PHONY: app

# Target app
app:
	@echo
	@echo "=== APP ==="
	@echo
	@echo "  make dev        - Inicia os containers"
	@echo "  make stop       - Para os containers"
	@echo "  make logs       - Exibe logs dos containers"
	@echo "  make remove-all - Remove todos os containers e volumes"
	@echo

# Inicia os containers
dev:
	docker-compose up -d

# Para os containers
stop:
	docker-compose down

# Exibe logs dos containers
logs:
	docker-compose logs -f

docker-remove-all:
	docker system prune --all --volumes

# Gera código do SQLC
sqlc:
	sqlc generate