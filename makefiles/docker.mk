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
	
	@echo