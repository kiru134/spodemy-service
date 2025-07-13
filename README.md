# Spodemy Service

A backend service for managing multi-sports academy venues and batches. Built with Go and Gin framework.

## Requirements

- Go 1.24 or later
- PostgreSQL 14 or later
- Git

## Features

- Venue Management (CRUD operations)
- Batch Management
- User Authentication
- Role-based Access Control
- RESTful API with Swagger Documentation
- PostgreSQL Database

## Installation

1. Clone the repository:

```bash
git clone https://github.com/kiru134/spodemy-service.git
cd spodemy-service
```

2. Install dependencies:

```bash
go mod download
```

3. Create a `config/local.json` file:

```json
{
  "database": {
    "host": "localhost",
    "port": 5432,
    "user": "your_username",
    "password": "your_password",
    "name": "spodemy_db"
  },
  "server": {
    "port": 8080
  }
}
```

## Database Setup

1. Create PostgreSQL database:

```sql
CREATE DATABASE spodemy_db;
```

2. Enable UUID extension:

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

3. Run migrations:

```bash
go run migrate/migrate.go
```

## Running the Application

1. Start the server:

```bash
go run main.go
```

2. The server will start at `http://localhost:8080`

## API Documentation

1. Generate Swagger documentation:

```bash
swag init
```

2. Access Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

## Project Structure

```
spodemy-backend/
├── config/             # Configuration files
├── controllers/        # Request handlers
├── database/          # Database connection and migrations
├── docs/              # Generated Swagger documentation
├── models/            # Data models
├── repositories/      # Database operations
├── routes/            # Route definitions
├── services/          # Business logic
└── main.go           # Application entry point
```

## API Endpoints

### Venues

- `GET /api/v1/venues` - List all venues
- `GET /api/v1/venues/{id}` - Get venue by ID
- `POST /api/v1/venues` - Create new venue
- `PUT /api/v1/venues/{id}` - Update venue
- `DELETE /api/v1/venues/{id}` - Delete venue

### Batches

- Similar CRUD operations for batches

## Development

1. Install Swagger tools:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Generate Swagger docs after API changes:

```bash
swag init
```

3. Run tests:

```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
