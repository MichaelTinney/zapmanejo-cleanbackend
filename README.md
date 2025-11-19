# zapmanejo-cleanbackend

Clean backend API for ZapManejo - livestock management system.

## Overview

ZapManejo is a livestock management application that helps ranchers track their animals, health records, and manage subscriptions via PayPal. This is the backend API built with Go, Fiber, and PostgreSQL.

## Features

- 🔐 JWT-based authentication
- 🐄 Animal management (CRUD operations)
- 🏥 Health records tracking
- 💳 Payment processing with PayPal
- 📱 WhatsApp webhook integration
- 🔄 Graceful shutdown handling
- 🏥 Health check endpoint for cloud deployments

## Prerequisites

- Go 1.23 or higher
- PostgreSQL database
- (Optional) Docker for containerized deployment

## Environment Variables

Copy `.env.example` to `.env` and configure the following variables:

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `DATABASE_URL` | PostgreSQL connection string | Yes | - |
| `PORT` | Server port | No | 5000 |
| `JWT_SECRET` | Secret key for JWT tokens | Yes | - |
| `AUTO_MIGRATE` | Run migrations on startup | No | false |
| `CORS_ALLOW_ORIGINS` | CORS allowed origins | No | * |
| `SHUTDOWN_TIMEOUT_SECONDS` | Graceful shutdown timeout | No | 30 |
| `PAYPAL_CLIENT_ID` | PayPal client ID | Yes | - |
| `PAYPAL_SECRET` | PayPal secret key | Yes | - |
| `PAYPAL_MODE` | PayPal mode (sandbox/live) | No | sandbox |
| `WHATSAPP_NUMBER` | WhatsApp number for webhooks | Yes | - |
| `WHATSAPP_VERIFY_TOKEN` | WhatsApp webhook verify token | Yes | - |
| `LIFETIME_TOTAL` | Total lifetime slots available | No | 200 |
| `LIFETIME_SOLD` | Number of lifetime slots sold | No | 0 |

## Local Development

### Install dependencies
```bash
go mod download
```

### Run the application
```bash
# Using go run
go run .

# Or using make
make run
```

### Build the application
```bash
make build
```

### Run tests
```bash
make test
```

### Run linter
```bash
make lint
```

## Deployment to DigitalOcean App Platform

### Prerequisites

1. A DigitalOcean account
2. A managed PostgreSQL database on DigitalOcean (or other provider)
3. This repository pushed to GitHub

### Step-by-Step Deployment Checklist

- [ ] **1. Create a PostgreSQL Database**
  - Go to DigitalOcean → Databases → Create Database
  - Choose PostgreSQL, select region and plan
  - Note down the connection details (see screenshot below for reference)

