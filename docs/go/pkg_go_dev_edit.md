b# 📘 Introdução ao Go - Estilo e Idiomas Idiomáticos	

Escrito em: 23/04/2025

Go é uma linguagem moderna com foco em simplicidade, eficiência e clareza.

---

## 💡 Conceito
> Para escrever Go bem, é importante seguir seu estilo idiomático e convenções.

Go não é Java nem C++. </br>
Evite traduções literais. Escreva "em Go".

---

## 📚 Exemplos Oficiais

Use os códigos-fonte da biblioteca padrão como referência:

- https://pkg.go.dev
- Busque por "Exemplo" na documentação de pacotes.
- Os exemplos citados aqui são direto ao ponto, sem enrolação.

---

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
// ✅ Correto: Comentários alinhados automaticamente
// comando: gofmt -w seu_arquivo.go

type T struct {
    name    string // name of the object
    value   int    // its value
}
```

```go
// ❌ Errado: desalinhado, não padronizado

type T struct {
    name string // name of the object
    value int // its value
}
```

🖨️ Saída esperada: alinhamento igual para todos os campos

---

## 🔧 Regras-chave de Formatação

### Indentação
```go
// ✅ Use TAB (não espaços)
if cond {
	fmt.Println("ok")
}
```

### Parênteses
```go
// ✅ Sem parênteses extras em if/for/switch
for i := 0; i < 10; i++ {
	fmt.Println(i)
}

// ❌ Errado:
// for (i := 0; i < 10; i++) { ... }
```

### Linha longa
```go
// ✅ Quebre linhas longas com identação extra
if longExpression ||
	anotherLongExpression {
	fmt.Println("OK")
}
```

---

## 🧠 Mapa Mental (Conceitual)

```text
Go Idiomático
|
|-- Nomenclatura clara e curta
|-- Nada de getters/setters artificiais
|-- Evite repetições de tipo (ex: "Car.CarName")
|-- Use interfaces pequenas
|-- Prefira composição à herança
```

---

## 📎 Dica final

Instale o `gofmt` ou configure seu editor para aplicar a formatação automaticamente ao salvar.

```sh
go fmt ./...
```

---

> Contin> Contin\u00ue com os próximos tópicos: nomes, pacotes e controle de fluxo ❯

