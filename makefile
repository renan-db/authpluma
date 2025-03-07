# Inclui outros Makefiles
include makefiles/app.mk
include makefiles/docker.mk
include makefiles/git.mk
include makefiles/go.mk
include makefiles/vim.mk
include makefiles/swagger.mk
include makefiles/setup.mk

# Define todos os targets como phony (make)
.PHONY: help

# Comando padrão
.DEFAULT_GOAL := help

# Menu de ajuda principal
help:
	@echo
	@echo "=== COMANDOS DISPONÍVEIS ==="
	@echo
	@echo "  make app    - Comandos da aplicacao"
	@echo "  make docker - Comandos Docker"
	@echo "  make git    - Comandos Git"
	@echo "  make vim    - Comandos Vim"
	@echo
	@echo "=== COMANDOS MAIS USADOS ==="
	@echo
	@echo " make setup-linux     - Instala dependencias no Linux"
	@echo " make go-mod-update   - Atualiza dependencias do projeto"
	@echo " make sw-docs         - Gera documentacao do swagger"