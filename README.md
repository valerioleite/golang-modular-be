# Golang Modular Backend

A modular Go backend built with **Clean Architecture** principles, using Go workspaces to organize independent services and shared libraries into a single deployable monolith.

## Tech Stack

| Component        | Technology                                  |
|------------------|---------------------------------------------|
| Language         | Go 1.25                                     |
| Web Framework    | Standard library `net/http`                 |
| Authentication   | OpenID Connect (Auth0, Keycloak, Okta)      |
| Database         | PostgreSQL 18                               |
| Object Storage   | AWS S3 (LocalStack for local dev)           |
| API Docs         | Swagger / OpenAPI (swaggo)                  |
| Container        | Docker (distroless multi-stage)             |
| CI/CD            | GitHub Actions                              |

## Project Structure

```
.
├── libraries/
│   ├── db/          # PostgreSQL connection & migration system
│   ├── domain/      # Shared domain error types
│   └── http/        # HTTP server, middleware, CORS, health checks, Swagger
├── services/
│   ├── monolith/    # Entry point — bootstraps and composes all services
│   ├── authentication/  # OIDC login, callback, token refresh
│   ├── user/        # User creation & retrieval
│   ├── tenant/      # Multi-tenant settings & branding
│   └── storage/     # AWS S3 file upload/download
├── go.work          # Go workspace definition
├── docker-compose.yaml
└── .github/workflows/
```

Each service follows Clean Architecture layers:

```
domain/ → repository/ → service/ → delivery/http/
```

## Getting Started

### Prerequisites

- Go 1.25+
- Docker & Docker Compose
- An OIDC provider (Auth0, Keycloak, Okta, etc.)

### Setup

1. **Clone the repository**

```bash
git clone git@github.com:valerioleite/golang-modular-be.git
cd golang-modular-be
```

2. **Configure environment variables**

```bash
cp .env.example .env
```

Edit `.env` with your values:

```env
# HTTP
PORT=8080
LOGGING_REQUESTS=true

# OIDC
OIDC_ISSUER_URL=https://your-provider.com/
OIDC_CLIENT_ID=your_client_id
OIDC_CLIENT_SECRET=your_client_secret
OIDC_REDIRECT_URI=http://localhost:8080/api/authentication/callback

# Database
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_DB=tenant-database
DATABASE_SSL_MODE=false

# AWS S3 (LocalStack for local dev)
AWS_REGION=us-east-1
AWS_ENDPOINT=http://localhost:4566
AWS_CREDENTIALS_KEY=test
AWS_CREDENTIALS_SECRET=test

# Internal services
STORAGE_BASE_URL=http://localhost:8080/api/storage/
```

3. **Start infrastructure**

```bash
docker compose up -d
```

This starts PostgreSQL on port `5432`.

4. **Run the application**

```bash
cd services/monolith
go run ./cmd
```

The server starts on `http://localhost:8080`. Database migrations run automatically on startup.

## API Endpoints

### Authentication

| Method | Endpoint                                              | Description         |
|--------|-------------------------------------------------------|---------------------|
| GET    | `/api/authentication/v1/authentication/authorize`     | Initiate OIDC login |
| GET    | `/api/authentication/v1/authentication/callback`      | OIDC callback       |
| POST   | `/api/authentication/v1/authentication/refresh`       | Refresh token       |
| GET    | `/api/authentication/v1/authentication/userinfo`      | Get user info       |

### Users

| Method | Endpoint                          | Description          |
|--------|-----------------------------------|----------------------|
| POST   | `/api/user/v1/users`              | Create user          |
| GET    | `/api/user/v1/users/sub/{sub}`    | Get user by OIDC sub |

### Tenants

| Method | Endpoint                                | Description          |
|--------|-----------------------------------------|----------------------|
| POST   | `/api/tenant/v1/tenants`                | Create tenant        |
| GET    | `/api/tenant/v1/tenants`                | List tenants         |
| GET    | `/api/tenant/v1/tenants/{id}`           | Get tenant by ID     |
| PUT    | `/api/tenant/v1/tenants/{id}`           | Update tenant        |
| PUT    | `/api/tenant/v1/tenants/{id}/image`     | Upload tenant image  |
| DELETE | `/api/tenant/v1/tenants/{id}`           | Delete tenant        |

### Storage

| Method | Endpoint                                      | Description    |
|--------|-----------------------------------------------|----------------|
| POST   | `/api/storage/v1/files`                       | Upload file    |
| GET    | `/api/storage/v1/files/{bucket}/{filename}`   | Download file  |

### Health Checks

Each service exposes `GET /api/{service}/v1/actuator/health`.

## Docker

Build and run the monolith container:

```bash
docker build -f services/monolith/Dockerfile -t golang-modular-be .
docker run --env-file .env -p 8080:8080 golang-modular-be
```

The image uses a multi-stage build with `distroless` for a minimal, secure runtime.

## License

This project is for learning and experimentation purposes.
