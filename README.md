# Api-Postgres-Go

Project for learning golang by applying it to create an API with chi

# Table of Contents

- [Api-Postgres-Go](#api-postgres-go)
- [Table of Contents](#table-of-contents)
- [Technologies](#technologies)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Without Docker](#without-docker)
  - [With Docker](#with-docker)
- [Plus](#plus)
- [References](#references)

# Technologies

- [Chi](https://go-chi.io/)
- [Golang](https://golang.org/)
- [Http-swagger](github.com/swaggo/http-swagger)
- [Migrate](github.com/golang-migrate/migrate/)
- [Postgres](https://www.postgresql.org/)
- [PQ](github.com/lib/pq)
- [Swag](github.com/swaggo/swag)
- [Validator](https://github.com/go-playground/validator)
- [Viper](github.com/spf13/viper)

# Project Structure

```bash
├── configs
│   └── config.go
├── db
│   └── migrations
│   └── connection.go
├── docs
│   └── docs.go
│   └── swagger.json
│   └── swagger.yaml
├── handlers
│   └── create.go
│   └── delete.go
│   └── list.go
│   └── retrieve.go
│   └── update.go
├── models
│   └── delete.go
│   └── entities.go
│   └── get_all.go
│   └── get.go
│   └── insert.go
│   └── update.go
├── .pre-commit-config.yaml
├── .commitlint.config.js
├── config.toml
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── README.md
```

# Getting Started

- Clone the repository

```bash
git clone https://github.com/SousaPedro11/api-postgres-go.git
```

## Without Docker

- Install dependencies

```bash
go mod download
```

- Run the project

```bash
go run main.go
```

## With Docker

- Run the project

```bash
docker-compose up
```

_After run You can access the API at http://localhost:9000 and the swagger at http://localhost:9000/swagger/index.html. The api port and the database port can be changed in the config.toml file._

# Plus

I added some tools to help me with the project:

- pre-commit to format the code and the commitlint to validate the commit messages, lint the code and run the tests.
- docker-compose to run the project with postgres database.
- swagger to document the API.
- viper to read the config file.
- validator to validate the request body.
- golang-migrate to run the migrations.

# References

- [Como implementar uma API do ZERO com Golang e PostgreSQL!!](https://www.youtube.com/watch?v=MD7b-iQMC24&ab_channel=AprendaGolang)
