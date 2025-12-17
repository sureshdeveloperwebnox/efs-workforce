# Workforce Service - Microservice Implementation

A comprehensive workforce management microservice built with Go, following Clean Architecture principles and microservices best practices.

## Architecture Overview

This service implements Clean Architecture with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────┐
│                    Presentation Layer                    │
│              (gRPC Handlers / HTTP Handlers)            │
└──────────────────────┬──────────────────────────────────┘
                       │
┌──────────────────────▼──────────────────────────────────┐
│                  Application Layer                       │
│         (Use Cases / Business Logic / DTOs)              │
└──────────────────────┬──────────────────────────────────┘
                       │
┌──────────────────────▼──────────────────────────────────┐
│                    Domain Layer                         │
│              (Entities / Business Rules)                 │
└──────────────────────┬──────────────────────────────────┘
                       │
┌──────────────────────▼──────────────────────────────────┐
│              Infrastructure Layer                        │
│    (PostgreSQL / Kafka / Adapters/Outbound)             │
└─────────────────────────────────────────────────────────┘
```

## Features

- ✅ Clean Architecture with clear layer separation
- ✅ gRPC API for internal service communication
- ✅ HTTP REST API for external access
- ✅ PostgreSQL for data persistence with UUID primary keys
- ✅ Kafka event publishing for async communication
- ✅ Complete CRUD operations for all workforce entities
- ✅ Database migrations with GORM auto-migration
- ✅ Docker support

## Domain Entities

The workforce service manages the following entities:

1. **Roles** - User roles (Admin, Manager, Technician, etc.)
2. **Users** - Workforce users with employee information
3. **Crews** - Groups of users working together
4. **CrewMembers** - Association between crews and users
5. **Equipment** - Equipment assigned to users
6. **Attendance** - User attendance tracking
7. **TimeOff** - Leave/time off requests
8. **Trips** - User trip tracking
9. **Permissions** - Role-based permissions for modules

## Project Structure

```
efs-workforce/
├── cmd/
│   └── server/              # Application entry point
│       └── main.go
├── internal/                # Private application code
│   ├── domain/              # Business entities and rules
│   │   ├── role.go
│   │   ├── user.go
│   │   ├── crew.go
│   │   ├── equipment.go
│   │   ├── attendance.go
│   │   ├── timeoff.go
│   │   ├── trip.go
│   │   └── permission.go
│   ├── ports/               # Interfaces (repositories, services)
│   │   ├── repositories/    # Repository interfaces
│   │   ├── services/        # Service interfaces
│   │   └── external/        # External service interfaces
│   ├── adapters/            # Implementations
│   │   ├── inbound/         # gRPC/HTTP handlers
│   │   │   ├── grpc/
│   │   │   └── http/
│   │   └── outbound/         # Database, Kafka
│   │       ├── postgres/
│   │       └── kafka/
│   ├── application/         # Use cases and business logic
│   │   ├── dto/             # Data Transfer Objects
│   │   ├── role_service.go
│   │   └── ...
│   ├── config/              # Configuration management
│   └── database/            # Database initialization
├── migrations/              # SQL migration files
├── pkg/                     # Public reusable libraries
│   ├── logger/              # Logging utilities
│   ├── errors/              # Error handling
│   └── validator/          # Validation utilities
├── proto/                   # gRPC protocol definitions
└── tests/                   # Test files
```

## Prerequisites

- Go 1.24+
- PostgreSQL 12+
- Kafka (optional, for event publishing)
- Protocol Buffers compiler

## Setup

### 1. Clone and Install Dependencies

```bash
cd efs-workforce
go mod download
```

### 2. Setup Database

Create a PostgreSQL database:

```bash
createdb workforce_db
```

Or using psql:

```sql
CREATE DATABASE workforce_db;
```

**Note**: Database schema is automatically created/updated by GORM when the service starts. You can also run the SQL migrations manually from `migrations/001_create_workforce_tables.sql`.

### 3. Configure Environment

Copy `.env.example` to `.env` and update values:

```bash
cp .env.example .env
```

Edit `.env`:

```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/workforce_db?sslmode=disable
GRPC_PORT=50051
HTTP_PORT=8081
KAFKA_BROKERS=localhost:9092
KAFKA_TOPIC_PREFIX=workforce
```

### 4. Run the Service

```bash
go run cmd/server/main.go
```

Or using Make:

```bash
make run
```

The service will start on:
- gRPC port: `50051` (default)
- HTTP port: `8081` (default)

## API Usage

### Health Check

```bash
curl http://localhost:8081/healthz
```

### Roles API

#### Create Role
```bash
curl -X POST http://localhost:8081/api/v1/roles \
  -H "Content-Type: application/json" \
  -d '{
    "role_name": "Manager",
    "description": "Manager role with team management capabilities"
  }'
```

#### Get Role
```bash
curl http://localhost:8081/api/v1/roles/{role_id}
```

#### List Roles
```bash
curl http://localhost:8081/api/v1/roles
```

#### Update Role
```bash
curl -X PUT http://localhost:8081/api/v1/roles/{role_id} \
  -H "Content-Type: application/json" \
  -d '{
    "role_name": "Senior Manager",
    "description": "Updated description"
  }'
```

#### Delete Role
```bash
curl -X DELETE http://localhost:8081/api/v1/roles/{role_id}
```

### Users API

Similar endpoints are available for Users, Crews, Equipment, Attendance, TimeOff, Trips, and Permissions.

## Database Schema

The service uses PostgreSQL with UUID primary keys for better distributed system support. Key tables:

- `roles` - User roles
- `users` - Workforce users
- `crews` - Crew groups
- `crew_members` - Crew-user associations
- `equipment` - Equipment inventory
- `attendance` - Attendance records
- `time_off` - Leave requests
- `trips` - Trip records
- `permissions` - Role permissions

See `migrations/001_create_workforce_tables.sql` for the complete schema.

## Event Publishing

The service publishes domain events to Kafka for async communication with other microservices:

- `workforce.RoleCreated`
- `workforce.RoleUpdated`
- `workforce.RoleDeleted`
- (Similar events for other entities)

Events are published to topics following the pattern: `{KAFKA_TOPIC_PREFIX}.{EventType}`

## Development

### Adding a New Feature

1. **Domain**: Add entities/rules in `internal/domain/`
2. **Ports**: Define interfaces in `internal/ports/`
3. **Application**: Implement use cases in `internal/application/`
4. **Adapters**: Implement interfaces in `internal/adapters/`
5. **Handler**: Wire up in `cmd/server/main.go`

### Code Generation

If you modify proto files:

```bash
make proto
```

## Docker

### Build and Run

```bash
docker-compose up -d
```

### View Logs

```bash
docker-compose logs -f workforce-service
```

## Testing

### Unit Tests

```bash
go test ./internal/domain/...
go test ./internal/application/...
```

### Integration Tests

```bash
go test ./tests/integration/...
```

## Deployment

### Build Binary

```bash
make build
```

### Docker Build

```bash
docker build -t efs-workforce .
docker run -p 50051:50051 -p 8081:8081 --env-file .env efs-workforce
```

## Architecture Principles

This service follows microservices best practices:

1. **Single Responsibility** - Each service handles one bounded context
2. **Loose Coupling** - Services communicate via well-defined interfaces
3. **Data Ownership** - Each service owns its database
4. **Event-Driven** - Uses Kafka for async communication
5. **API Gateway** - Designed to work with KrakenD gateway
6. **Observability** - Structured logging and metrics ready

## Contributing

1. Follow Clean Architecture principles
2. Write tests for new features
3. Update documentation
4. Follow Go best practices

## License

MIT
