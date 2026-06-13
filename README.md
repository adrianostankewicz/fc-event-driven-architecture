# Event-Driven Architecture

Um projeto de arquitetura orientada a eventos desenvolvido em **Go**, demonstrando padrões e práticas para construir sistemas escaláveis e desacoplados.

## Visão Geral

Este projeto implementa uma arquitetura baseada em eventos, permitindo que componentes se comuniquem de forma assíncrona e desacoplada.

## Tecnologias

- **Linguagem**: Go 1.24.5
- **HTTP Framework**: Chi (go-chi)
- **Banco de Dados**: MySQL 5.7
- **UUID**: Google UUID
- **Containerização**: Docker & Docker Compose
- **Testes**: Testify

## Estrutura do Projeto

```
fc-event-driven-architecture/
├── api/              # Endpoints e handlers HTTP
├── internal/         # Código interno da aplicação
├── pkg/             # Pacotes reutilizáveis
├── docker-compose.yaml
├── go.mod
└── README.md
```

## Como executar

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

## Arquitetura

A arquitetura segue o padrão **Event-Driven**, onde:

- **Eventos** são os blocos construtivos fundamentais
- **Produtores** geram eventos ao executar ações
- **Consumidores** reagem aos eventos de forma desacoplada
- **Banco de Dados** mantém o estado persistido

### Camadas

- **API**: Interface HTTP para interação externa
- **Internal**: Lógica de negócio, handlers de eventos, repositórios
- **Pkg**: Componentes reutilizáveis (utils, modelos, etc.)
