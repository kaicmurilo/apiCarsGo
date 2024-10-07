# API Rest Go - Carros

Uma API REST simples em Go para gerenciar uma lista de carros, utilizando o framework [Echo](https://echo.labstack.com/) e o banco de dados SQLite.

## Requisitos

- [Go](https://golang.org/) 1.19 ou superior
- [SQLite3](https://www.sqlite.org/download.html)
- Biblioteca Echo: `github.com/labstack/echo/v4`
- Driver SQLite para Go: `github.com/mattn/go-sqlite3`

## Instalação

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/seu-usuario/nome-do-repositorio.git
   cd nome-do-repositorio
   ```

2. **Instale as dependências:**

   Execute o seguinte comando no diretório do projeto para baixar as dependências:

   ```bash
   go mod tidy
   ```

3. **Crie o banco de dados SQLite:**

   ```bash
   touch cars.db
   sqlite3 cars.db "CREATE TABLE cars (name TEXT, price NUMBER);"
   ```

## Executando a Aplicação

1. Inicie o servidor com o seguinte comando:

   ```bash
   go run main.go
   ```

2. A API estará disponível em `http://localhost:8080`.

## Endpoints da API

### GET /cars

Retorna uma lista de todos os carros no banco de dados.

- **Resposta de Sucesso:** `200 OK`
- **Exemplo de Resposta:**

  ```json
  [
    {
      "name": "Toyota",
      "price": 10000
    },
    {
      "name": "Honda",
      "price": 15000
    }
  ]
  ```

### POST /cars

Adiciona um novo carro ao banco de dados.

- **Requisição:**

  - **Body:** JSON com o nome e o preço do carro
  - **Exemplo:**

    ```json
    {
      "name": "Ford",
      "price": 20000
    }
    ```

- **Resposta de Sucesso:** `201 Created`
- **Exemplo de Resposta:**

  ```json
  {
    "name": "Ford",
    "price": 20000
  }
  ```

## Estrutura do Código

- **`Car struct`**: Representa o modelo de um carro com `Name` e `Price`.
- **`getCars`**: Handler para o endpoint GET `/cars`, que retorna a lista de carros em formato JSON.
- **`createCar`**: Handler para o endpoint POST `/cars`, que cria um novo carro no banco de dados.
- **`saveCar`**: Função auxiliar que insere um carro no banco de dados SQLite.
- **`main`**: Configura e inicia o servidor, definindo as rotas da API.

## Dependências

- Echo Framework: `github.com/labstack/echo/v4`
- SQLite Driver: `github.com/mattn/go-sqlite3`
