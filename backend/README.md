

# Lema AI Users-Posts Backend

A robust backend service for managing users and posts, built with Go. This project uses GORM as the ORM and SQLite as the database.

## Prerequisites

- Go 1.23.0 or higher
- SQLite3
- Git

## Quick Start

1. Clone the repository:
   ```bash
   git clone https://github.com/TimiBolu/lema-ai-users-service.git
   cd lema-ai-users-service
   ```

2. Run the installation script:
   ```bash
   chmod +x bin/install.sh
   ./bin/install.sh
   ```

3. Start the server:
   ```bash
   ./server
   ```

The server will start on `localhost:8080` (or your configured port)

## API Documentation

Visit `/docs` route after starting the server to access the complete API documentation.

## Development

For development with hot-reload:
```bash
go install github.com/gravityblast/fresh@latest
fresh
```

## Project Structure
```bash

├── README.md
├── bin
│   └── install.sh         # Installation script
├── config
│   └── config.go          # Application configuration
├── database
│   ├── database.go        # Database connection and setup
│   └── seed.go            # Data seeding utilities
├── docs
│   ├── api.md             # API documentation in markdown
│   └── index.html         # Interactive API documentation
├── handlers
│   ├── handlers_test.go   # Handler tests
│   ├── posts.go           # Post-related handlers
│   └── users.go           # User-related handlers
├── main.go                 # Application entry point
├── models
│   ├── Address.go         # Address model
│   ├── Post.go            # Post model
│   └── User.go            # User model
├── router
│   └── router.go          # Application routing
└── test.db                 # SQLite database file
```

## Testing

Run tests with:
```bash
go test ./... -v
```

## Dependencies

- gorilla/mux - HTTP router and URL matcher
- gorm.io/gorm - ORM library
- go-faker/faker - Test data generation
- joho/godotenv - Environment variable loading
- rs/cors - CORS middleware
- stretchr/testify - Testing toolkit
- And more (see go.mod for complete list)

## Environment Variables

Copy the example environment file and adjust as needed:
```bash
cp .env.example .env
```
