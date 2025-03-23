# Define todos os targets
.PHONY: go

# Target go
go:
	@echo
	@echo "=== COMANDOS GO ==="
	@echo
	@echo "  go mod init <module>        - Inicializa um novo módulo Go"
	@echo "  go get -u ./...             - Atualiza todas as dependências"
	@echo "  go mod tidy                 - Sincroniza dependências"
	@echo "  go run main.go              - Executa o arquivo main.go"
	@echo "  go build                    - Compila o projeto"
	@echo "  go test ./...               - Executa todos os testes"
	@echo