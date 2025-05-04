# 📘 Introdução ao Go - Estilo e Idiomas Idiomáticos	

Escrito em: 23/04/2025

Go é uma linguagem moderna com foco em simplicidade, eficiência e clareza.

> Antes de escrever código Go, leia as convenções da comunidade. </br>

> Não é regra, você pode criar suas próprias convenções.

> Este documento e baseado no repositório [layout-go](https://github.com/golang-standards/project-layout)

## 📏 Convenções da Comunidade


- **Arquivos de errors**: Arquivo com responsabilidade de lidar com erros deve iniciar com "error" ou "errors" para indicar claramente que o arquivo lida com erros. Ex: `errorsuser.go` ou `errorsuser.go`
  
- **Nomes de arquivos**: Nome do arquivo deve seguir o padrão de nomenclatura `snake_case` e `lowercase` para consistência e legibilidade. Ex: `user_entity.go` ou `user_usecase.go`

- **Módulos/Camadas**: Nome do módulo/camada deve seguir o padrão de nomenclatura `lowercase` para manter a simplicidade e evitar confusões. Ex: `module/user/` ou `module/user/usecase`

- **Declaração de retorno**: Deve ser idiomática, utilizando múltiplos retornos quando necessário, e sempre retornando um valor e um erro quando aplicável. 
</br>
Ex: </br> ✅`func (u *User) GetName() (string, error)` </br>❌`func (u *User) GetName() string`

- **Formatação automática**: Use `gofmt` para garantir que o código siga o estilo idiomático do Go. Isso elimina discussões sobre estilo de código e garante que todo código Go siga o mesmo padrão.</br>
⚠️Use outro linter após usar o `gofmt` para garantir que o código siga suas convenções.
##  Conceito
> Para escrever Go bem, é importante seguir seu estilo idiomático e convenções.

Go não é Java nem C++. </br>
Evite traduções literais. Escreva "em Go".

## 📚 Exemplos Oficiais

> Use os códigos-fonte da biblioteca padrão como referência:

- https://pkg.go.dev
- Os exemplos citados aqui são direto ao ponto, sem enrolação.

## 🎨 Formatação com `gofmt`

> A formatação não é debatida: use `gofmt`.

✅ Vantagens
- Aplica o estilo idiomático do Go.
- Elimina discussões sobre estilo de código.
- Garante que todo código Go siga o mesmo padrão.
- Facilita a leitura e manutenção do código.
- É integrado na maioria dos editores (VS Code, GoLand, etc).
- `gofmt -w meu_arquivo.go` para formatar um arquivo.
- `go fmt ./...` para formatar todos os arquivos.

</br>

```go
// ✅ Correto: Comentários alinhados

type T struct {
    name    string // nome do objeto
    value   int    // valor do objeto
}
```

```go
// ❌ Errado: desalinhado, não padronizado

type T struct {
    name string // nome do objeto
    value int // valor do objeto
}
```

## 🔧 Regras-chave de Formatação

### Indentação
```go
// ✅ Correto: Use TAB (não espaços)
if cond {
	fmt.Println("ok")
}
```

### Parênteses

> Evite parênteses extras em if/for/switch. </br>
> O idioma Go presa a simplicidade

```go
// ✅ Correto
for i := 0; i < 10; i++ {
	fmt.Println(i)
}

// ❌ Errado
for (i := 0; i < 10; i++) {
	fmt.Println(i)
}
```

```go
// ✅ Correto
if cond <= 10 {
	fmt.Println("ok")
}

// ❌ Errado
if (cond <= 10) {
	fmt.Println("ok")
}
```

```go
// ✅ Correto
switch cond {
	case 1:
		fmt.Println("ok")
	case 2:
		fmt.Println("ok")
}

// ❌ Errado
switch (cond) {
	case 1:
		fmt.Println("ok")
	case 2:
		fmt.Println("ok")
}
```

### Linha longa

> Quebre linhas longas com identação extra.

```go
// ✅ Correto
if longExpression ||
	anotherLongExpression {
	fmt.Println("OK")
}

// ❌ Errado
if longExpression || anotherLongExpression {
	fmt.Println("OK")
}

// ⚠️ Apesar de errado, é comum ver esse tipo de código.
// ⚠️ Em alguns casos, é preferível manter a linha longa.
// ⚠️ Mas não é idiomático.
if longExpression || anotherLongExpression {
	fmt.Println("OK")
}
```

## 💡 Dica final

 > Instale o `gofmt` ou configure seu editor para aplicar a formatação automaticamente ao salvar.

```sh
go fmt ./...
```

## 🏷️ Nomes

### Getters

> Em Go, não é necessário prefixar métodos de acesso com "Get" ou "Set". O nome do método deve descrever claramente o que ele faz.

```go
// ✅ Correto
type Person struct {
    Name string
}

func (p *Person) Name() string {
    return p.Name
}

// ❌ Errado
// ⚠️ Mas não é idiomático
// ⚠️ As vezes é apropriado usar getters/setters
func (p *Person) GetName() string {
    return p.Name
}
```

### Nomes de interfaces

> Interfaces em Go são nomeadas com base no comportamento que descrevem. Use o sufixo "er" para indicar a ação.

```go
// ✅ Correto
type Reader interface {
    Read(p []byte) (n int, err error)
}

// ❌ Errado
type DataReader interface {
    ReadData(p []byte) (n int, err error)
}
```

```go
// ✅ Correto
type Voicer interface {
    Speak(text string) (n int, err error)
}

// ❌ Errado
type Voice interface {
    Speak(text string) (n int, err error)
}
```

### MixedCaps

> Use o estilo MixedCaps ou camelCase para nomes de variáveis, funções e tipos. Evite underscores.

```go
// ✅ Correto
var userName string
func fetchUserData() {}

// ✅ Correto
// ⚠️ Primeira letra maiúscula quando é exportação
var UserName string
func FetchUserData() {}

// ❌ Errado
var user_name string
func fetch_user_data() {}
```

### Nomes de arquivos

> Use nomes de arquivos que refletam o que eles contêm.

```go
// ✅ Correto

`user.go` // tudo minúsculo
`user_entity.go` // sufixo para indicar partes específicas
`user_usecase.go` // sufixo para indicar partes específicas
`user_test.go`    // sufixo para indicar testes

// ❌ Errado

`User.go` // nunca use maiúsculas
`user_entity.go`  // sem underline
```

## 🔄 Estruturas de Controle

### if

> Em Go, a estrutura `if` `for` `switch` não requer parênteses ao redor da condição. É comum declarar variáveis dentro do `if` `for` `switch`.

```go
// ✅ Correto
if x := x < y {
    fmt.Println("x é menor que y")
}

// ❌ Errado
if (x < y) {
    fmt.Println("x é menor que y")
}
```	

```go
// ✅ Correto
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// ❌ Errado
for (i := 0; i < 10; i++) {
    fmt.Println(i)
}
```

```go
// ✅ Correto
switch day {
case "Monday":
    fmt.Println("Início da semana")
case "Friday":
    fmt.Println("Quase fim de semana")
default:
    fmt.Println("Dia comum")
}

// ❌ Errado
switch (day) {
case "Monday":
    fmt.Println("Início da semana")
case "Friday":
    fmt.Println("Quase fim de semana")
default:
    fmt.Println("Dia comum")
}
```

### Redeclaration and reassignment

> Go permite a reatribuição de variáveis usando `:=` dentro de blocos de controle, mas apenas se pelo menos uma das variáveis for nova.

```go
// ✅ Correto
x, y := 1, 2 // x, y são novos
x, z := 3, 4 // x já existe, mas z é novo ⇒ válido!

// ❌ Errado
x, y := 5, 6 // x e y já existem ⇒ inválido!
```

### For

> O laço `for` é a única estrutura de loop em Go. Pode ser usado de várias formas, incluindo como um `while`.

```go
// ✅ loop clássico
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// ✅ loop tipo while
for x < y {
    x *= 2
}

// ✅ loop infinito
for {
    fmt.Println("loop infinito")
}
```

### For com tag

> O `for` com tag é uma forma de loop que pode ser usado para iterar sobre uma coleção de dados.

```go
// ✅ loop com tag
// saida: 0 1 2 3 4
minhaTag:
for i := 0; i < 10; i++ {
    if i == 5 {
        break // sai do loop
    }
    fmt.Println(i)
}

// ✅ loop com tag
// saida: 0 1 2 3 4 6 7 8 9
// pulou o 5 e continuou o loop
minhaTag:
for i := 0; i < 10; i++ {
    fmt.Println(i)
    if i == 5 {
        continue minhaTag
    }
}
```
### Switch

> O `switch` em Go é mais flexível que em outras linguagens. Não precisa de `break` para sair do case.

```go
// ✅ Correto
switch day {
case "Monday":
    fmt.Println("Início da semana")
case "Friday":
    fmt.Println("Quase fim de semana")
default:
    fmt.Println("Dia comum")
}

// ❌ Errado
switch day {
case "Monday":
    fmt.Println("Início da semana")
    break // não é necessário, pode até dar warning 
case "Friday":
    fmt.Println("Quase fim de semana")
    break // não é necessário, pode até dar warning
}
```

### Type switch

> Use `type switch` para determinar o tipo de uma interface.

```go
// ✅ Correto
var i interface{} = "hello"

switch v := i.(type) {
case string:
    fmt.Printf("string %s\n", v)
case int:
    fmt.Printf("int %d\n", v)
default:
    fmt.Printf("outro tipo %T\n", v)
}
```
```go
// ✅ ⚠️ Correto, mas verboso
var i interface{} = "hello"

if v, ok := i.(string); ok {
    fmt.Printf("string %s\n", v)
} else if v, ok := i.(int); ok {
    fmt.Printf("int %d\n", v)
} else {
    fmt.Printf("outro tipo %T\n", v)
}
```

## 🔧 Funções

### Declaração de Funções

> Em Go, as funções são declaradas usando a palavra-chave `func`, seguida pelo nome da função, parâmetros e tipo de retorno.

> 💡💡💡 Dica: Escolha um padrão para seu projeto e siga ele. </br>
> 💡💡💡 Se necessário use algum linter para validar a sintaxe.


```go
// ✅ Tipos de parâmetros separados
func add(a int, b int) int {
    return a + b
}

// ✅ Tipos de parâmetros agrupados
func add(a, b int) int {
    return a + b
}

// ✅ Tipos de parâmetros agrupados e retorno nomeado
func add(a, b int) (soma int) {
    soma = a + b
    return
}

// ✅ Tipos de parâmetros agrupados e retorno nomeado
func add(a, b int) int {
    soma := a + b
    return soma
}

// ✅ Tipos de parâmetros agrupados e retorno nomeado e erro
// ⚠️ Não é idiomático, mas é útil para retornar um erro
// ⚠️O erro é retornado como segundo valor
func add(a, b int) (soma int, erro error) {
    soma = a + b
    return
}

```

### Funções Nomeadas e Anônimas

> Funções podem ser nomeadas ou anônimas. Funções anônimas são úteis para funções de callback ou quando uma função precisa ser passada como argumento.

```go
// ✅ função nomeada
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// ✅ função anônima
greet := func(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
greet("World")
```

### Retorno Múltiplo

> Funções em Go podem retornar múltiplos valores, o que é útil para retornar um resultado e um erro, por exemplo.

```go
// ✅ 100% idiomático
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    return a / b, nil
}

res, err := divide(10, 5)
if err != nil {
    fmt.Println("Erro:", err)      // divisão por zero
} else {
    fmt.Println("Resultado:", res) // 2
}
```	

### Funções Variádicas

> Funções variádicas aceitam um número variável de argumentos. Use `...` para indicar que a função aceita múltiplos argumentos do mesmo tipo.

```go
// ✅ 100% idiomático
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

fmt.Println(sum(1, 2, 3, 4)) // 10
```

### Defer

> A palavra-chave `defer` adia a execução de uma função até que a função que a contém retorne. É útil para garantir que recursos sejam liberados.

```go
// ✅ 100% idiomático
func readFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    // Garante que o arquivo será fechado mesmo que ocorra um erro
    defer file.Close()
}
```

## 📅 Manipulação de Datas

### Pacote time

> O pacote `time` do Go fornece funcionalidades para manipulação de datas e horas. É a principal ferramenta para lidar com datas em Go.

```go
import "time"

// ✅ 100% idiomático

// 2025-04-23 15:04:05
now := time.Now()

// Data e hora atual: 2025-04-23 15:04:05
fmt.Println("Data e hora atual:", now) 
```

```go
// ✅ 100% idiomático

// 23-04-2025 15:04:05
formatted := now.Format("30-02-2006 15:04:05")

// Data formatada: 23-04-2025 15:04:05
fmt.Println("Data formatada:", formatted)
```

### Parse e Format

> Use `time.Parse` para converter strings em objetos `time.Time` e `Format` para converter objetos `time.Time` em strings.

```go
// ✅ Parse de string para time.Time
dateStr := "2023-10-15"

date, err := time.Parse("2006-01-02", dateStr)

if err != nil {
    fmt.Println("Erro ao fazer parse da data:", err)
} else {
    // Data:2023-10-15 00:00:00 +0000 UTC
    fmt.Println("Data:", date)
}
```
```go	
// ✅ Format de time.Time para string
formattedDate := date.Format("02/01/2006")

// Data formatada: 15/10/2023
fmt.Println("Data formatada:", formattedDate)
```

### Duração e Intervalos

> O tipo `time.Duration` representa a diferença entre dois instantes. Use para calcular intervalos de tempo.

```go
// ✅ Calcular a diferença entre duas datas
// time.Date(ano, mês, dia, hora, minuto, segundo, nanosegundo, localização)


// 2023-10-01 00:00:00 +0000 UTC 
start := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)

// 2023-10-15 00:00:00 +0000 UTC
end := time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC)

// 14d0h0m0s
duration := end.Sub(start)

// Duração em dias: 14
fmt.Println("Duração em dias:", duration.Hours()/24)
```

### Timezones

> Manipule fusos horários usando `time.Location`. Isso é importante para aplicações globais.

```go
// ✅ Converter para um fuso horário específico

// "America/New_York"
// "America/Sao_Paulo"
// "Europe/London"
// "Asia/Tokyo"
// "Australia/Sydney"
// "Asia/Dubai"
// "Africa/Johannesburg"
loc, err := time.LoadLocation("America/New_York")

if err != nil {
    fmt.Println("Erro ao carregar localização:", err)
} else {
    locTime := now.In(loc)
    fmt.Println("Hora em Nova York:", locTime)
}
```

## 📊 Manipulação de Dados

### Alocação com `new`

> A função `new` aloca memória, mas não inicializa a memória além de definir como zero. Retorna um ponteiro para o tipo.

```go
✅ Criando um ponteiro para um inteiro
p := new(int)   // p é um ponteiro para um int
fmt.Println(*p) // Saída: 0
```

```go
✅ Criando um ponteiro para um array de inteiros
p := new([5]int)   // p é um ponteiro para um []int

fmt.Println(p)   // &[0 0 0 0 0]
fmt.Println(*p)  // [0 0 0 0 0]
// p é o ponteiro: mostra o endereço + conteúdo
// *p é o array em si
```

### Alocação com `make`

> Use `make` para alocar e inicializar slices, maps e channels. Retorna um valor inicializado (não um ponteiro).

```go
✅ Criando um slice de inteiros
s := make([]int, 5) // s= [0 0 0 0 0]
```	

```go
✅ Criando um map de strings para inteiros
m := make(map[string]int) // m= map[]
m["key"] = 42 // m= {"key": 42}
m["key2"] = 43 // m= {"key": 42, "key2": 43}
```	

```go
✅ Criando um map de strings para inteiros
estoque := make(map[string]int)
estoque["pão"] = 50
estoque["leite"] = 30
fmt.Println(estoque["pão"]) // Saída: 50
fmt.Println(estoque["leite"]) // Saída: 30
```

```go
✅ Criando um map de strings para inteiros e atribuindo valores
contagem := make(map[string]int)
frase := []string{"gato", "cachorro", "gato", "pássaro"}

for _, palavra := range frase {
    contagem[palavra]++
}

fmt.Println(contagem)
// Saída: map[cachorro:1 gato:2 pássaro:1]
```

> Quando criamos um slice com tamanho fixo, o go aloca a memória para o slice. </br>
> Quando criamos um slice flexível, também é alocado na memória. </br>
> Se a memória alocada não for suficiente, o go irá copiar o slice para uma nova posição na memória. </br>
> Se a slice for de tamanho fixo, crie com make ou new. </br>
> A troca dos dados para uma nova posição na memória perde performance.

```go
// ✅ Criando um slice flexível de inteiros
// Sem manke ou new
var s = []int{} // s= []
s = append(s, 1, 2, 3) // s= [1 2 3]
```

### Construtores e literais compostos

> Use literais compostos para inicializar structs, arrays, slices e maps. Construtores são funções que retornam um novo objeto.

```go
// ✅ Criando um struct com literais compostos
type Point struct {
    X, Y int
}
p := Point{X: 1, Y: 2}

func NewPoint(x, y int) *Point {
    return &Point{X: x, Y: y}
}
p := NewPoint(1, 2)
```

### Arrays

> Arrays têm tamanho fixo e são raramente usados diretamente. Prefira slices para flexibilidade.

```go
// ✅ Criando um array de inteiros com tamanho 3
var a [3]int // a= [0 0 0]
a[0] = 1 // a= [1 0 0]
```

### Slices

> Slices são mais comuns que arrays. Eles são dinâmicos e podem crescer conforme necessário.

```go
// ✅ Criando um slice de inteiros
s := []int{1, 2, 3} // s= [1 2 3]
s = append(s, 4)    // s= [1 2 3 4]
```

### Slices bidimensionais

> Slices podem ser aninhados para criar estruturas bidimensionais.

```go
// ✅ Criando um slice bidimensional de inteiros
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
// matrix= [[1 2 3] [4 5 6] [7 8 9]]
// matrix[0]= [1 2 3]
// matrix[0][0]= 1
```

### Maps

> Maps são coleções de pares chave-valor. Use `make` para inicializar.

```go
// ✅ Criando um map de strings para inteiros
m := make(map[string]int)
m["key"] = 42
m["key2"] = 43
// m= {"key": 42, "key2": 43}
// m["key"]= 42
// m["key2"]= 43
```

### Impressão

> Use `fmt` para imprimir valores. `Println` adiciona uma nova linha, `Printf` formata a saída.

```go
// ✅ Imprimindo um valor
// Hello, World!
fmt.Println("Hello, World!")
```

```go
// ✅ Imprimindo um valor formatado
// Valor: 42
fmt.Printf("Valor: %d\n", 42)
```

### Append

> Use `append` para adicionar elementos a um slice. `append` retorna um novo slice.

```go
// ✅ Adicionando elementos a um slice
s := []int{1, 2, 3} // s= [1 2 3]
s = append(s, 4, 5) // s= [1 2 3 4 5]
```

## 🚀 Inicialização

### Constantes

> Constantes em Go são valores imutáveis que são definidos usando a palavra-chave `const`. Elas podem ser de tipos básicos como `int`, `float`, `string`, etc.

```go
// ✅ Definindo uma constantes
const Pi = 3.14
const Greeting = "Hello, World!"

const number = int32(10) // 32 bits
const number = int64(10) // 64 bits
const number = 10        // Arquitetura decide o tipo

```

### Variáveis

> Variáveis em Go são declaradas usando a palavra-chave `var`. Elas podem ser inicializadas com um valor ou deixadas com o valor zero do tipo.

```go
// ✅ Declarando variáveis
var x int = 10
var y = 20   // Tipo inferido como int
var z string // '' string vazia
var w int    // 0 zero como inteiro
```

> Você também pode usar a declaração curta `:=` para declarar e inicializar variáveis dentro de funções.

```go
// ✅ Declarando variáveis com tipo inferido
a := 5
b := "Go"
```

### A função `init`

> A função `init` é uma função especial em Go que é executada automaticamente antes da execução da função `main`. Ela é usada para inicializar variáveis ou executar código de configuração.

> A função `init` pode ser definida múltiplas vezes em um mesmo pacote, e todas serão executadas na ordem em que aparecem no arquivo.

```go
var config string

func init() {
    config = "Configuração inicializada"
    fmt.Println(config)
}

func main() {
    fmt.Println("Função principal")
}
```

## 🛠️ Métodos

### Ponteiros vs. Valores

> Em Go, métodos podem ser definidos em tipos de valor ou tipos de ponteiro. A escolha entre usar um ponteiro ou um valor como receptor de método afeta como o método pode ser chamado e se ele pode modificar o estado do objeto.

#### Métodos com Receptores de Valor

> Quando um método tem um receptor de valor, ele opera em uma cópia do valor original. Isso significa que qualquer modificação feita ao receptor dentro do método não afeta o valor original.

```go
type Rectangle struct {
    Width, Height float64
}

// Método com receptor de valor
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

rect := Rectangle{Width: 10, Height: 5}
fmt.Println("Área:", rect.Area()) // Saída: Área: 50
```

#### Métodos com Receptores de Ponteiro

> Quando um método tem um receptor de ponteiro, ele pode modificar o valor original. Isso é útil para métodos que precisam alterar o estado do objeto.

> **Dica**: Use receptores de ponteiro quando o método precisa modificar o estado do objeto ou quando o objeto é grande e a cópia seria cara em termos de desempenho. Use receptores de valor quando o método não precisa modificar o estado do objeto e o objeto é pequeno.

```go
type Circle struct {
    Radius float64
}

// Método com receptor de ponteiro
func (c *Circle) Scale(factor float64) {
    c.Radius *= factor
}

circle := Circle{Radius: 5}
circle.Scale(2)
fmt.Println("Novo raio:", circle.Radius) // Saída: Novo raio: 10
```

## 🔗 Interfaces e Outros Tipos

### Interfaces

> Interfaces em Go definem um conjunto de métodos que um tipo deve implementar. Elas são usadas para definir comportamentos comuns.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### Conversões

> Em Go, você pode converter tipos explicitamente. Isso é útil quando você precisa tratar um tipo como outro.

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

### Conversões de Interface e Aserções de Tipo

> Use asserções de tipo para converter uma interface para seu tipo subjacente. Isso é útil quando você precisa acessar métodos ou campos específicos de um tipo.

```go
var i interface{} = "hello"

// Asserção de tipo
s, ok := i.(string)
if ok {
    fmt.Println(s) // Saída: hello
} else {
    fmt.Println("não é uma string")
}
```

### Generalidade

> Interfaces permitem que você escreva funções que podem operar em qualquer tipo que implemente a interface, promovendo a generalidade e a reutilização de código.

```go
func PrintAll(values []interface{}) {
    for _, value := range values {
        fmt.Println(value)
    }
}
```

### Interfaces e Métodos

> Métodos podem ser definidos em tipos que implementam interfaces, permitindo que esses tipos satisfaçam as interfaces.

```go
type Stringer interface {
    String() string
}

type Person struct {
    Name string
}

func (p Person) String() string {
    return p.Name
}

var p Stringer = Person{Name: "Alice"}
fmt.Println(p.String()) // Saída: Alice
```

## 🔍 O Identificador em Branco

### O Identificador em Branco em Atribuições Múltiplas

> O identificador em branco `_` é usado para ignorar valores em atribuições múltiplas. Isso é útil quando você precisa de apenas alguns dos valores retornados por uma função.

```go
func getData() (int, int, int) {
    return 1, 2, 3
}

a, _, c := getData()
fmt.Println(a, c) // Saída: 1 3
```

### Importações e Variáveis Não Utilizadas

> Em Go, importações e variáveis não utilizadas causam erros de compilação. O identificador em branco pode ser usado para evitar esses erros ao ignorar importações ou variáveis não utilizadas.

```go
import (
    "fmt"
    "os"
)

func main() {
    _, err := os.Open("file.txt")
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo")
    }
}
```

### Importação para Efeito Colateral

> Às vezes, um pacote é importado apenas para garantir que seu `init` seja executado. Use o identificador em branco para importar pacotes apenas para seus efeitos colaterais.

```go
import (
    _ "net/http/pprof"
)
```

### Verificações de Interface

> O identificador em branco pode ser usado para garantir que um tipo implementa uma interface em tempo de compilação.

```go
type Stringer interface {
    String() string
}

type MyType struct{}

func (m MyType) String() string {
    return "MyType"
}

// Verificação de interface
var _ Stringer = (*MyType)(nil)
```

## 📦 Embedding

### Embedding de Structs

> Em Go, você pode "embed" (incorporar) um struct dentro de outro. Isso permite que o struct incorporado herde os métodos e campos do struct pai, promovendo a reutilização de código.

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return "..."
}

