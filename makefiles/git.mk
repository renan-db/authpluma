# Define todos os targets como phony
.PHONY: git git-help

# Alias para git-help
git-help: git

# Target git
git:
	@echo
	@echo "=== GIT ==="
	@echo
	@echo "=== Comandos Básicos ==="
	@echo "  git status                       - Mostra o estado atual do repositório"
	@echo "  git add <arquivo>                - Adiciona arquivos ao stage"
	@echo "  git add .                        - Adiciona todos os arquivos ao stage"
	@echo "  git commit -m 'mensagem'         - Cria um commit com uma mensagem"
	@echo "  git commit -m 't' -m 'd'         - Cria um commit com um titulo e uma descrição"
	@echo "  git push                         - Envia commits para o repositório remoto"
	@echo "  git pull                         - Atualiza o repositório local"
	@echo
	@echo "=== Branches ==="
	@echo "  git branch                       - Lista todas as branches"
	@echo "  git branch <nome>                - Cria uma nova branch"
	@echo "  git checkout <branch>            - Muda para uma branch"
	@echo "  git checkout -b <nome>           - Cria e muda para uma nova branch"
	@echo "  git merge <branch>               - Mescla uma branch com a atual"
	@echo
	@echo "=== Histórico ==="
	@echo "  git log                          - Mostra histórico de commits"
	@echo "  git log --oneline                - Mostra histórico resumido"
	@echo "  git blame <arquivo>              - Mostra quem alterou cada linha"
	@echo
	@echo "=== Desfazer Alterações ==="
	@echo "  git restore <arquivo>            - Descarta alterações não commitadas"
	@echo "  git reset --soft HEAD~1          - Desfaz último commit mantendo alterações"
	@echo "  git reset --hard HEAD~1          - Desfaz último commit e alterações"
	@echo "  git revert <commit>              - Cria commit que reverte outro commit"
	@echo
	@echo "=== Stash ==="
	@echo "  git stash                        - Salva alterações temporariamente"
	@echo "  git stash list                   - Lista stashes salvos"
	@echo "  git stash pop                    - Recupera e remove último stash"
	@echo "  git stash apply                  - Recupera último stash sem remover"
	@echo
	@echo "=== Tags ==="
	@echo "  git tag                          - Lista todas as tags"
	@echo "  git tag -a v1.0 -m 'msg'         - Cria uma tag anotada"
	@echo "  git push origin v1.0             - Envia tag para o remoto"
	@echo
	@echo "=== Remoto ==="
	@echo "  git remote -v                    - Lista repositórios remotos"
	@echo "  git fetch                        - Atualiza refs do remoto"
	@echo "  git pull origin <branch>         - Atualiza branch específica"
	@echo "  git push origin <branch>         - Envia branch para remoto"
	@echo
	@echo "=== Configuração ==="
	@echo "  git config --global user.name    - Define nome do usuário"
	@echo "  git config --global user.email   - Define email do usuário"
	@echo "  git config --list                - Lista configurações"
	@echo

# Alias para git-help
git-help: git
