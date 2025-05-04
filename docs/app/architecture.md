# Arquitetura do Projeto

Este documento descreve a arquitetura geral do projeto, suas camadas, responsabilidades e padrões utilizados.

## Estrutura do Projeto

Somente os arquivos e camadas extremamente relevantes são citados.

```
.
├── cmd/                          # Ponto de entrada da aplicação (main.go, etc.)
├── internal/                     # Código interno da aplicação (não exposto para outros projetos)
│   └── modules/                  # Organização por módulos (ex: user, auth, etc.)
│       └── user/                 # Exemplo: módulo de usuários
│           ├── doc.go            # Documentação do módulo
│           ├── entity/           # Definições de entidades de domínio
│           ├── dto/              # Data Transfer Objects (requisição/resposta)
│           ├── interface/        # Interfaces/Contratos entre camadas
│           ├── usecase/          # Lógica de negócio (application layer)
│           ├── handler/          # Entrypoints HTTP (controllers)
│           ├── repository/       # Comunicação com banco de dados ou fontes externas
│           ├── route/            # Configuração de rotas
│           └── helper/           # Utilitários internos (ex: binder, validator)
├── docs/                         # Documentação do projeto
│   ├── app/                      # Documentação do projeto
│   │   ├── architecture.md       # Arquitetura do projeto
│   │   ├── default_comments.md   # Documentação de comentários padrão
│   └── makefiles/                # Arquivos Makefile para facilitar a execução de comandos
│   └── swagger/                  # Documentação Swagger (gerada pelo Swaggo)
├── go.mod                        # Gerenciamento de dependências
├── go.sum                        # Checksum das dependências
├── Dockerfile                    # Dockerfile para construir a imagem Docker
├── docker-compose.yml            # Docker Compose para iniciar o projeto
├── README.md                     # Descrição principal do projeto
```



## Padrões Utilizados

- **Clean Architecture** (inspiração principal)
- **DDD (Domain-Driven Design)** no nível de organização de entidades
- **SOLID Principles** para separar responsabilidades
- **Dependency Injection** manual para gerenciar dependências


## Camadas da Arquitetura

### 1. Entity

- Define as entidades centrais do domínio.
- Camada 100% independente de frameworks e tecnologias externas.
- São estruturas puras sem dependência de frameworks ou tecnologias externas.
- Exemplo: `UserEntity` com campos que representam um usuário.

### 2. DTO (Data Transfer Object)

- Representam o formato de entrada/saída da API.
- Úteis para separar a representação interna da exposição externa.
- O DTO é reponsavel por receber os dados da requisição e devolver os dados da resposta padronizados.

### 3. Interface

- Define contratos (interfaces) que descrevem comportamentos esperados.
- Exemplo: `UserUseCase`, `UserRepository`, `UserHandler`.
- Facilita a substituição de implementações (mock para testes, etc).

### 4. Usecase

- Deve ser quase 100% pura, se precosar de algo externo, deve implementar através de interfaces.
- Contém a lógica de aplicação e orquestra chamadas entre entidades e repositórios.
- Regras de negócio vivem aqui, não nos handlers.
- Exemplo: Criar usuário, atualizar, validar regras específicas.

### 5. Handler

- Responsável por receber requisições HTTP.
- Faz o bind dos dados (entrada), chama o caso de uso apropriado e devolve uma resposta formatada.
- Utiliza o Echo Framework (`echo.Context`).

### 6. Repository

- Implementa o acesso a dados (banco de dados, API externa, etc).
- Nunca expõe detalhes da tecnologia usada para quem está acima (usecases).

### 7. Helper

- Utilitários que auxiliam handlers e usecases.
- Exemplo:
  - `binder/` para extrair dados da requisição para o DTO.
  - `validate/` para validar os dados recebidos.

### 8. Route

- Define as rotas HTTP da aplicação.
- Conecta endpoints (URLs) aos handlers correspondentes.


## Principais Convenções

- Cada módulo (`user`, `auth`, etc) é isolado dentro de `internal/modules/`.
- Um módulo deve ser o mais independente possível de outros módulos.
- Funções públicas começam com letra maiúscula; funções privadas com minúscula.
- Comentários para documentação de pacotes (`doc.go`) sempre que necessário.


## Benefícios da Arquitetura

✅ Fácil manutenção  
✅ Alta testabilidade  
✅ Isolamento de responsabilidade  
✅ Flexibilidade para evoluir camadas sem impactar o restante  
✅ Independência de frameworks (Echo usado apenas no handler)


# Fim
