# How to Start the Go Backend Service

## Prerequisites

1. **Go 1.21+** installed ([Download](https://golang.org/dl/))
2. **PostgreSQL** running (local or Docker)
3. **Environment variables** configured

## Quick Start Steps

### Step 1: Navigate to Backend Directory

```bash
cd backend
```

### Step 2: Install Dependencies

```bash
go mod download
```

This will download all required Go packages listed in `go.mod`.

### Step 3: Set Up Environment Variables

**Create `.env` file in the backend directory:**

```bash
# On Windows (PowerShell)
Copy-Item .env.example .env

# On Linux/Mac
cp .env.example .env
```

**Or create `.env` manually** with the following content:

**Required Settings:**
```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=medical_user
DB_PASSWORD=medical_password
DB_NAME=medical_records
DB_SSLMODE=disable

# Server Configuration
SERVER_PORT=8080

# JWT Configuration (IMPORTANT: Change in production!)
JWT_SECRET=your-super-secret-jwt-key-change-in-production
JWT_EXPIRATION_HOURS=24
```

**Optional Settings** (for file uploads and notifications):
- AWS S3 credentials (for storing prescription PDFs and lab reports)
- SMTP settings (for email sharing)
- Twilio credentials (for SMS sharing)

### Step 4: Set Up Database

**Option A: Automatic Setup (Recommended)**
```bash
# This will create the database and user automatically
go run cmd/setup/main.go
```

**Option B: Using Docker**
```bash
# From project root
docker-compose up -d postgres
```

**Option C: Manual PostgreSQL Setup**
If you have PostgreSQL running, create the database manually:
```sql
CREATE DATABASE medical_records;
CREATE USER medical_user WITH PASSWORD 'medical_password';
GRANT ALL PRIVILEGES ON DATABASE medical_records TO medical_user;
```

### Step 5: Start the Server

**Option A: Direct Run (Development)**
```bash
go run cmd/server/main.go
```

**Option B: Build and Run**
```bash
# Build the binary
go build -o bin/server cmd/server/main.go

# Run the binary
./bin/server
```

**Option C: Using Makefile**
```bash
make run
```

### Step 6: Verify Server is Running

You should see output like:
```
Server starting on port 8080
```

The server will be available at:
- **API Base URL**: `http://localhost:8080/api/v1`
- **Health Check**: `http://localhost:8080/health`
- **Swagger Docs**: `http://localhost:8080/swagger/index.html`

## Troubleshooting

### Error: "Failed to connect to database"
- Check PostgreSQL is running: `pg_isready` or check Docker container
- Verify database credentials in `.env`
- Ensure database exists: `psql -U medical_user -d medical_records`

### Error: "Port already in use"
- Change `SERVER_PORT` in `.env` to a different port (e.g., 8081)
- Or stop the process using port 8080

### Error: "No .env file found"
- Make sure you've copied `.env.example` to `.env`
- Or set environment variables directly in your shell

### Error: "go: cannot find module"
- Run `go mod download` to install dependencies
- Run `go mod tidy` to clean up dependencies

## Development Tips

### Hot Reload (Optional)

For automatic reloading during development, use tools like:
- **Air**: `go install github.com/cosmtrek/air@latest` then `air`
- **Fresh**: `go install github.com/gravityblast/fresh@latest` then `fresh`

### Running Tests

```bash
go test ./...
```

### Building for Production

```bash
# Build optimized binary
go build -o bin/server cmd/server/main.go

# Or use Makefile
make build
```

## Environment Variables Reference

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `DB_HOST` | Yes | localhost | PostgreSQL host |
| `DB_PORT` | Yes | 5432 | PostgreSQL port |
| `DB_USER` | Yes | medical_user | Database user |
| `DB_PASSWORD` | Yes | medical_password | Database password |
| `DB_NAME` | Yes | medical_records | Database name |
| `SERVER_PORT` | No | 8080 | Server port |
| `JWT_SECRET` | Yes | - | Secret key for JWT tokens |
| `JWT_EXPIRATION_HOURS` | No | 24 | Token expiration time |

## Next Steps

Once the backend is running:
1. Start the frontend: `cd ../frontend && npm start`
2. Access Swagger docs: `http://localhost:8080/swagger/index.html`
3. Test the API: `curl http://localhost:8080/health`