type Dog struct {
    Animal // Embedding
    Breed string
}

func main() {
    dog := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
    fmt.Println(dog.Name)   // Saída: Rex
    fmt.Println(dog.Speak()) // Saída: ...
}
```

### Embedding de Interfaces

> Interfaces também podem ser incorporadas em outras interfaces. Isso permite que uma interface herde os métodos de outra, criando interfaces compostas.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}

func main() {
    var rw ReadWriter
    // rw agora pode usar métodos de Reader e Writer
}
```

### Vantagens do Embedding

> O embedding promove a composição sobre a herança, um princípio importante em Go. Ele permite que você crie tipos mais complexos a partir de tipos mais simples, sem a necessidade de uma hierarquia de classes.

- **Reutilização de Código**: Permite que você reutilize métodos e campos de structs ou interfaces incorporadas.
- **Flexibilidade**: Facilita a criação de tipos compostos sem a rigidez da herança tradicional.
- **Simplicidade**: Mantém o código simples e fácil de entender, seguindo o princípio KISS (Keep It Simple, Stupid).

## ⚙️ Concorrência

### Compartilhar por Comunicação

> Em Go, a concorrência é gerenciada através do compartilhamento de dados por comunicação, em vez de compartilhamento de memória. Isso é feito usando goroutines e channels.

