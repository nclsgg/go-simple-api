# Go Simple API

Uma simples API com métodos HTTP e conexão com banco de dados.
Esse projeto foi feito para fins de estudo da linguagem Go. 

## Tecnologias Utilizadas

- [Fiber](https://gofiber.io/)
- [mongo-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)

## Requisitos

- [Docker](https://www.docker.com/get-started)
- [Go](https://golang.org/dl/)

## Configuração do Ambiente

1. Clone o repositório:

    ```bash
    git clone https://github.com/seu-usuario/seu-repositorio.git
    cd seu-repositorio
    ```

2. Copie o arquivo `.env.example` para `.env` e edite conforme necessário:

    ```bash
    cp .env.example .env
    ```

3. Inicie o banco de dados com Docker Compose:

    ```bash
    docker-compose up -d
    ```

## Instalando Dependências

Para instalar as dependências do projeto, execute:

```bash
go mod tidy
```

## Executando o Projeto

Para rodar o projeto, execute:

```bash
go run main.go
```
