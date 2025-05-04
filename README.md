"# authpluma" 
Autenticador centralizado desenvolvido em Go com foco em escalabilidade e segurança.

## 🎯 Funcionalidades

- Autenticação de usuários
- Autorização de acesso a recursos
- Gerenciamento de sessões

## 📚 Documentação

O projeto adota boas práticas de desenvolvimento, como os princípios SOLID e a Clean Architecture, entre outras. 
Além disso, a estrutura foi baseada no repositório golang-api-template, cuja documentação está disponível na camada /docs.

## 🚀 Execução Local

O projeto foi desenvolvido em ambiente Linux, utilizando Ubuntu 24.04.1 LTS via WSL2. Recomenda-se utilizar o mesmo ambiente para garantir compatibilidade na execução local.

- Instale o [Make](https://www.gnu.org/software/make/manual/make.html)
- Clone o repositório e entre na pasta do projeto
- Crie um arquivo `.env` na raiz baseado no `.env.example` (se existir).
- Execute o comando `make setup-linux` para instalar as dependências e configurar o ambiente.
- Execute o comando `make dev` para iniciar os containers.

## Comandos Essenciais

- `make`        - Exibe a lista de comandos disponíveis.
