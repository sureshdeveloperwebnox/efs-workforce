# Gin Framework - Quick Reference Guide

## 🚀 Quick Start

### Start Server
```bash
cd /home/webnox/Development/Microservice-EFS/workforce/efs-workforce
go run cmd/server/main.go
```

### Check Health
```bash
curl http://localhost:8082/healthz
```

### Test API
```bash
# List all roles
curl http://localhost:8082/api/v1/roles

# Create a role
curl -X POST http://localhost:8082/api/v1/roles \
  -H "Content-Type: application/json" \
  -d '{"role_name":"Manager","description":"Team manager"}'

# Get specific role
curl http://localhost:8082/api/v1/roles/{id}

# Update role
curl -X PUT http://localhost:8082/api/v1/roles/{id} \
  -H "Content-Type: application/json" \
  -d '{"role_name":"Senior Manager","description":"Updated"}'

# Delete role
curl -X DELETE http://localhost:8082/api/v1/roles/{id}
```

## 📁 Project Structure

```
efs-workforce/
├── cmd/server/main.go                      # Entry point (Gin-based)
├── internal/adapters/inbound/http/
│   ├── gin_server.go                       # Main server setup
│   ├── handlers/role_handler.go            # Request handlers
│   ├── middleware/                         # Custom middleware
│   ├── routes/routes.go                    # Route definitions
│   └── utils/response.go                   # Response helpers
├── GIN_ARCHITECTURE.md                     # Full architecture docs
└── MIGRATION_SUMMARY.md                    # Migration details
```

## 🔧 Adding New Endpoints

### 1. Create Handler
```go
// internal/adapters/inbound/http/handlers/user_handler.go
type UserHandler struct {
    userService *application.UserService
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    var req dto.CreateUserRequest
    if !utils.BindJSON(c, &req) {
        return
    }
    
    resp, err := h.userService.CreateUser(&req)
    if err != nil {
        utils.HandleError(c, err)
        return
    }
    
    utils.RespondCreated(c, resp)
}
```

### 2. Add Routes
```go
// internal/adapters/inbound/http/routes/routes.go
func SetupUserRoutes(rg *gin.RouterGroup, userHandler *handlers.UserHandler) {
    users := rg.Group("/users")
    {
        users.POST("", userHandler.CreateUser)
        users.GET("", userHandler.ListUsers)
        users.GET("/:id", userHandler.GetUser)
        users.PUT("/:id", userHandler.UpdateUser)
        users.DELETE("/:id", userHandler.DeleteUser)
    }
}
```

### 3. Register in Server
```go
// In SetupV1Routes
SetupUserRoutes(v1, userHandler)
```

## 🎯 Response Helpers

```go
// Success responses
utils.RespondOK(c, data)                    // 200 OK
utils.RespondCreated(c, data)               // 201 Created
utils.RespondNoContent(c)                   // 204 No Content

// Error handling
utils.HandleError(c, err)                   // Auto-maps to correct status

// Custom response
utils.RespondSuccess(c, statusCode, data, "message")
```

## 🛡️ Middleware

### Available Middleware
- `middleware.Recovery()` - Panic recovery
- `middleware.Logger()` - Request logging
- `middleware.CORS()` - CORS handling
- `middleware.RequestID()` - Request tracking
- `middleware.SecurityHeaders()` - Security headers
- `middleware.RateLimiter()` - Rate limiting (placeholder)

### Apply Middleware
```go
// Global
router.Use(middleware.Logger())

// Route group
v1.Use(middleware.RateLimiter())

// Single route
router.GET("/special", middleware.Custom(), handler)
```

## 📊 Monitoring

### Logs
```bash
# View live logs
tail -f server.log

# Search logs
grep "ERROR" server.log
grep "api/v1/roles" server.log
```

### Health Check
```bash
# Simple check
curl http://localhost:8082/healthz

# With details
curl -s http://localhost:8082/healthz | jq .
```

## 🐛 Debugging

### Common Issues

**Port already in use:**
```bash
fuser -k 8082/tcp
fuser -k 50056/tcp
```

**Database connection:**
```bash
# Check DATABASE_URL in .env
cat .env | grep DATABASE_URL

# Test connection
psql $DATABASE_URL -c "SELECT 1"
```

**Import errors:**
```bash
go mod tidy
go mod download
```

## 📈 Performance Tips

1. **Use connection pooling** (already configured)
2. **Enable Gin release mode** in production:
   ```go
   gin.SetMode(gin.ReleaseMode)
   ```
3. **Add caching** for frequently accessed data
4. **Use database indexes** on queried fields
5. **Implement rate limiting** for public APIs

## 🔐 Security Checklist

- ✅ CORS configured
- ✅ Security headers enabled
- ✅ Panic recovery active
- ✅ Request ID tracking
- ⏳ Rate limiting (to implement)
- ⏳ Authentication middleware (to implement)
- ⏳ Input validation (to implement)

## 📚 Resources

- **Gin Documentation**: https://gin-gonic.com/docs/
- **Architecture Guide**: `GIN_ARCHITECTURE.md`
- **Migration Summary**: `MIGRATION_SUMMARY.md`

## 🆘 Quick Commands

```bash
# Start server
go run cmd/server/main.go

# Run tests
go test ./...

# Build binary
go build -o workforce-service cmd/server/main.go

# Format code
go fmt ./...

# Check for issues
go vet ./...

# Update dependencies
go mod tidy

# Rollback to Gorilla Mux
mv cmd/server/main.go.backup cmd/server/main.go
```

## 🎨 Code Style

### Handler Pattern
```go
func (h *Handler) Action(c *gin.Context) {
    // 1. Bind request
    var req dto.Request
    if !utils.BindJSON(c, &req) {
        return
    }
    
    // 2. Call service
    resp, err := h.service.Action(&req)
    if err != nil {
        utils.HandleError(c, err)
        return
    }
    
    // 3. Return response
    utils.RespondOK(c, resp)
}
```

### Error Handling
```go
// Service layer returns errors
if err != nil {
    return nil, fmt.Errorf("resource not found")
}

// Handler uses utils.HandleError
// Automatically maps to 404, 409, 500, etc.
```

## 🚦 Status Codes

- `200` - OK
- `201` - Created
- `204` - No Content
- `400` - Bad Request
- `404` - Not Found
- `409` - Conflict
- `500` - Internal Server Error

---

**Framework**: Gin v1.11.0
**Port**: 8082 (HTTP), 50056 (gRPC)
**Status**: ✅ Production Ready
