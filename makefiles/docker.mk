# Define todos os targets
.PHONY: docker

# Target docker
docker:
	@echo
	@echo "=== DOCKER ==="
	@echo
	@echo "=== Docker Compose ==="
	@echo
	@echo "  docker-compose up -d                            - Inicia todos os containers em background"
	@echo "  docker-compose down                             - Para e remove todos os containers e redes criadas pelo Compose"
	@echo "  docker-compose build                            - Constrói todas as imagens definidas no arquivo docker-compose.yml"
	@echo "  docker-compose pull                             - Baixa as versões mais recentes das imagens do Compose"
	@echo "  docker-compose push                             - Envia as imagens criadas para um registro remoto"
	@echo "  docker-compose logs                             - Exibe os logs de todos os containers gerenciados pelo Compose"
	@echo "  docker-compose logs -f <container>              - Exibe os logs de um container específico em tempo real"
	@echo "  docker-compose logs -t <container>              - Exibe os logs de um container com timestamps"
	@echo "  docker-compose logs -n <container> <number>     - Exibe os últimos <number> logs de um container"
	@echo
	@echo "=== Docker Cleaning Up ==="
	@echo
	@echo "  docker system prune                             - Remove containers parados, imagens não utilizadas, redes não utilizadas e caches de build"
	@echo "  docker system prune -a                          - Remove todos os containers e imagens que não estão sendo usados por nenhum container em execução"
	@echo "  docker system prune -f                          - Executa a limpeza sem pedir confirmação"
	@echo "  docker system prune -f -a                       - Remove todos os containers e imagens não utilizadas sem pedir confirmação"
	@echo "  docker system prune --volumes                   - Remove todos os volumes que não estão em uso por nenhum container"
	@echo "  docker system prune --all                       - Equivalente a '-a', remove containers, imagens, redes e caches de build"
	@echo "  docker system prune --all --volumes             - Remove todos os containers, imagens não utilizadas, redes não utilizadas e volumes órfãos"
	@echo "  docker container prune                          - Remove apenas containers parados"
	@echo "  docker image prune                              - Remove apenas imagens não utilizadas (dangling)"
	@echo "  docker volume prune                             - Remove apenas volumes não utilizados por nenhum container"
	@echo "  docker network prune                            - Remove redes criadas manualmente que não estão conectadas a nenhum container"
	@echo
	@echo "=== Docker Containers ==="
	@echo
	@echo "  docker ps                                       - Lista os containers em execução"
	@echo "  docker ps -a                                    - Lista todos os containers (executando e parados)"
	@echo "  docker run <image>                              - Cria e inicia um novo container a partir de uma imagem"
	@echo "  docker run -d <image>                           - Cria e inicia um novo container em background"
	@echo "  docker run --name <name> <image>                - Cria e inicia um novo container atribuindo um nome específico"
	@echo "  docker stop <container>                         - Para um container em execução"
	@echo "  docker start <container>                        - Inicia um container parado"
	@echo "  docker restart <container>                      - Reinicia um container"
	@echo "  docker rm <container>                           - Remove um container"
	@echo "  docker exec -it <container> bash                - Executa um shell interativo dentro de um container"
	@echo
	@echo "=== Docker Logs / Inspect ==="
	@echo
	@echo "  docker logs <container>                         - Exibe os logs de um container"
	@echo "  docker logs -f <container>                      - Exibe os logs de um container em tempo real"
	@echo "  docker logs -t <container>                      - Exibe os logs de um container com timestamps"
	@echo "  docker logs -n <container> <number>             - Exibe os últimos <number> logs de um container"
	@echo "  docker inspect <container/image>                - Exibe informações detalhadas sobre um container ou imagem"
	@echo "  docker stats <container>                        - Exibe estatísticas de uso de recursos de um container"
	@echo
	@echo "=== Docker Export e Import ==="
	@echo
	@echo "  docker export <container> > <file.tar>          - Exporta o sistema de arquivos de um container para um arquivo tar"
	@echo "  docker import <file.tar> <image>                - Cria uma nova imagem a partir de um arquivo tar exportado"
	@echo
	@echo "=== Docker Images ==="
	@echo
	@echo "  docker images                                   - Lista todas as imagens armazenadas localmente"
	@echo "  docker pull <image>                             - Baixa a versão mais recente de uma imagem do registro"
	@echo "  docker build -t <name>:<tag> <path>             - Constrói uma nova imagem a partir de um Dockerfile"
	@echo "  docker rmi <image>                              - Remove uma imagem local"
	@echo "  docker tag <image> <new name>:<tag>             - Cria uma nova tag para uma imagem existente"
	@echo "  docker push <name>:<tag>                        - Envia uma imagem para um registro remoto"
	@echo
	@echo "=== Docker Networks ==="
	@echo
	@echo "  docker network ls                               - Lista todas as redes do Docker"
	@echo "  docker network create <name>                    - Cria uma nova rede"
	@echo "  docker network rm <network>                     - Remove uma rede específica"
	@echo "  docker network inspect <network>                - Exibe informações detalhadas sobre uma rede"
	@echo "  docker network connect <network> <container>    - Conecta um container a uma rede existente"
	@echo "  docker network disconnect <network> <container> - Desconecta um container de uma rede"
	@echo
	@echo "=== Docker Volumes ==="
	@echo
	@echo "  docker volume ls                                - Lista todos os volumes disponíveis"
	@echo "  docker volume create <name>                     - Cria um novo volume"
	@echo "  docker volume inspect <volume>                  - Exibe informações detalhadas sobre um volume"
	@echo "  docker volume rm <volume>                       - Remove um volume específico"
	@echo "  docker run -v <volume>:/path <image>            - Monta um volume em um container durante a execução"
