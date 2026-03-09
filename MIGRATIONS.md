# Database Migrations

This project uses a migration registry system that allows you to:
- Run migrations independently without starting the main application
- Separate migrations by service/model
- Easily add new model migrations

## Migration Structure

```
internal/
├── migrations/
│   └── migrations.go          # Migration registry and runner
├── student/
│   ├── migrations.go          # Student model migrations
│   └── student.migration.go   # Registers student migrations
```

## Running Migrations

### Method 1: Standalone Migration CLI (Recommended)

Run migrations without starting the main application:

```bash
# Build the migration tool
go build -o ./build/migrate ./cmd/migrate

# Run migrations
./build/migrate
```

Or in one command:
```bash
go run ./cmd/migrate
```

### Method 2: Via Main Application

Migrations run automatically when the application starts:

```bash
go run ./cmd/api
```

## Adding New Model Migrations

When you add a new service/model, follow these steps:

1. Create migration file (e.g., `internal/course/migrations.go`):

```go
package course

import "gorm.io/gorm"

type CourseMigration struct{}

func (cm *CourseMigration) Migrate(db *gorm.DB) error {
    return db.AutoMigrate(&Course{})
}

func NewCourseMigration() *CourseMigration {
    return &CourseMigration{}
}
```

2. Create registration file (e.g., `internal/course/course.migration.go`):

```go
package course

import (
    "github.com/Derikklok/go-college-server/internal/migrations"
)

func init() {
    migrations.Register(NewCourseMigration())
}
```

3. Import the package in `cmd/migrate/main.go`:

```go
import (
    _ "github.com/Derikklok/go-college-server/internal/course"
)
```

## Environment Setup

Ensure `.env` file is configured with database credentials:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=college_db
```

## Migration Best Practices

- ✅ Use the migration registry for all model changes
- ✅ Keep migrations atomic (one logical change per migration)
- ✅ Always test migrations locally before deployment
- ✅ Use the standalone CLI for production migrations
- ✅ Register all models in init() function