![DigitalOcean Database Connection Details](https://placeholder-for-screenshot)

- [ ] **2. Create an App on DigitalOcean App Platform**
  - Go to DigitalOcean → Apps → Create App
  - Connect your GitHub repository
  - Select the branch to deploy (e.g., `main`)
  - DigitalOcean will auto-detect the Dockerfile

- [ ] **3. Configure Build Settings**
  - **Build Command:** (Auto-detected from Dockerfile)
  - **Run Command:** (Auto-detected from Dockerfile)
  - Dockerfile path: `Dockerfile` (in root)

- [ ] **4. Set Environment Variables**
  - In App Settings → Environment Variables, add:
    - `DATABASE_URL`: Copy the connection string from your database panel
      - Format: `postgres://username:password@host:port/database?sslmode=require`
      - ⚠️ **Use the "Connection String" from your database, not the public one**
    - `JWT_SECRET`: Generate with `openssl rand -base64 32`
    - `AUTO_MIGRATE`: Set to `true` for first deployment
    - `PAYPAL_CLIENT_ID`: Your PayPal client ID
    - `PAYPAL_SECRET`: Your PayPal secret
    - `WHATSAPP_VERIFY_TOKEN`: Your WhatsApp verify token
    - Other variables as needed from `.env.example`

- [ ] **5. Deploy the Application**
  - Click "Create Resources" or "Deploy"
  - Wait for the build and deployment to complete
  - Monitor the deployment logs for any errors

- [ ] **6. Verify Deployment**
  - Access the health endpoint: `https://your-app.ondigitalocean.app/healthz`
  - Should return: `{"status":"healthy","service":"zapmanejo-backend"}`
  - Check application logs in the DigitalOcean dashboard

- [ ] **7. Post-Deployment**
  - Set `AUTO_MIGRATE=false` after first successful deployment
  - Consider setting up a custom domain
  - Enable application metrics and monitoring
  - Set up alerts for application health

### Important Security Notes

- ⚠️ **Never commit secrets or credentials to the repository**
- 🔄 **Rotate all credentials immediately if they are accidentally exposed**
- 🔐 **Use strong, randomly generated values for `JWT_SECRET` and `WHATSAPP_VERIFY_TOKEN`**
- 🌐 **In production, set `CORS_ALLOW_ORIGINS` to your specific frontend domain**
- 📊 **Monitor application logs regularly for security issues**

### Getting Database Connection String

The `DATABASE_URL` should be obtained from your DigitalOcean database panel:

1. Go to Databases → Your Database → Connection Details
2. Select "Connection String" or "Connection Parameters"
3. Copy the full connection string including username, password, host, and port
4. The format should be: `postgres://username:password@host:port/database?sslmode=require`

**Important:** If you accidentally expose your database password in commits:
1. Immediately reset the database password in DigitalOcean
2. Update the `DATABASE_URL` in your App Platform settings
3. Redeploy the application

## API Endpoints

### Health Check
- `GET /healthz` - Health check endpoint (returns 200 OK)

### Authentication
- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login and get JWT token

### Animals (Protected)
- `GET /api/animals` - Get all animals for user
- `POST /api/animals` - Create new animal
- `PUT /api/animals/:id` - Update animal
- `DELETE /api/animals/:id` - Delete animal

### Health Records (Protected)
- `GET /api/health` - Get all health records
- `POST /api/health` - Create health record

### Payments (Protected)
- `POST /api/payment/create` - Create PayPal order
- `POST /api/payment/capture` - Capture PayPal payment

### WhatsApp Webhook
- `GET /webhook` - Verify webhook
- `POST /webhook` - Handle incoming messages

## Database Migrations

The application supports two migration approaches:

1. **AutoMigrate (Development)**: Set `AUTO_MIGRATE=true` to automatically run GORM migrations
2. **Manual Migrations (Production)**: Use [golang-migrate](https://github.com/golang-migrate/migrate) for production

See [migrations/README.md](migrations/README.md) for detailed migration instructions.

## Docker

### Build Docker image
```bash
make docker-build
```

### Run with Docker
```bash
docker run -p 5000:5000 --env-file .env zapmanejo-backend:latest
```

## CI/CD

This repository includes a GitHub Actions workflow that:
- Runs on every push to `main` and on pull requests
- Checks out code and sets up Go 1.23
- Downloads dependencies
- Builds the application
- Runs tests
- Runs golangci-lint for code quality

See [.github/workflows/go-ci.yml](.github/workflows/go-ci.yml) for details.

## Development Tools

- **Makefile**: Common development commands (`make help` for all options)
- **golangci-lint**: Code quality and linting (`.golangci.yml` for configuration)
- **Docker**: Containerized builds and deployments

## Project Structure

```
.
├── .github/
│   └── workflows/
│       └── go-ci.yml          # GitHub Actions CI workflow
├── internal/
│   ├── database/              # Database connection and migrations
│   ├── middleware/            # JWT middleware
│   ├── models/                # Data models
│   ├── routes/                # API route handlers
│   └── utils/                 # Utility functions
├── migrations/                # Database migration files
├── .env.example               # Environment variables template
├── .gitignore                 # Git ignore rules
├── .golangci.yml              # Linter configuration
├── Dockerfile                 # Multi-stage Docker build
├── Makefile                   # Development commands
├── go.mod                     # Go module definition
├── go.sum                     # Go module checksums
├── main.go                    # Application entry point
└── README.md                  # This file
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

See LICENSE file for details.

## Support

For issues or questions, please open an issue on GitHub.