### Goroutines

> Goroutines são funções que são executadas concorrentemente com outras funções. Use a palavra-chave `go` para iniciar uma goroutine. </br>

>No Go, quando você coloca go na frente de uma função, você está dizendo: Executa essa função ao mesmo tempo que o resto do programa, sem bloquear a execução principal. </br>

> É como se você mandasse o Go: "Vai lá, executa essa função enquanto eu continuo aqui".

```go
func sayHello() {
    fmt.Println("Hello, World!")
}

func main() {
    go sayHello() // Inicia uma goroutine
    fmt.Println("Main function")
}
```

### Channels

> Channels são usados para comunicação entre goroutines. Eles permitem que você envie e receba valores entre goroutines.
- Uma goroutine é uma função que roda em paralelo com o resto do programa.
- Um canal, é um túnel de comunicação entre goroutines.
```go
func main() {
    ch := make(chan string)          // Canal que carrega strings

    go func() {
        ch <- "Hello from goroutine" // Envia uma string para o canal
    }()

    msg := <-ch
    fmt.Println(msg)                 // Saída: Hello from goroutine
}
```

### Channels de Channels

> Channels podem ser usados para enviar outros channels, permitindo a construção de pipelines complexos.

```go
func main() {
    ch1 := make(chan string)         // Canal que carrega strings
    ch2 := make(chan chan string)    // Canal que carrega canais que carregam strings

    go func() {
        ch1 <- "Hello"               // Envia uma string para o canal
    }()

    go func() {
        ch2 <- ch1                  // Envia o canal ch1 para o canal ch2
    }()

    ch := <-ch2                    // Recebe o canal ch2
    fmt.Println(<-ch)              // Saída: Hello
}
```

