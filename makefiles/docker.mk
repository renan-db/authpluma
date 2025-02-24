# Define todos os targets
.PHONY: docker docker-help

# Target docker
docker:
	@echo
	@echo "=== DOCKER ==="
	@echo
	@echo "=== Docker Compose ==="
	@echo
	@echo "  docker-compose up -d                        - Inicia todos os containers em background"
	@echo "  docker-compose down                         - Para e remove todos os containers"
	@echo "  docker-compose build                        - Cria todas as imagens"
	@echo "  docker-compose pull                         - Atualiza todas as imagens"
	@echo "  docker-compose push                         - Envia todas as imagens para o registro"
	@echo "  docker-compose logs                         - Exibe os logs de todos os containers"
	@echo "  docker-compose logs -f <container>          - Exibe os logs de um container em tempo real"
	@echo "  docker-compose logs -t <container>          - Exibe os logs de um container com timestamps"
	@echo "  docker-compose logs -n <container> <number> - Exibe os ultimos <number> logs de um container"
	@echo
	@echo "=== Docker Cleaning Up ==="
	@echo
	@echo "  docker system prune                     - Remove todos os containers e imagens não utilizadas"
	@echo "  docker system prune -a                  - Remove todos os containers e imagens não utilizadas"
	@echo "  docker system prune -f                  - Remove todos os containers e imagens não utilizadas"
	@echo "  docker system prune -f -a               - Remove todos os containers e imagens não utilizadas"
	@echo "  docker system prune --volumes           - Remove todos os volumes não utilizados"
	@echo "  docker system prune --all               - Remove todos os containers e imagens não utilizadas"
	@echo "  docker system prune --all --volumes     - Remove todos os containers e imagens não utilizadas e volumes não utilizados"
	@echo "  docker container prune                  - Remove todos os containers não utilizados"
	@echo "  docker image prune                      - Remove todas as imagens não utilizadas"
	@echo "  docker volume prune                     - Remove todos os volumes não utilizados"
	@echo "  docker network prune                    - Remove todas as redes não utilizadas"
	@echo
	@echo "=== Docker Containers ==="
	@echo
	@echo "  docker ps                               - Lista containers em execucao"
	@echo "  docker ps -a                            - Lista todos os containers"
	@echo "  docker run <image>                      - Inicia um novo container"
	@echo "  docker run -d <image>                   - Inicia um novo container em background"
	@echo "  docker --name <name> <image>            - Inicia um novo container com um nome"
	@echo "  docker stop <container>                 - Para um coItanner em exeiucio"
	@echo "  docker start <container>                - Inicia um coItanner parado"
	@echo "  docker restart <container>              - Reinicia um coItanner"
	@echo "  docker rm <container>                   - Removea umco tainer"
	@echo "  docker exec -it <container> bash        - Executa um comando no container"
	@echo
	@echo "=== Docker Logs / Inspect ==="
	@echo
	@echo "  docker logs <container>                 - Exibe os logs de um container"
	@echo "  docker logs -f <container>              - Exibe os logs de um container em tempo real"
	@echo "  docker logs -t <container>              - Exibe os logs de um container com timestamps"
	@echo "  docker logs -n <container> <number>     - Exibe os ultimos <number> logs de um container"
	@echo "  docker inspect <container/image>        - Exibe informacoes detalhadas sobre um container/imagem"
	@echo "  docker status <container>               - Exibe o status de um container"
	@echo
	@echo "=== Docker Export E Import ==="
	@echo
	@echo "  docker export <container> > <file.tar>  - Exporta um container para um arquivo"
	@echo "  docker import <file.tar> <image>        - Importa uma imagem de um arquivo"
	@echo
	@echo "=== Docker Images ==="
	@echo
	@echo "  docker images                           - Lista todas as imagens"
	@echo "  docker pull <image>                     - Atualiza uma imagem"
	@echo "  docker build -t <name>:<tag> <path>     - Cria uma nova imagem"
	@echo "  docker rmi <image>                      - Remnoe uma imagem"
	@echo "  docker tag <image> <new name>:<tag>     - Renomeia uma imagem"
	@echo "  docker push <name>:<tag>                - Envia uma imagem para um registro"
	@echo
	@echo "=== Docker Networks ==="
	@echo
	@echo "  docker network ls                       - Lista todas as redes"
	@echo "  docker network create <name>            - Cria uma nova rede"
	@echo "  docker network rm <network>             - Remove uma rede"
	@echo "  docker inspect <network>                - Exibe informacoes detalhadas sobre uma rede"
	@echo "  docker inspect <network> <container>    - Exibe informacoes detalhadas sobre uma rede e um container"
	@echo "  docker connect <network> <container>    - Conecta um container a uma rede"
	@echo "  docker disconnect <network> <container> - Desconecta um container de uma rede"
	@echo
	@echo "=== Docker Volumes ==="
	@echo
	@echo "  docker volume ls                        - Lista todos os volumes"
	@echo "  docker volume create <name>             - Cria um novo volume"
	@echo "  docker volume inspect <volume>          - Exibe informacoes detalhadas sobre um volume"
	@echo "  docker volume rm <volume>               - Remove um volume"
	@echo "  docker -v <volume>:/path <image>        - Monta um volume em um container"
