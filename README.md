"# authpluma" 
Autenticador centralizado desenvolvido em Go com foco em escalabilidade e segurança.

## Endpoints

- `http://localhost:8080/` - Página inicial, execução local.

## Funcionalidades

- Autenticação de usuários
- Autorização de acesso a recursos
- Gerenciamento de sessões

## Comandos Essenciais

- `go mod tidy` Instala e organiza as dependências do projeto.
- `go run ./cmd/api/main.go` Executa a aplicação.
- `go test ./...` Executa os testes do projeto.
- `go build -o app ./cmd/api` Compila a aplicação.
- `go list -u -m all` Lista as dependências desatualizadas.

- `go run cmd/api/main.go dev` Desenvolvimento (usa .env.dev)
- `go run cmd/api/main.go prod` Produção (usa .env.prod)
- `go run cmd/api/main.go test` Teste (usa .env.test)
