# Define todos os targets como phony
.PHONY: go go-help

# Target go
go:
	@echo
	@echo "=== COMANDOS GO ==="
	@echo
	@echo "=== Gerenciamento de Dependências ==="
	@echo "  go-mod-init                    - Inicializa um novo módulo"
	@echo "  go-mod-update                  - Atualiza dependências do projeto"
	@echo "  go-mod-vendor                  - Cria pasta vendor com dependências"
	@echo "  go-mod-verify                  - Verifica integridade das dependências"
	@echo "  go-list-u-m-all                 - Lista dependências desatualizadas"
	@echo
	@echo "=== Execução ==="
	@echo "  go-run-api                    - Executa a aplicação"
	@echo "  go-run-api-dev                - Executa em modo desenvolvimento"
	@echo "  go-run-api-prod               - Executa em modo produção"
	@echo "  go-run-api-test               - Executa em modo teste"
	@echo
	@echo "=== Build ==="
	@echo "  go-build-api                  - Compila a aplicação"
	@echo "  go-install                    - Instala o binário"
	@echo "  go-clean                      - Remove arquivos de build"
	@echo
	@echo "=== Testes ==="
	@echo "  go-test-all                     - Executa todos os testes"
	@echo "  go-test-v-all                   - Executa testes com detalhes"
	@echo "  go-test-cover-all               - Executa testes com cobertura"
	@echo "  go-test-race-all                - Testa condições de corrida"
	@echo
	@echo "=== Ferramentas ==="
	@echo "  go-fmt-all                     - Formata o código"
	@echo "  go-vet-all                      - Analisa código por erros"
	@echo "  go-doc                         - Gera documentação"
	@echo "  go-generate-all                - Executa geradores de código"
	@echo
	@echo "=== Debug ==="
	@echo "  go-tool-pprof                  - Analisa performance"
	@echo "  go-tool-trace                  - Analisa traces de execução"
	@echo "  go-tool-cover                  - Analisa cobertura de testes"
	@echo
	@echo "=== Ambiente ==="
	@echo "  go-env                         - Mostra variáveis de ambiente"
	@echo "  go-version                     - Mostra versão do Go"
	@echo "  go-help                        - Mostra ajuda do Go"
	@echo

	go-mod-update:
		go mod tidy