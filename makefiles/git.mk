# Define todos os targets como phony
.PHONY: git

# Target git
git:
	@echo
	@echo "=== GIT ==="
	@echo
	@echo "=== Comandos Básicos ==="
	@echo
	@echo "  git status                             - Mostra o estado atual do repositório"
	@echo "  git add <arquivo>                      - Adiciona um arquivo ao stage"
	@echo "  git add .                              - Adiciona todos os arquivos ao stage"
	@echo "  git commit -m 'mensagem'               - Cria um commit com uma mensagem"
	@echo "  git commit -m 'titulo' -m 'descricao'  - Cria um commit com título e descrição"
	@echo "  git push                               - Envia commits para o repositório remoto"
	@echo "  git pull                               - Atualiza o repositório local com as alterações do remoto"
	@echo
	@echo "  git commit --amend -m 't' -m 'd'       - Modifica a mensagem do último commit"
	@echo "  git commit --amend                     - Abre o editor para editar o último commit"
	@echo "  esc :wq                                - Salva e sai do editor de commit"
	@echo "  git log -1 --pretty=full               - Mostra os detalhes do último commit"
	@echo
	@echo "=== Branches ==="
	@echo
	@echo "  git branch                             - Lista todas as branches"
	@echo "  git branch <nome>                      - Cria uma nova branch"
	@echo "  git checkout <branch>                  - Alterna para uma branch existente"
	@echo "  git checkout -b <nome>                 - Cria e muda para uma nova branch"
	@echo "  git merge <branch>                     - Mescla uma branch com a atual"
	@echo
	@echo "=== Histórico ==="
	@echo
	@echo "  git log                                - Mostra o histórico de commits"
	@echo "  git log --oneline                      - Exibe um histórico resumido"
	@echo "  git blame <arquivo>                    - Mostra quem alterou cada linha do arquivo"
	@echo
	@echo "=== Desfazer Alterações ==="
	@echo
	@echo "  git restore <arquivo>                  - Descarta alterações não commitadas"
	@echo "  git reset --soft HEAD~1                - Desfaz o último commit mantendo as alterações"
	@echo "  git reset --hard HEAD~1                - Desfaz o último commit e descarta alterações"
	@echo "  git revert <commit>                    - Cria um commit que reverte outro commit"
	@echo
	@echo "=== Stash ==="
	@echo
	@echo "  git stash                              - Salva alterações temporariamente"
	@echo "  git stash list                         - Lista stashes salvos"
	@echo "  git stash pop                          - Recupera e remove o último stash"
	@echo "  git stash apply                        - Recupera o último stash sem removê-lo"
	@echo
	@echo "=== Tags ==="
	@echo
	@echo "  git tag                                - Lista todas as tags"
	@echo "  git tag -a v1.0 -m 'msg'               - Cria uma tag anotada"
	@echo "  git push origin v1.0                   - Envia a tag para o repositório remoto"
	@echo
	@echo "=== Remoto ==="
	@echo
	@echo "  git remote -v                          - Lista os repositórios remotos"
	@echo "  git fetch                              - Atualiza as referências do remoto"
	@echo "  git pull origin <branch>               - Atualiza a branch específica com o remoto"
	@echo "  git push origin <branch>               - Envia a branch para o repositório remoto"
	@echo
	@echo "=== Configuração ==="
	@echo
	@echo "  git config --global user.name 'Nome'   - Define o nome do usuário"
	@echo "  git config --global user.email 'Email' - Define o email do usuário"
	@echo "  git config --list                      - Lista as configurações do Git"
	@echo