### Paralelização

> Go permite a execução paralela de goroutines em múltiplos núcleos, maximizando o uso de recursos de hardware.

```go
import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    runtime.GOMAXPROCS(4)          // Define o número de núcleos a serem usados

    var wg sync.WaitGroup          // Grupo de espera
    wg.Add(2)                      // Adiciona 2 goroutines ao grupo de espera

    go func() {
        defer wg.Done()            // Indica que a goroutine terminou
        fmt.Println("Goroutine 1") // Imprime a mensagem
    }()

    go func() {
        defer wg.Done()            // Indica que a goroutine terminou
        fmt.Println("Goroutine 2") // Imprime a mensagem
    }()

    wg.Wait()                      // Aguarda todas as goroutines terminarem

    fmt.Println("Todas as goroutines terminaram")

    // Saída:
    // Goroutine 1
    // Goroutine 2
    // Todas as goroutines terminaram
}
```

### Um Buffer Vazante

> Um buffer vazante é um padrão onde um buffer é usado para suavizar picos de carga, permitindo que o sistema continue a processar dados mesmo quando a entrada é intermitente.

 **Qual é a prática comum disso?**

- Útil quando os dados chegam em rajadas (picos) mas o consumidor lê num ritmo mais lento.
- Exemplo prático:
    - Leituras de sensores a cada segundo
    - Mas o processamento/armazenamento pode levar 2s
    - O buffer impede que você perca leituras

