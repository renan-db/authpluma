# 📦 Diretrizes de Estilo para Pacotes Go

Escrito em: 25/04/2025

> Antes de escrever código Go, leia as convenções da comunidade. </br>

> Não é regra, você pode criar suas próprias convenções.

> Este documento e baseado no repositório [layout-go](https://github.com/golang-standards/project-layout)

## 📂 Organização de Pacotes

### Use Múltiplos Arquivos

> Um pacote é um diretório com um ou mais arquivos Go. Separe seu código em quantos arquivos forem necessários para uma leitura ideal.

```plaintext
- doc.go       // documentação do pacote
- headers.go   // tipos e código de cabeçalhos HTTP
- cookies.go   // tipos e código de cookies HTTP
- http.go      // implementação do cliente HTTP, tipos de requisição e resposta, etc.
```

> ⚠️ Exemplo da arquitetura usado no projeto

```go
user/
├── docs/        // Documentação do pacote
├── dto/         // Objetos de transferência de dados (Data Transfer Object)
├── entity/      // Entidades de domínio (modelos de dados principais)
├── handle/      // Funções de manipulação (handlers) para requests (HTTP)
├── helper/      // Funções auxiliares ou utilitários
├── interface/   // Interfaces (abstrações de comportamentos)
├── repository/  // Acesso e persistência de dados (banco, arquivos, etc.)
├── route/       // Definição de rotas HTTP ou de comunicação
├── usecase/     // Casos de uso que implementam a lógica de negócio

```

### Mantenha Tipos Próximos

> Mantenha os tipos próximos de onde são usados. Isso facilita a localização de um tipo por qualquer mantenedor.

```go
// headers.go
package http

// Header representa um cabeçalho HTTP.
type Header struct {...}
```

### Organize por Responsabilidade

> Em Go, organize o código por responsabilidades funcionais, em vez de criar pacotes genéricos como `models` ou `types`.

```go
package mngtservice

// User representa um usuário no sistema.
type User struct {...}

func UsersByQuery(ctx context.Context, q *Query) ([]*User, *Iterator, error)
func UserIDByEmail(ctx context.Context, email string) (int64, error)
```

### Otimize para Godoc

> Use o godoc nas fases iniciais do design da API do seu pacote para ver como seus conceitos serão renderizados na documentação.

```sh
godoc -http=:6060
```

### Forneça Exemplos para Preencher Lacunas

> Inclua exemplos para ajudar o usuário a descobrir e entender como os tipos são usados juntos.

```go
// Exemplo de uso do NewClient
func ExampleNewClient() {
    client, err := datastore.NewClient(ctx, "project-id")
    if err != nil {
        log.Fatal(err)
    }
    // Use o cliente...
}
```

### Não Exporte do Main

> Não exporte identificadores de pacotes main, pois eles não são importáveis.

```go
// Correto
package main

func main() {
    // código
}

// Errado
package main

// Exportar identificadores aqui não é necessário
```

## 📛 Nomeação de Pacotes

### Apenas Minúsculas

> Nomes de pacotes devem ser em minúsculas. Não use snake_case ou camelCase.

```go
// Correto
package http

// Errado
package HTTP
```

### Nomes Curtos, mas Representativos

> Nomes de pacotes devem ser curtos, mas únicos e representativos. Evite nomes amplos como "common" e "util".

```go
// Correto
package httputil

// Errado
package common
```

### Caminhos de Importação Limpos

> Evite expor a estrutura do seu repositório para os usuários. Alinhe-se bem com as convenções do GOPATH.

```plaintext
// Correto
github.com/user/repo/httputil

// Errado
github.com/user/repo/src/httputil
```

### Sem Plurais

> Em Go, nomes de pacotes não são plurais. Use a forma singular.

```go
// Correto
package httputil

// Errado
package httputils
```

### Renomeações Devem Seguir as Mesmas Regras

> Se você estiver importando mais de um pacote com o mesmo nome, pode renomeá-los localmente. As renomeações devem seguir as mesmas regras.

```go
import (
    gourl "net/url"
    "myother.com/url"
)
```

### Enforce Vanity URLs

> Use URLs personalizadas para servir pacotes com um domínio e caminho personalizados.

```plaintext
package datastore // import "cloud.google.com/go/datastore"
```

## 📄 Documentação de Pacotes

> Sempre documente o pacote. A documentação do pacote é um comentário de nível superior imediatamente antes da cláusula do pacote.

```go
// Package ioutil implementa algumas funções utilitárias de I/O.
package ioutil

// Command gops lista todos os processos em execução no seu sistema.
package main
```

Essas diretrizes ajudam a manter o código Go bem organizado e fácil de entender. Se precisar de mais detalhes ou ajustes, sinta-se à vontade para pedir!
