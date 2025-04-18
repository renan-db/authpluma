Significado das camadas

internal/modules/user/
├── dto/           → Modelos de entrada/saída da API (Request/Response)
├── entity/        → Entidades do domínio (ex: User struct com regras de negócio)
├── usecase/       → Casos de uso (ex: CreateUser, LoginUser, etc)
├── handler/       → Controladores HTTP (ex: métodos que lidam com Echo ou Gin)
├── route/         → Definição das rotas e agrupamentos
├── repository/    → Implementações de acesso a dados (ex: PostgreSQL, Redis, Memória Local)
├── interface/     → Interfaces de abstração (ex: user_interface, Hasher)