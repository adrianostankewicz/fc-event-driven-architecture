# Event-Driven Architecture

Um projeto de arquitetura orientada a eventos desenvolvido em **Go**, demonstrando padrões e práticas para construir sistemas escaláveis e desacoplados.

## 📋 Visão Geral

Este projeto implementa uma arquitetura baseada em eventos, permitindo que componentes se comuniquem de forma assíncrona e desacoplada. É uma excelente base para aprender e explorar padrões de microsserviços e sistemas distribuídos.

## 🛠️ Tecnologias

- **Linguagem**: Go 1.24.5
- **HTTP Framework**: Chi (go-chi)
- **Banco de Dados**: MySQL 5.7
- **UUID**: Google UUID
- **Containerização**: Docker & Docker Compose
- **Testes**: Testify

## 📦 Estrutura do Projeto

```
fc-event-driven-architecture/
├── api/              # Endpoints e handlers HTTP
├── internal/         # Código interno da aplicação
├── pkg/             # Pacotes reutilizáveis
├── docker-compose.yaml
├── go.mod
└── README.md
```

## 🚀 Começando

### Pré-requisitos

- Go 1.24.5 ou superior
- Docker e Docker Compose
- MySQL (opcional, pode usar via Docker)

### Instalação

1. **Clone o repositório**

```bash
git clone https://github.com/adrianostankewicz/fc-event-driven-architecture.git
cd fc-event-driven-architecture
```

2. **Instale as dependências**

```bash
go mod download
```

3. **Inicie os serviços com Docker Compose**

```bash
docker-compose up -d
```

Isso iniciará uma instância MySQL com o banco de dados `wallet` e senha `root`.

### Desenvolvimento

4. **Execute a aplicação**

```bash
go run ./cmd/main.go
```

Ou compile:

```bash
go build -o app ./cmd/main.go
./app
```

## 🧪 Testes

Execute todos os testes com:

```bash
go test ./...
```

Com cobertura:

```bash
go test -cover ./...
```

## 🏗️ Arquitetura

A arquitetura segue o padrão **Event-Driven**, onde:

- **Eventos** são os blocos construtivos fundamentais
- **Produtores** geram eventos ao executar ações
- **Consumidores** reagem aos eventos de forma desacoplada
- **Banco de Dados** mantém o estado persistido

### Camadas

- **API**: Interface HTTP para interação externa
- **Internal**: Lógica de negócio, handlers de eventos, repositórios
- **Pkg**: Componentes reutilizáveis (utils, modelos, etc.)

## 📊 Banco de Dados

O projeto usa MySQL 5.7 com os seguintes serviços:

- Host: `localhost`
- Porta: `3306`
- Root Password: `root`
- Database: `wallet`
- User: `root`

## 📝 Dependências Principais

| Pacote | Versão | Descrição |
|--------|--------|-----------|
| chi | v5.2.3 | Router HTTP |
| google/uuid | v1.6.0 | Geração de UUIDs |
| go-sql-driver/mysql | v1.9.3 | Driver MySQL |
| mattn/go-sqlite3 | v1.14.32 | Driver SQLite |
| testify | v1.11.1 | Framework de testes |

## 📚 Recursos

- [Go Documentation](https://golang.org/doc)
- [Chi Router](https://github.com/go-chi/chi)
- [Event-Driven Architecture Pattern](https://www.ibm.com/cloud/architecture/architectures/event-driven)

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para:

1. Fazer um Fork
2. Criar uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abrir um Pull Request

## 📄 Licença

Este projeto está disponível sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

## 👤 Autor

**Adriano Stankewicz**

- GitHub: [@adrianostankewicz](https://github.com/adrianostankewicz)

---

**Última atualização**: Junho 2026
