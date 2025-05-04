# 📘 Comentários de Revisão de Código Go

Escrito em: 25/04/2025

> Antes de escrever código Go, leia as convenções da comunidade. </br>

> Não é regra, você pode criar suas próprias convenções.

> Este documento e baseado no repositório [layout-go](https://github.com/golang-standards/project-layout)

## 🎨 Gofmt

> Execute `gofmt` no seu código para corrigir automaticamente a maioria dos problemas de estilo mecânico. Quase todo o código Go usa `gofmt`.

```sh
gofmt -w seu_arquivo.go
```

## 📝 Sentenças de Comentário

> Comentários que documentam declarações devem ser frases completas, começando com o nome do item descrito e terminando com um ponto.

```go
// ✅ Certo
// Request representa uma solicitação para executar um comando.
type Request struct { ... }

// ❌ Errado
// representa uma solicitação para executar um comando
type Request struct { ... }
```

## 📚 Contextos

> Valores do tipo `context.Context` carregam credenciais de segurança, informações de rastreamento, prazos e sinais de cancelamento. Passe Contexts explicitamente ao longo da cadeia de chamadas de função.

```go
// ✅ Certo
func F(ctx context.Context, /* outros argumentos */) {}

// ❌ Errado
func F(/* outros argumentos */) {}
```

## 🔄 Copiando

> Evite copiar valores de tipos cujos métodos estão associados ao tipo ponteiro, `*T`.

```go
// ❌ Errado
var buf1, buf2 bytes.Buffer
buf2 = buf1 // Pode causar aliasing inesperado

// ✅ Certo
var buf1 bytes.Buffer
var buf2 bytes.Buffer // Novo buffer vazio, independente

// ⚠️ Copiando o conteúdo
var buf1 bytes.Buffer
buf1.WriteString("exemplo")

var buf2 bytes.Buffer
buf2.Write(buf1.Bytes()) // Copia só o conteúdo, não o estado

```

## 🔒 Crypto Rand

> Não use `math/rand` para gerar chaves. Use `crypto/rand.Reader`.

```go
// ✅ Certo
package main

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
)

func main() {
    b := make([]byte, 16)       // 16 bytes == 128 bits
    _, err := rand.Read(b)      // Lê bytes aleatórios seguros
    if err != nil {
        panic(err)
    }
    fmt.Println(hex.EncodeToString(b)) // Ex: "4f3c2d1a9e8b7c6d5f0e1d2c3b4a5968"
}


// ❌ Errado
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano()) // Seed é previsível
    token := fmt.Sprintf("%x", rand.Int63())
    fmt.Println("Token inseguro:", token)
}
```

## 📏 Declaração de Slices Vazios

> Prefira declarar slices vazios como `var t []string` em vez de `t := []string{}`.

```go
// ✅ Certo
var t []string // Declara um slice nil, nao ocupa memória e espaço

// ❌ Errado
t := []string{} // Cria um slice vazio, não nil, ocupa memória e espaço
```

## 📚 Comentários de Documentação

> Comentários de documentação devem ser claros e descritivos, explicando o propósito e o uso de pacotes, funções e tipos.

```go
// ✅ Certo
// Package math fornece funções básicas de matemática.
package math

// ❌ Errado
// math fornece funções básicas de matemática.
package math
```

## 🚫 Não Entre em Pânico

> Não use `panic` para tratamento de erros normais. Use valores de erro e retornos múltiplos.

```go
// ✅ Certo
func doSomething() error {
    if err := someOperation(); err != nil {
        return fmt.Errorf("falha na operação: %w", err)
    }
    return nil
}

// ❌ Errado
func doSomething() {
    if err := someOperation(); err != nil {
        panic("falha na operação")
    }
}
```

## ⚠️ Strings de Erro

> Strings de erro não devem ser capitalizadas ou terminar com pontuação.

```go
// ✅ Certo
return fmt.Errorf("algo deu errado")

// ❌ Errado
return fmt.Errorf("Algo deu errado.")
```

## 📚 Exemplos

> Exemplos de código devem ser incluídos para ilustrar o uso de pacotes e funções. </br>

> O comentário // Output: (com "O" maiúsculo) é especialmente interpretado pelo go test como o valor esperado da saída.

```go
// ✅ Certo
func ExampleSqrt() {
    fmt.Println(Sqrt(4))
    // Output: 2
}

// ❌ Errado
func ExampleSqrt() {
    fmt.Println(Sqrt(4))
    // output: 2
}
```

## 🧵 Tempo de Vida de Goroutines

> Certifique-se de que o tempo de vida das goroutines seja claro e documentado.

```go
// ✅ Certo
go func() {
    defer wg.Done()
    // lógica da goroutine
}()

// ❌ Errado
go func() {
    // lógica da goroutine
}()
```

## ⚠️ Lidar com Erros

> Sempre verifique e lide com erros retornados por funções.

```go
// ✅ Certo
if err := someFunction(); err != nil {
    log.Fatal(err)
}

// ❌ Errado
someFunction()
```

## 📥 Importações

> Evite renomear importações, exceto para evitar colisões de nomes. Organize importações em grupos.

```go
// ✅ Certo
import (
    "fmt"
    "os"

    "github.com/foo/bar"
    "rsc.io/goversion/version"
)

// ❌ Errado
import (
    "fmt"
    "os"
    "github.com/foo/bar"
    "rsc.io/goversion/version"
)
```

## 📥 Importação em Branco

> Use importações em branco para importar pacotes apenas para seus efeitos colaterais.

```go
// ✅ Certo
import _ "net/http/pprof"

// ❌ Errado
import "net/http/pprof"
```

## 📥 Importação com Ponto

> Evite usar importações com ponto, pois elas poluem o namespace global.

```go
// ❌ Errado
import . "fmt"

// ✅ Certo
import "fmt"
```

## 🚫 Erros In-Band

> Evite usar valores especiais para indicar erros. Use valores de erro explícitos.

```go
// ✅ Certo
if err := doSomething(); err != nil {
    return err
}

// ❌ Errado
if result == -1 {
    return errors.New("erro")
}
```

## 📏 Indentação do Fluxo de Erro

> Indente o fluxo de erro para melhorar a legibilidade.

```go
// ✅ Certo
if err := doSomething(); err != nil {
    return err
}
// Continue com o fluxo normal

// ❌ Errado
if err := doSomething(); err != nil {
    return err
} else {
    // Continue com o fluxo normal
}
```

## 🔠 Siglas

> Use siglas em maiúsculas em nomes de variáveis e funções.

```go
// ✅ Certo
var userID string

// ❌ Errado
var userid string
```

## 🔗 Interfaces

> Defina interfaces pequenas e focadas em um único propósito.

```go
// ✅ Certo
type Reader interface {
    Read(p []byte) (n int, err error)
}

// ❌ Errado
type ReadWriteCloser interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
    Close() error
}
```

## 📏 Comprimento da Linha

> Mantenha o comprimento das linhas em 80 caracteres ou menos para melhorar a legibilidade.

```go
// ✅ Certo
if longCondition && anotherLongCondition && yetAnotherLongCondition {
    // código
}

// ❌ Errado
if longCondition && anotherLongCondition && yetAnotherLongCondition && oneMoreCondition {
    // código
}
```

## 🔠 Caps Misturados

> Use caps misturados (camelCase) para nomes de variáveis, funções e tipos.

```go
// ✅ Certo
var userName string
func fetchUserData() {}

// ❌ Errado
var user_name string
func fetch_user_data() {}
```

## 📛 Parâmetros de Resultado Nomeados

> Use parâmetros de resultado nomeados para melhorar a clareza do código.

```go
// ✅ Certo
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = fmt.Errorf("divisão por zero")
        return
    }
    result = a / b
    return
}

// ❌ Errado
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    return a / b, nil
}
```

## 🚫 Retornos Nus

> Evite retornos nus, pois eles podem dificultar a leitura do código.

```go
// ✅ Certo
func add(a, b int) int {
    return a + b
}

// ❌ Errado
func add(a, b int) (sum int) {
    sum = a + b
    return
}
```

## 📚 Comentários de Pacote

> Inclua comentários de pacote para descrever o propósito e o uso do pacote.

```go
// ✅ Certo
// Package math fornece funções básicas de matemática.
package math

// ❌ Errado
package math
```

## 📦 Nomes de Pacotes

> Evite nomes de pacotes sem significado, como `util`, `common`, `misc`.

```go
// ✅ Certo
package chubby

// ❌ Errado
package util
```

## 📦 Passar Valores

> Prefira passar valores em vez de ponteiros para tipos pequenos e imutáveis.

```go
// ✅ Certo
func process(data string) { ... }

// ❌ Errado
func process(data *string) { ... }
```

## 📦 Nomes de Receptores

> Use nomes de receptores consistentes e significativos.

```go
type User struct {
    Name string
}

// ✅ Certo
func (u *User) GetName() string {
    return u.Name
}

// ❌ Errado
func (this *User) GetName() string {
    return this.Name
}
```

## 📦 Tipo de Receptor

> Use receptores de valor para tipos pequenos e imutáveis, e receptores de ponteiro para tipos grandes ou mutáveis.

```go
type Point struct {
    X, Y int
}

// ✅ Certo
func (p Point) Distance() int {
    return p.X*p.X + p.Y*p.Y
}

// ❌ Errado
func (p *Point) Distance() int {
    return p.X*p.X + p.Y*p.Y
}
```

## 📦 Funções Síncronas

> Prefira funções síncronas a funções assíncronas, a menos que haja uma boa razão para usar goroutines.

```go
// ✅ Certo
func fetchData() (Data, error) { ... }

// ❌ Errado
func fetchDataAsync() <-chan Data { ... }
```

## 📦 Falhas Úteis de Teste

> Inclua mensagens de falha de teste úteis que descrevam o que deu errado.

```go
// ✅ Certo
if got != want {
    t.Errorf("got %v, want %v", got, want)
}

// ❌ Errado
if got != want {
    t.Error("falha no teste")
}
```

## 📦 Nomes de Variáveis

> Use nomes de variáveis claros e descritivos.

```go
// ✅ Certo
var userName string

// ❌ Errado
var u string
```

Essas diretrizes ajudam a manter o código Go limpo, legível e consistente. Se precisar de mais detalhes ou ajustes, sinta-se à vontade para pedir!