**Quando usar?**
- Quando a produção de dados é mais rápida que o consumo.
- Quando você quer evitar bloqueios ou perda de dados temporariamente.
- Para balancear a carga em sistemas concorrentes.

```go
func main() {
    // Canal com buffer de tamanho 3
    ch := make(chan int, 3)

    go func() {
        // Vai mandando 0 até 4 para o canal
        for i := 0; i < 5; i++ {
            ch <- i
            // Vai imprimindo o valor enviado
            fmt.Println("Enviado:", i)
        }
        // Fecha o canal, quando terminou de enviar tudo
        close(ch)
    }()

    // Lê o canal até que ele seja fechado
    for val := range ch {
        // Cada valor recebido é impresso como "Recebido: valor"
        fmt.Println("Recebido:", val)
    }
}

// O que acontece aqui?
// 1. Cria um canal com buffer de tamanho 3 que não precisa ser  lido imediatamente
// 2. Inicia uma goroutine que empurra os valores para o canal
// 3. Quando o buffer encher (3 valores), ela para e espera alguém ler (senão trava)
// 4. A main vai consumindo os valores, liberando espaço no buffer
// 5. Isso cria um ritmo equilibrado entre a produção e o consumo

// Saída provável:
// Enviado: 0
// Enviado: 1
// Enviado: 2
// Recebido: 0
// Recebido: 3
// Recebido: 1
// Recebido: 4
// Recebido: 2
// Recebido: 3
// Recebido: 4

```

