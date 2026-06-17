# Go Gin API Boilerplate

A production-ready REST API boilerplate built with Go and Gin, following Clean Architecture principles and best practices for maintainability, scalability, and developer experience.

This boilerplate provides a solid foundation for building backend services such as inventory systems, warehouse management systems, e-commerce platforms, SaaS applications, and internal business tools.

---

## Features

### Core

* Go (Golang)
* Gin Web Framework
* RESTful API Architecture
* Clean Architecture Structure
* Dependency Injection Pattern
* Environment-Based Configuration

### Database

* PostgreSQL
* GORM ORM
* Repository Pattern
* Database Migration Support

### Authentication & Security

* JWT Authentication
* Authorization Middleware
* Protected Routes
* Password Hashing (bcrypt)
* Environment Secret Management

### API Development

* Request Validation
* Standardized API Responses
* Error Handling
* Route Grouping
* Middleware Support

### Developer Experience

* Docker Support
* Docker Compose
* Hot Reload Support (Air)
* Structured Logging
* Makefile Commands
* Environment Templates

### Testing & Quality

* Unit Testing Ready
* Integration Testing Ready
* Linting Support
* Clean Project Structure

---

## Project Structure

```text
.
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── config/
│   │   └── config.go
│   │
│   ├── database/
│   │   └── postgres.go
│   │
│   ├── handlers/
│   │   ├── auth_handler.go
│   │   └── user_handler.go
│   │
│   ├── middleware/
│   │   ├── auth.go
│   │   └── logger.go
│   │
│   ├── models/
│   │   ├── user.go
│   │   └── role.go
│   │
│   ├── repositories/
│   │   └── user_repository.go
│   │
│   ├── services/
│   │   └── user_service.go
│   │
│   └── routes/
│       └── routes.go
│
├── pkg/
│   ├── response/
│   │   └── response.go
│   │
│   └── utils/
│       └── helpers.go
│
├── migrations/
│
├── .env.example
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

---

## Architecture Overview

```text
Client
   │
   ▼
Gin Router
   │
   ▼
Handler Layer
   │
   ▼
Service Layer
   │
   ▼
Repository Layer
   │
   ▼
PostgreSQL Database
```

### Layer Responsibilities

| Layer      | Responsibility                      |
| ---------- | ----------------------------------- |
| Handler    | HTTP request and response handling  |
| Service    | Business logic                      |
| Repository | Database operations                 |
| Model      | Domain entities                     |
| Middleware | Authentication, logging, validation |
| Config     | Environment configuration           |

---

## Prerequisites

Before running the application, ensure the following are installed:

* Go 1.24+
* PostgreSQL 15+
* Docker (Optional)
* Docker Compose (Optional)

---

## Getting Started

### 1. Clone Repository

```bash
git clone https://github.com/indahcs/go-gin-api-boilerplate.git

cd go-gin-api-boilerplate
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Configure Environment Variables

Create a local environment file:

```bash
cp .env.example .env
```

Example:

```env
APP_ENV=development

PORT=8080

DATABASE_URL=postgres://postgres:password@localhost:5432/app_db

JWT_SECRET=your-super-secret-key

JWT_EXPIRE_HOURS=24
```

### 4. Start PostgreSQL

Using Docker:

```bash
docker-compose up -d
```

### 5. Run Application

```bash
go run ./cmd/server
```

Server will be available at:

```text
http://localhost:8080
```

---

## Development Commands

### Run Application

```bash
make run
```

### Build Application

```bash
make build
```

### Run Tests

```bash
make test
```

### Run Linter

```bash
make lint
```

### Start Docker Services

```bash
docker-compose up -d
```

### Stop Docker Services

```bash
docker-compose down
```

---

## Environment Variables

| Variable         | Description                  |
| ---------------- | ---------------------------- |
| APP_ENV          | Application environment      |
| PORT             | API server port              |
| DATABASE_URL     | PostgreSQL connection string |
| JWT_SECRET       | JWT signing secret           |
| JWT_EXPIRE_HOURS | Token expiration time        |

---

## Example API Endpoints

### Health Check

```http
GET /api/v1/health
```

Response:

```json
{
  "success": true,
  "message": "Server is running"
}
```

### Authentication

```http
POST /api/v1/auth/register
POST /api/v1/auth/login
```

### Users

```http
GET    /api/v1/users
GET    /api/v1/users/:id
POST   /api/v1/users
PUT    /api/v1/users/:id
DELETE /api/v1/users/:id
```

---

## Creating a New Project from This Boilerplate

### Option 1: Use GitHub Template

1. Click **Use this template**
2. Create a new repository
3. Clone your new repository

### Option 2: Clone and Rename

```bash
git clone https://github.com/indahputri/go-gin-api-boilerplate.git my-project

cd my-project
```

Update the Go module:

```bash
go mod edit -module github.com/your-username/my-project

go mod tidy
```

Update:

* Project name
* Environment variables
* Database schema
* Models
* Routes
* Business logic

---

## Recommended Enhancements

Depending on project requirements, consider adding:

* Refresh Tokens
* Role-Based Access Control (RBAC)
* Swagger / OpenAPI Documentation
* Redis Caching
* Email Service Integration
* Background Jobs
* File Upload Support
* Rate Limiting
* Audit Logging
* CI/CD Pipelines
* Kubernetes Deployment

---

## Contributing

Contributions, issues, and feature requests are welcome.

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push your branch
5. Open a Pull Request

---