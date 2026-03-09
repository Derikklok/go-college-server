# College Server API

A production-ready Go REST API for managing college student records. Built with Gin, GORM, and MySQL.

## Features

- ✅ RESTful API for managing students
- ✅ MySQL database integration with GORM ORM
- ✅ Scalable migration system
- ✅ Clean architecture (Controller → Service → Repository)
- ✅ Environment-based configuration
- ✅ Docker support
- ✅ Error handling and structured responses

## Project Structure

```
go-college-server/
├── cmd/
│   ├── api/
│   │   └── main.go              # Application entry point
│   └── migrate/
│       └── main.go              # Migration CLI tool
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── database/
│   │   └── mysql.go             # Database connection
│   ├── migrations/
│   │   └── migrations.go        # Migration registry
│   ├── shared/
│   │   └── response.go          # Shared response types
│   └── student/
│       ├── student.controller.go
│       ├── student.service.go
│       ├── student.repository.go
│       ├── student.model.go
│       ├── student.dto.go
│       ├── student.routes.go
│       ├── migrations.go
│       └── student.migration.go
├── build/                       # Build output directory
├── .env                         # Environment variables
├── .env.prod                    # Production environment variables
├── .gitignore
├── Dockerfile                   # Docker configuration
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums
├── README.md                    # This file
├── MIGRATIONS.md                # Migration documentation
├── migrate.sh                   # Linux/Mac migration script
├── migrate.bat                  # Windows migration script
└── Dockerfile
```

## Installation

### Prerequisites

- Go 1.19 or higher
- MySQL 5.7 or higher
- Git

### Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/Derikklok/go-college-server.git
   cd go-college-server
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Configure environment variables**
   
   Create a `.env` file in the project root:
   ```env
   APP_PORT=8080
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=your_password
   DB_NAME=college_db
   ```

4. **Run migrations**
   ```bash
   # Linux/Mac
   ./migrate.sh
   
   # Windows
   ./migrate.bat
   
   # Or directly
   go run ./cmd/migrate
   ```

5. **Start the server**
   ```bash
   go run ./cmd/api
   ```

   Server will run on `http://localhost:8080`

## API Documentation

### Base URL
```
http://localhost:8080/api
```

### Endpoints

#### Create Student
```http
POST /api/students
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 20
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "ID": 1,
    "Name": "John Doe",
    "Email": "john@example.com",
    "Age": 20,
    "CreatedAt": "2026-03-09T10:30:00Z",
    "UpdatedAt": "2026-03-09T10:30:00Z"
  }
}
```

#### Get All Students
```http
GET /api/students
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "ID": 1,
      "Name": "John Doe",
      "Email": "john@example.com",
      "Age": 20,
      "CreatedAt": "2026-03-09T10:30:00Z",
      "UpdatedAt": "2026-03-09T10:30:00Z"
    }
  ]
}
```

#### Delete Student
```http
DELETE /api/students/:id
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": "Student deleted"
}
```

### Error Response
```json
{
  "success": false,
  "message": "Failed to create student",
  "error": "error details here"
}
```

## Database Migrations

The project uses a migration registry system for managing database schema changes. See [MIGRATIONS.md](MIGRATIONS.md) for detailed information.

### Quick Migration Commands

```bash
# Run all migrations
go run ./cmd/migrate

# Via shell script (Linux/Mac)
./migrate.sh

# Via batch script (Windows)
./migrate.bat
```

## Docker

### Build Docker Image

```bash
docker build -t college-server:latest .
```

### Run with Docker

```bash
docker run -p 8080:8080 \
  -e DB_HOST=host.docker.internal \
  -e DB_PORT=3306 \
  -e DB_USER=root \
  -e DB_PASSWORD=your_password \
  -e DB_NAME=college_db \
  --name college-server \
  college-server:latest
```

### Docker Compose (Optional)

Create a `docker-compose.yml`:

```yaml
version: '3.8'

services:
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: sachin1605
      MYSQL_DATABASE: college_db
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      DB_HOST: mysql
      DB_PORT: 3306
      DB_USER: root
      DB_PASSWORD: sachin1605
      DB_NAME: college_db
    depends_on:
      - mysql

volumes:
  mysql_data:
```

Run with:
```bash
docker-compose up
```

## Development

### Building the Project

```bash
# Build API
go build -o ./build/api ./cmd/api

# Build Migration Tool
go build -o ./build/migrate ./cmd/migrate
```

### Running Tests

```bash
go test ./...
```

### Code Structure

- **Controller**: Handles HTTP requests and responses
- **Service**: Contains business logic
- **Repository**: Handles database operations
- **Model**: Defines data structures
- **DTO**: Data Transfer Objects for request/response

This follows the **repository pattern** for clean, testable code.

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `APP_PORT` | 8080 | Server port |
| `DB_HOST` | localhost | MySQL host |
| `DB_PORT` | 3306 | MySQL port |
| `DB_USER` | root | MySQL username |
| `DB_PASSWORD` | - | MySQL password |
| `DB_NAME` | college_db | Database name |

## Adding New Services

To add a new service (e.g., Courses):

1. Create `internal/course/` directory
2. Add model, DTO, controller, service, repository, routes files
3. Create `internal/course/migrations.go` and `internal/course/course.migration.go`
4. Import in `cmd/migrate/main.go`
5. Register routes in `cmd/api/main.go`

See [MIGRATIONS.md](MIGRATIONS.md) for detailed migration setup.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For issues and questions, please open an issue on GitHub or contact the maintainers.

## Roadmap

- [ ] Authentication & Authorization
- [ ] Student enrollment management
- [ ] Course management
- [ ] Grade tracking
- [ ] Attendance management
- [ ] Email notifications
- [ ] API documentation (Swagger/OpenAPI)
- [ ] Unit and integration tests
- [ ] API rate limiting
- [ ] Logging and monitoring

## Acknowledgments

- Built with [Gin Web Framework](https://github.com/gin-gonic/gin)
- ORM by [GORM](https://gorm.io/)
- Configuration management with [godotenv](https://github.com/joho/godotenv)
