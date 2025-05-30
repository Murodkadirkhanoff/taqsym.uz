# Taqsym.uz Microservices Architecture

This repository contains a collection of microservices for the Taqsym.uz platform, built using Go and following clean architecture principles.

## Architecture Overview

The project follows a microservices architecture with the following components:

1. **API Gateway**: Entry point for all client requests, handles routing to appropriate services
2. **User Service**: Manages user accounts, authentication, and authorization
3. **Task Service**: Manages tasks and their assignments
4. **Product Service**: Manages products and categories
5. **Account Service**: Manages financial accounts and transactions
6. **Catalog Service**: Manages product catalogs
7. **Order Service**: Manages orders and order processing

Each microservice is built following clean architecture principles, with clear separation of concerns:

```
service/
├── cmd/                  # Application entry points
│   └── main.go           # Main application
├── config/               # Configuration
│   └── config.yml        # Configuration file
├── internal/             # Private application code
│   ├── api/              # API layer (HTTP handlers)
│   ├── domain/           # Domain models and interfaces
│   ├── infrastructure/   # Infrastructure implementations
│   ├── repository/       # Data access implementations
│   ├── usecase/          # Business logic implementations
│   └── config/           # Configuration code
└── app.dockerfile        # Dockerfile for the service
```

## Clean Architecture

Each microservice follows the clean architecture pattern with the following layers:

1. **Domain Layer**: Contains the business entities and repository interfaces
   - Located in `internal/domain/`
   - Defines the core business models and interfaces
   - Has no dependencies on other layers

2. **Use Case Layer**: Contains the business logic
   - Located in `internal/usecase/`
   - Implements the business rules
   - Depends only on the domain layer

3. **Repository Layer**: Contains the data access implementations
   - Located in `internal/repository/`
   - Implements the repository interfaces defined in the domain layer
   - Depends on the domain layer and infrastructure layer

4. **API Layer**: Contains the HTTP handlers
   - Located in `internal/api/`
   - Handles HTTP requests and responses
   - Depends on the use case layer

5. **Infrastructure Layer**: Contains external implementations
   - Located in `internal/infrastructure/`
   - Implements database connections, external services, etc.
   - Used by the repository layer

## Dependency Flow

The dependencies flow from the outer layers to the inner layers:

```
API Layer → Use Case Layer → Domain Layer ← Repository Layer ← Infrastructure Layer
```

This ensures that the domain layer remains independent of implementation details.

## Best Practices

### Code Organization

1. **Package Structure**: Organize code by layer, not by feature
2. **Interface Segregation**: Define small, focused interfaces
3. **Dependency Injection**: Use constructor injection to provide dependencies
4. **Error Handling**: Return descriptive errors and handle them appropriately
5. **Configuration**: Use environment variables and configuration files

### API Design

1. **RESTful Endpoints**: Follow REST principles for API design
2. **Consistent Response Format**: Use consistent JSON response format
3. **Proper Status Codes**: Use appropriate HTTP status codes
4. **Validation**: Validate input data before processing
5. **Pagination**: Implement pagination for list endpoints

### Database Access

1. **Repository Pattern**: Use the repository pattern to abstract data access
2. **Parameterized Queries**: Use parameterized queries to prevent SQL injection
3. **Transaction Management**: Use transactions for operations that require atomicity
4. **Connection Pooling**: Configure connection pooling appropriately

### Containerization

1. **Multi-stage Builds**: Use multi-stage builds to minimize image size
2. **Environment Variables**: Use environment variables for configuration
3. **Health Checks**: Implement health check endpoints
4. **Graceful Shutdown**: Handle shutdown signals gracefully

## Getting Started

### Prerequisites

- Go 1.24 or later
- Docker and Docker Compose
- PostgreSQL

### Running the Services

1. Clone the repository:
   ```
   git clone https://github.com/Murodkadirkhanoff/taqsym.uz.git
   cd taqsym.uz
   ```

2. Start the services using Docker Compose:
   ```
   docker-compose up -d
   ```

3. Access the services:
   - API Gateway: http://localhost:8000
   - User Service: http://localhost:8081
   - Product Service: http://localhost:8082
   - Account Service: http://localhost:8001
   - Catalog Service: http://localhost:8002
   - Order Service: http://localhost:8003

## Development

### Adding a New Service

1. Create a new directory for the service:
   ```
   mkdir -p new-service/{cmd,config,internal/{api,domain,infrastructure,repository,usecase,config}}
   ```

2. Implement the required layers following the clean architecture pattern

3. Add the service to the docker-compose.yml file

4. Update the API Gateway to route requests to the new service

### Testing

Each service should have unit tests for the domain and use case layers, and integration tests for the repository and API layers.

Run tests with:
```
go test ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.