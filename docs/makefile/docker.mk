# Define todos os targets
.PHONY: docker

# Target docker
docker:
	@echo
	@echo "=== DOCKER ==="
	@echo
	@echo "  docker-compose up -d                        - Inicia todos os containers em background"
	@echo "  docker-compose down                         - Para e remove todos os containers"
	@echo "  docker-compose build                        - Constrói todas as imagens"
	@echo "  docker-compose logs                         - Exibe os logs de todos os containers"
	@echo "  docker system prune                         - Remove containers e imagens não utilizados"
	@echo "  docker ps                                   - Exibe os containers em execução"
	@echo "  docker exec -it nome_do_container /bin/sh   - Acessa o container"
	@echo "  docker exec -it nome_do_container /bin/bash - Acessa o container se tiver bash"
	@echo "  "
	@echo "=== GUIA PARA LINUX COM DOCKER ==="
	@echo ""
	@echo "  docker search linux                         - Busca imagens de linux"
	@echo "  docker pull <nome_da_imagem>                - Baixa a imagem, exemplo: docker pull ubuntu"
	@echo "  docker run -d -it <nome_da_imagem> bash     - Inicia o container, exemplo: docker run -d -it ubuntu bash"
	@echo "  docker exec -it <id_container> bash         - Acessa o container, exemplo: docker exec -it ubuntu bash"
	@echo "  df -ah                                      - No linux: df=espaço em disco, -a=toda a partição, -h=formato humano (KB, MB, GB, etc)"
	@echo "  apt-get update                              - Atualiza os pacotes"
	@echo "  apt-get install curl                        - Instala o curl, serve para fazer requisições HTTP"
	@echo "  apt-get install net-tools                   - Instala o net-tools, serve para ver as portas e etc"
