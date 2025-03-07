# Define todos os targets
.PHONY: sw sw-help

# Target swagger
sw:
	@echo "=== SWAGGER ==="
	@echo
	@echo "  make sw-docs - Gera documentacao do swagger"
	@echo

# Inicia os containers
sw-docs:
	swag init -g cmd/api/main.go --output docs/swagger
