# Define todos os targets
.PHONY: setup

# Target setup help
setup:
	@echo "=== SETUP ==="
	@echo
	@echo "  make setup-linux - Instala dependências necessárias no Linux"
	@echo

# Instala dependências no Linux
setup-linux:
	@echo "=== VOCÊ ESTÁ PRESTES: ==="
	@echo
	@echo "- Atualizar os pacotes do Linux"
	@echo "- Instalar Docker e Docker Compose"
	@echo "- Instalar Go e configurar GOPATH"
	@echo "- Instalar Swagger"
	@echo "- Configurar permissões do Docker"
	@echo
	@read -p "Deseja prosseguir? (y/n) " answer; \
	if [ "$$answer" != "y" ]; then \
		echo "Operação cancelada"; \
		exit 1; \
	fi
	
	@echo "=== Instalando dependências no Linux ==="
	@echo "1. Atualizando pacotes..."
	sudo apt update

	@echo "2. Instalando Docker..."
	sudo apt install -y docker.io docker-compose

	@echo "3. Instalando Go..."
	sudo apt install -y golang-go

	@echo "4. Configurando GOPATH..."
	echo 'export GOPATH=$$HOME/go' >> ~/.bashrc
	echo 'export PATH=$$PATH:$$GOPATH/bin' >> ~/.bashrc
	source ~/.bashrc

	@echo "5. Instalando Swagger..."
	go install github.com/swaggo/swag/cmd/swag@latest

	@echo "6. Configurando permissões Docker..."
	sudo usermod -aG docker $$USER

	@echo "=== Setup completo! ==="
	@echo "Por favor, faça logout e login novamente para aplicar todas as alterações"
	@echo "Depois execute: make dev"