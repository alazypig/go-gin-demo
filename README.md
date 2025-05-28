# Go Gin Video API

A RESTful API built with Go and Gin framework for managing video resources with JWT authentication.

## Features

- CRUD operations for video resources
- JWT-based authentication
- Swagger documentation
- HTML template rendering
- SQLite database with GORM
- Input validation
- Structured logging with Zap
- Basic auth middleware

## Prerequisites

- Go 1.24.0 or higher
- SQLite

## Project Structure

```
go-gin/
├── api/            # API handlers
├── controller/     # Business logic controllers
├── dto/           # Data transfer objects
├── entity/        # Database models
├── middlewares/   # Custom middleware functions
├── repository/    # Database operations
├── service/       # Business logic services
├── templates/     # HTML templates
├── validators/    # Custom validators
└── docs/         # Swagger documentation
```

## Installation

1. Clone the repository

```bash
git clone <repository-url>
```

2. Install dependencies

```bash
go mod download
```

3. Generate Swagger documentation

```bash
./swagger.sh
```

4. Build the application

```bash
./build.sh
```

## Running the Application

```bash
./bin/application
```

The server will start on port 8080 by default. You can change the port by setting the `PORT` environment variable.

## API Documentation

Access the Swagger documentation at:

```
http://localhost:8080/swagger/index.html
```

## Authentication

The API uses JWT tokens for authentication. To get a token:

1. Make a POST request to `/api/v1/auth/token`
2. Default credentials:
   - Username: `admin`
   - Password: `123456`

## Environment Variables

- `PORT`: Server port (default: 8080)
- `JWT_SECRET`: Secret key for JWT tokens (default: "edward-secret-key")

## License

[MIT License](LICENSE)