## 🚨 Erros

### Panic

> `panic` é uma função embutida que interrompe a execução normal de um programa. Quando uma função chama `panic`, a execução do programa para imediatamente e começa a desenrolar a pilha de chamadas, executando qualquer função `defer` ao longo do caminho.

```go
func main() {
    fmt.Println("Início do programa")
    panic("Algo deu errado!") // Interrompe a execução
    fmt.Println("Isso não será impresso")
}
```

### Recover

> `recover` é uma função embutida que permite que um programa lide com uma situação de pânico. `recover` só funciona dentro de uma função `defer`. Se for chamado fora de uma função `defer`, ele não faz nada e retorna `nil`.

> **Dica**: Use `panic` e `recover` com moderação. Eles são mais adequados para situações em que o programa não pode continuar a execução normal, como erros de programação ou falhas inesperadas. Para erros comuns, prefira retornar um valor de erro.

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado do pânico:", r)
        }
    }()

    fmt.Println("Início do programa")
    panic("Algo deu errado!") // Interrompe a execução
    fmt.Println("Isso não será impresso")
}
```

## 🌐 Servidor Web

### Criando um Servidor Web Simples

> Go facilita a criação de servidores web com o pacote `net/http`. Você pode definir manipuladores de rota e iniciar um servidor em poucas linhas de código.

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloHandler)
    fmt.Println("Servidor iniciado na porta 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
```

### Manipuladores de Rota

> Manipuladores de rota são funções que lidam com solicitações HTTP para uma rota específica. Use `http.HandleFunc` para associar uma função a uma rota.

```go
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Sobre nós")
}

func main() {
    http.HandleFunc("/about", aboutHandler)
    http.ListenAndServe(":8080", nil)
}
```

### Servindo Arquivos Estáticos

> Use `http.FileServer` para servir arquivos estáticos, como HTML, CSS e JavaScript.

```go
func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    fmt.Println("Servidor iniciado na porta 8080")
    http.ListenAndServe(":8080", nil)
}
```

### Middleware

> Middleware são funções que envolvem manipuladores de rota para adicionar funcionalidades, como autenticação ou registro de logs.

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("Solicitação recebida: %s %s\n", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func main() {
    helloHandler := http.HandlerFunc(helloHandler)
    http.Handle("/", loggingMiddleware(helloHandler))
    http.ListenAndServe(":8080", nil)
}
```

Essas explicações fornecem uma visão clara sobre como criar e gerenciar um servidor web em Go, destacando a criação de servidores, manipuladores de rota, arquivos estáticos e middleware. Se precisar de mais detalhes ou ajustes, sinta-se à vontade para pedir!



