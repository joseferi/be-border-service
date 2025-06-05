# <div align="center">Loyalty Team</div>

<div align="center"><img src="image.png"></img></div>
<!-- ![icon logo](image.png) -->

# 🗺️ Overview
This boilerplate is designed to help developers get started quickly with native Go http with mux router. It includes basic structure, middleware setup, configuration management, and sample modules.

## ✅ Features
- Clean architecture using Domain Design Driven
- Config management using <b>Viper by spf13</b>
- Middleware
- Gracefully shutdown
- Structured logging with <b>zap</b>
- Dependency Injection
- Docker support with multi stage method build
- Health check support
- Migrations tools generate table using <b>Goose</b>
- Command line using <b>Cobra by spf13 </b>
- Background jobs process with <b>Asynq</b>
- [TODO] Dashboard background jobs
- Validation tools with go-validator v10


## 🛠️ Folder Structure
### Folder Details

- `cmd/`: Contains the main application file. Acts as the entry point for the service.
- `database/`: Migrations file .sql.
- `deployment/`: Dockerfile & docker-compose file for deployment local setup.
- `internal/`: Core logic of the application, divided into multiple modules.
  - `bootstrap/`: Contains function registry booting up package.
  - `common/`: Common function reusable, like base response, encrypting, base64, etc.
  - `config/`: Define struct of file configuration.
  - `delivery/`: Define struct of params from request incoming like payload or query params.
  - `handler/http.go`: Wrapping http method or another protocol like grpc or MQ.
    - `handler/tasks/`: handler process queue do.
  - `middleware/`: middlewares file.
  - `model/`: Representation attribute table from database.
  - `providers/`: Define thirdparty function, like sendgrid mail, payment-gateway or another service dependency.
  - `repository/`: Implementation of data persistence or external calls.
  - `routes/`: Implementation of usecase process into route path.
  - `usecase/`: Application use case logic (service layer).
  - `validators/`: Shared function with specfic case / rules, like phone number format, identity number.
- `pkg/`: Shared libraries that can be reused across the app.
- `Makefile`: CLI commands for development tasks

## 📖 How to install
1. Please use golang version manager, we recommend using ```g``` version manager <br>
https://github.com/voidint/g


2. Installing ```golangci-lint``` first for linters tool in local device. <br>
```sh
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6
```
3. Create file ```.env``` on main directory
4. Build or running service can using docker compose setup.
```sh
docker-compose -f deployment/docker-compose.yml up -d
```

## 📑 References
- https://dev.to/kittipat1413/understanding-the-options-pattern-in-go-390c
- https://github.com/hibiken/asynq/wiki/Getting-Started
- https://golangci-lint.run/
- https://github.com/spf13
- https://refactoring.guru/
- https://www.conventionalcommits.org/en/v1.0.0/