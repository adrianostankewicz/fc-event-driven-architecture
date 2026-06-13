# fc-event-driven-architecture

Projeto exemplo em Go demonstrando uma arquitetura orientada a eventos (event-driven) para um sistema financeiro simples (contas, clientes e transações).

> Nota: README escrito em Português (pt-BR).

## Visão Geral

Este repositório contém um exemplo de aplicação em Go que demonstra conceitos de arquitetura orientada a eventos: separação de camadas (entity, usecase, gateway), publicação/consumo de eventos e persistência (SQLite/MySQL via drivers configuráveis).

Principais responsabilidades:
- Gerenciar clientes e contas.
- Registrar e executar transações entre contas.
- Publicar eventos relacionados a mudanças (ex.: conta criada, transação efetuada).

## Estrutura do repositório

- /internal - código da aplicação (cmd, database, entity, event, gateway, usecase, web)
- /pkg - pacotes reutilizáveis (ex.: eventos, unit-of-work)
- /api - exemplos de requisições HTTP (arquivo client.http)
- docker-compose.yaml - config mínima para rodar um banco MySQL local
- go.mod / go.sum - dependências do projeto

## Requisitos

- Go 1.24+
- Docker e docker-compose (opcional, recomendado para iniciar banco de dados)

## Como rodar (desenvolvimento)

1. Inicie o banco de dados (MySQL) com Docker Compose:

   docker-compose up -d

   Isso cria um serviço MySQL (porta 3306) com banco `wallet` e usuário root. Ajuste variáveis de ambiente conforme necessário.

2. Instale dependências e execute a aplicação:

   - No modo rápido (executar todos os pacotes para desenvolvimento):

     go run ./...

   - Ou compile e execute o binário do serviço (por exemplo):

     go build -o bin/app ./...
     ./bin/app

Observação: dependendo da organização dos comandos (p. ex. se houver um pacote em internal/cmd específico para iniciar o servidor), substitua `./...` pelo pacote de entrada correto (ex.: `go run ./internal/cmd`).

## Endpoints de exemplo

Veja `api/client.http` para exemplos de requisições HTTP:

- Criar cliente: POST http://localhost:3000/clients
- Criar conta: POST http://localhost:3000/accounts
- Criar transação: POST http://localhost:8080/transactions

Exemplos estão incluídos no arquivo `api/client.http` no repositório.

## Testes

Execute os testes com:

    go test ./...

O projeto depende de `github.com/stretchr/testify` para assertions em testes.

## Dependências e versão do Go

- go 1.24.5 (definida em go.mod)
- Drivers e libs importantes: go-chi, mysql driver, sqlite3, google/uuid, testify

## Contribuição

Contribuições são bem-vindas. Abra issues para reportar bugs ou discutir melhorias. Para pull requests:

1. Crie uma branch com um nome descritivo.
2. Escreva testes para novas funcionalidades/bugs corrigidos.
3. Abra um Pull Request descrevendo as mudanças.

## License

Licença não especificada — adicione um arquivo LICENSE se desejar publicar este projeto sob uma licença permissiva.

---

Se quiser, posso adicionar uma seção de execução mais detalhada (com exemplos de variáveis de ambiente ou comandos específicos para os pacotes `internal/cmd` ou `internal/web`) — informe qual binário/entrada você prefere como ponto de partida.
