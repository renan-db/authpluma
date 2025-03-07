# Define todos os targets
.PHONY: sw

# Target swagger
sw:
	@echo "=== SWAGGER ==="
	@echo
	@echo "  swag init -g cmd/api/main.go --output docs/swagger  - Gera documentação do Swagger dentro da pasta docs/swagger"
	@echo

# Gera a documentação do Swagger
sw-docs:
	swag init -g cmd/api/main.go --output docs/swagger
