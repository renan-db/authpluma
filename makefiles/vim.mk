# Define todos os targets como phony
.PHONY: vim

# Target vim
vim:
	@echo
	@echo "=== COMANDOS VIM ==="
	@echo
	@echo "=== Modos ==="
	@echo
	@echo "  i                             - Modo de inserção (antes do cursor)"
	@echo "  a                             - Modo de inserção (depois do cursor)"
	@echo "  v                             - Modo visual (seleção)"
	@echo "  ESC                           - Volta para modo normal"
	@echo
	@echo "=== Navegação ==="
	@echo
	@echo "  h j k l                       - Esquerda, baixo, cima, direita"
	@echo "  w                             - Próxima palavra"
	@echo "  b                             - Palavra anterior"
	@echo "  0                             - Início da linha"
	@echo "  $$                             - Fim da linha"
	@echo "  gg                            - Início do arquivo"
	@echo "  G                             - Fim do arquivo"
	@echo
	@echo "=== Edição ==="
	@echo
	@echo "  x                             - Deleta caractere"
	@echo "  dd                            - Deleta linha"
	@echo "  yy                            - Copia linha"
	@echo "  p                             - Cola depois do cursor"
	@echo "  P                             - Cola antes do cursor"
	@echo "  u                             - Desfaz última ação"
	@echo "  CTRL+r                        - Refaz última ação"
	@echo
	@echo "=== Salvar e Sair ==="
	@echo
	@echo "  :w                            - Salva arquivo"
	@echo "  :q                            - Sai (se não houver mudanças)"
	@echo "  :q!                           - Sai sem salvar"
	@echo "  :wq ou :x                     - Salva e sai"
	@echo
	@echo "=== Busca ==="
	@echo
	@echo "  /palavra                      - Busca 'palavra' para frente"
	@echo "  ?palavra                      - Busca 'palavra' para trás"
	@echo "  n                             - Próxima ocorrência"
	@echo "  N                             - Ocorrência anterior"
	@echo
	@echo "=== Substituição ==="
	@echo
	@echo "  :%s/antigo/novo/g            - Substitui todas ocorrências"
	@echo "  :s/antigo/novo               - Substitui primeira ocorrência na linha"
	@echo
	@echo "=== Janelas ==="
	@echo
	@echo "  :sp arquivo                   - Abre arquivo em split horizontal"
	@echo "  :vsp arquivo                  - Abre arquivo em split vertical"
	@echo "  CTRL+w w                      - Alterna entre janelas"
	@echo "  CTRL+w h/j/k/l                - Move para janela esquerda/baixo/cima/direita"
	@echo
