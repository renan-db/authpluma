# Define todos os targets
.PHONY: git

# Target git
git:
	@echo
	@echo "=== GIT ==="
	@echo
	@echo "  git status                   - Mostra o estado atual do repositório"
	@echo "  git add <arquivo>            - Adiciona um arquivo ao stage"
	@echo "  git commit -m 'mensagem'     - Cria um commit com uma mensagem"
	@echo "  git push                     - Envia commits para o repositório remoto"
	@echo "  git pull                     - Atualiza o repositório local"
	@echo
	@echo "  git pull--rebase origin main - Atualiza o repositório local e faz rebase com a branch main""
	@echo "                               - Isso pode causar conflitos se houver alterações na branch main"
	@echo "                               - Por segurança, crie uma branch temporária antes de fazer o rebase"
	@echo "                               - git push origin main"
	@echo
	@echo "  git branch                   - Lista todas as branches"
	@echo "  git checkout <branch>        - Alterna para uma branch"
	@echo "  git merge <branch>           - Mescla uma branch com a atual"
	@echo "  git commit --amend           - Abre o editor para editar o último commit"
	@echo "  git log -1 --pretty=full     - Mostra os detalhes do último commit"
	@echo "  git rm -r --cached <path>    - Remove arquivos versionados"
	@echo
