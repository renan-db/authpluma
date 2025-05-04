# 📚 Guia de Comentários — Projeto Go

Este guia define as boas práticas para escrever comentários no projeto, seguindo o padrão oficial do [GoDoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc).

## 🎯 Objetivo

- Facilitar a leitura e manutenção do código.
- Permitir a geração automática de documentação.
- Manter um padrão único e profissional em todo o projeto.

---

## ✍️ Regras Básicas

- Comentários devem iniciar com o **nome do pacote, struct, interface ou função**.
- Utilize **frases diretas** no **presente do indicativo**.
- **Não use vírgulas** após o nome.
- **Não adicione ponto final** em comentários de uma linha.
- **Máximo de 80 caracteres** por linha.

---

## 📦 Comentando Pacotes

**Formato:**

```go
// Package nome define [descrição breve do que o pacote faz].
package nome
```

**Exemplo:**

```go
// Package handler define os handlers para o módulo de usuário.
package handler
```

---

## 🏛️ Comentando Structs

**Formato:**

```go
// NomeStruct representa [o que a struct representa].
type NomeStruct struct { ... }
```

**Exemplo:**

```go
// UserEntity representa a entidade de usuário no sistema.
type UserEntity struct {
    ID int
    Name string
}
```

---

## 🧩 Comentando Interfaces

**Formato:**

```go
// NomeInterface define [a responsabilidade da interface].
type NomeInterface interface { ... }
```

**Exemplo:**

```go
// Binder define a interface para realizar o bind de dados em uma estrutura.
type Binder interface {
    Bind(target interface{}) error
}
```

---

## 🔧 Comentando Funções e Métodos

**Formato:**

```go
// NomeFunção [descrição objetiva da ação realizada].
func NomeFunção(...) { ... }
```

**Exemplos:**

```go
// NewUserHandler cria uma nova instância de userHandler.
func NewUserHandler(...) *userHandler { ... }
```

```go
// Create processa a requisição para criar um novo usuário.
func (h *userHandler) Create(c echo.Context) error { ... }
```

---

## ⚡ Resumo Rápido

| Elemento  | Exemplo de Comentário                           |
|-----------|--------------------------------------------------|
| Pacote    | `Package handler define os handlers...`          |
| Struct    | `UserEntity representa a entidade...`            |
| Interface | `Binder define a interface para binding...`      |
| Função    | `NewUserHandler cria uma nova instância...`      |
| Método    | `Create processa uma requisição...`              |

---

## 📌 Observações Finais

- Mantenha os comentários atualizados conforme o código evolui.
- Prefira clareza em vez de complexidade.
- Comentários são uma ferramenta de comunicação entre desenvolvedores.

---

# 🚀 Vamos construir um projeto limpo e fácil de entender!
