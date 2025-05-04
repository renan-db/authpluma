# Inclui outros Makefiles
include docs/makefile/app.mk
include docs/makefile/docker.mk
include docs/makefile/git.mk
include docs/makefile/go.mk
include docs/makefile/vim.mk
include docs/makefile/swagger.mk
include docs/makefile/setup.mk

# Define todos os targets como phony (make)
.PHONY: help

# Comando padrão
.DEFAULT_GOAL := help

# Menu de ajuda principal
help:
	@echo
	@echo "=== COMANDOS DISPONÍVEIS ==="
	@echo
	@echo " make app     - Comandos principais da aplicacao"
	@echo " make docker  - Comandos Docker"
	@echo " make git     - Comandos Git"
	@echo " make go      - Comandos Go"
	@echo " make setup   - Comandos Setup"
	@echo " make swagger - Comandos Swagger"
	@echo " make vim     - Comandos Vim"
	@echo
	
