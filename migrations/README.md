# Database Migrations

This directory is intended for database migration files.

## Current Setup

The application currently uses GORM's AutoMigrate feature, which can be enabled by setting `AUTO_MIGRATE=true` in your environment variables.

## Recommended for Production

For production environments, it's recommended to use a proper migration tool like [golang-migrate](https://github.com/golang-migrate/migrate) instead of AutoMigrate.

### Installing golang-migrate

```bash
# macOS
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/
```

### Creating Migrations

```bash
# Create a new migration
migrate create -ext sql -dir migrations -seq create_users_table
```

### Running Migrations

```bash
# Apply all migrations
migrate -path migrations -database "postgres://user:pass@host:port/dbname?sslmode=require" up

# Rollback last migration
migrate -path migrations -database "postgres://user:pass@host:port/dbname?sslmode=require" down 1
```

## Example Migration Files

Migration files should be placed in this directory with the following naming convention:
- `000001_create_users_table.up.sql`
- `000001_create_users_table.down.sql`

### Sample Up Migration

```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    plan VARCHAR(50) DEFAULT 'trial',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Sample Down Migration

```sql
DROP TABLE IF EXISTS users;
```

## Integration with Application

To integrate golang-migrate with your application:

1. Disable AUTO_MIGRATE in production
2. Run migrations as part of your deployment process
3. Consider using a migration runner in your CI/CD pipeline
