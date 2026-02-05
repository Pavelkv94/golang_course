# Go Demo - Learning Path

A collection of Go language examples and projects demonstrating various concepts from basics to advanced topics.

## 📚 Lessons Overview

### `lesson1/` - Basics & Data Structures
- Structs and custom types
- Methods on structs
- Arrays and slices (`[]`, `make()`, capacity vs length)
- Maps (creating, accessing, checking existence)
- Packages and imports
- User input with `bufio.Scanner`
- String manipulation (`strings.Fields`)
- Modules and package organization

### `lesson2/` - Interfaces & Concurrency Basics
- Interface implementation
- Dependency injection pattern
- Error handling (`error` type, custom errors)
- Goroutines (`go` keyword)
- Channels (`chan`, sending/receiving)
- Anonymous functions

### `lesson3/` - Channel Selection
- `select` statement for multiple channels
- Non-blocking channel operations
- Concurrent message handling
- Infinite loops with channels

### `lesson4/` - Channel Lifecycle
- Closing channels (`close()`)
- Detecting closed channels (`value, ok := <-ch`)
- Range over channels
- Graceful goroutine termination

### `lesson5/` - Context
- `context.Context` for cancellation
- Parent and child contexts
- `context.WithCancel()`
- Graceful shutdown of goroutines
- Context propagation

### `lesson6/` - WaitGroups
- `sync.WaitGroup` for goroutine synchronization
- `Add()`, `Done()`, `Wait()` methods
- Coordinating multiple goroutines without return values
- `defer` with WaitGroup

### `lesson7/` - Mutex
- `sync.Mutex` for data protection
- Race conditions and data races
- `Lock()` and `Unlock()`
- Critical sections
- Protecting shared resources

### `lesson8/` - RWMutex
- `sync.RWMutex` for read/write locks
- `RLock()` and `RUnlock()` for readers
- Performance optimization with concurrent reads
- Running with race detector (`go run -race`)

### `lesson9/` - Advanced Concurrency
- Complex concurrency patterns
- Worker pools
- Multiple channel coordination with `select`
- Context-based cancellation in pools
- Combining contexts, channels, and goroutines

## 🌐 HTTP & Web

### `http1/` - Basic HTTP Server
- `net/http` package basics
- HTTP handlers (`http.HandleFunc`)
- Request and response handling
- Reading request body (`io.ReadAll`)
- HTTP status codes
- Routing
- `sync.Mutex` with HTTP handlers
- `sync/atomic` for concurrent operations

### `http2/` - JSON & HTTP APIs
- JSON marshaling/unmarshaling
- Struct tags for JSON (`json:"field"`)
- Request validation
- Query parameters (`r.URL.Query()`)
- HTTP response formatting
- `json.MarshalIndent` for pretty printing
- Building RESTful endpoints

### `todo/` - Todo List REST API
- Complete REST API implementation
- Project structure (separation of concerns)
- API handlers and routing
- DTOs (Data Transfer Objects)
- Custom error types
- Business logic layer
- Chi router integration

## 🗄️ Database

### `postgres/` - PostgreSQL Integration
- PostgreSQL connection with `pgx/v5`
- Database migrations (`golang-migrate/migrate`)
- CRUD operations (Create, Read, Update, Delete)
- SQL queries in Go
- Connection pooling
- Environment variables for configuration
- Docker setup with Dockerfile
- Docker Compose for multi-container apps
- Entrypoint scripts for initialization
- Database initialization at container startup
- HTTP server with database backend
- Makefile for common tasks

## 🐳 DevOps & Deployment

The `postgres/` folder also demonstrates:
- Dockerfile best practices
- Multi-stage builds optimization
- Docker networking
- Container environment configuration
- Database migrations in Docker
- Health checks and startup scripts

## 🛠️ Tools & Dependencies

Common tools used across projects:
- `github.com/k0kubun/pp` - Pretty printer for debugging
- `github.com/jackc/pgx/v5` - PostgreSQL driver
- `github.com/golang-migrate/migrate/v4` - Database migrations
- `github.com/go-chi/chi/v5` - HTTP router

## 📖 Learning Path Recommendation

1. Start with `lesson1` through `lesson9` in order for core Go concepts
2. Move to `http1` and `http2` for web development basics
3. Study `todo` for API design and project structure
4. Finish with `postgres` for full-stack Go application with database

## 🚀 Running Examples

Each folder typically includes:
- `go.mod` - Module definition
- `main.go` - Entry point
- Additional packages for organization

Run any example:
```bash
cd <folder-name>
go run main.go
```

For the postgres project:
```bash
cd postgres
docker-compose up
```
