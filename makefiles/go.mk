# Define todos os targets como phony
.PHONY: go

# Target go
go:
	@echo
	@echo "=== COMANDOS GO ==="
	@echo
	@echo "=== Gerenciamento de Dependências ==="
	@echo ""
	@echo "  go mod init <module>                - Inicializa um novo módulo Go (executar apenas uma vez no início do projeto)"
	@echo "  go get -u ./...                     - Atualiza todas as dependências para suas versões mais recentes"
	@echo "  go mod tidy                         - Sincroniza dependências: adiciona necessárias, remove não utilizadas"
	@echo "  go mod vendor                       - Cria pasta vendor com cópia local das dependências (útil para deploys)"
	@echo "  go mod verify                       - Verifica se as dependências não foram modificadas (segurança)"
	@echo "  go list -u -m all                   - Lista todas as dependências e suas atualizações disponíveis"
	@echo
	@echo "=== Execução ==="
	@echo ""
	@echo "  go run main.go                      - Executa arquivo main.go diretamente sem gerar binário"
	@echo "  go run .                            - Executa o pacote atual (útil quando há múltiplos arquivos)"
	@echo
	@echo "=== Build ==="
	@echo ""
	@echo "  go build                            - Compila o projeto gerando um executável"
	@echo "  go install                          - Compila e instala o binário no GOPATH/bin"
	@echo "  go clean                            - Remove arquivos de compilação e cache"
	@echo
	@echo "=== Testes ==="
	@echo ""
	@echo "  go test ./...                       - Executa todos os testes do projeto recursivamente"
	@echo "  go test -v ./...                    - Executa testes com output detalhado"
	@echo "  go test -cover ./...                - Mostra porcentagem de cobertura de testes"
	@echo "  go test -race ./...                 - Detecta condições de corrida em código concorrente"
	@echo
	@echo "=== Ferramentas ==="
	@echo ""
	@echo "  go fmt ./...                        - Formata código seguindo padrão Go"
	@echo "  go vet ./...                        - Analisa código buscando erros comuns"
	@echo "  go doc <package>                    - Mostra documentação de um pacote"
	@echo "  go generate ./...                   - Executa comandos de geração de código (//go:generate)"
	@echo
	@echo "=== Debug ==="
	@echo ""
	@echo "  go tool pprof <arquivo>             - Analisa profile de CPU/memória"
	@echo "  go tool trace <arquivo>             - Visualiza trace de execução"
	@echo "  go tool cover -html=coverage.out    - Abre relatório de cobertura no navegador"
	@echo
	@echo "=== Ambiente ==="
	@echo ""
	@echo "  go env                              - Lista todas as variáveis de ambiente Go"
	@echo "  go version                          - Mostra versão do Go instalada"
	@echo "  go help                             - Mostra ajuda geral dos comandos"
	@echo

go-mod-update:
	go mod tidy