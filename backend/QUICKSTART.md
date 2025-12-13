# Quick Start Guide - Go Backend

## Minimum Setup (5 Steps)

### 1. Install Go Dependencies
```bash
cd backend
go mod download
```

### 2. Create .env File
Create a file named `.env` in the `backend` directory.

**If using existing PostgreSQL:**
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_postgres_username
DB_PASSWORD=your_postgres_password
DB_NAME=medical_records
DB_SSLMODE=disable
SERVER_PORT=8080
JWT_SECRET=change-this-to-a-random-secret-key-in-production
JWT_EXPIRATION_HOURS=24
APP_ENV=development
```

**If using Docker or new setup:**
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=medical_user
DB_PASSWORD=medical_password
DB_NAME=medical_records
DB_SSLMODE=disable
SERVER_PORT=8080
JWT_SECRET=change-this-to-a-random-secret-key-in-production
JWT_EXPIRATION_HOURS=24
APP_ENV=development
```

### 3. Set Up Database

**If using existing PostgreSQL:**
```sql
-- Connect to PostgreSQL and run:
CREATE DATABASE medical_records;
```

**If using Docker:**
```bash
# From project root directory
docker-compose up -d postgres
```

**If setting up new database:**
```bash
# Automatic setup (creates database and user)
go run cmd/setup/main.go
```

Or manually:
```sql
CREATE DATABASE medical_records;
CREATE USER medical_user WITH PASSWORD 'medical_password';
GRANT ALL PRIVILEGES ON DATABASE medical_records TO medical_user;
```

### 4. Start the Go Server
```bash
# The server will automatically create all tables on first run
go run cmd/server/main.go
```

### 5. Verify It's Running
Open your browser and visit:
- Health check: http://localhost:8080/health
- API docs: http://localhost:8080/swagger/index.html

You should see the server running on port 8080!

## Common Issues

**"Failed to connect to database"**
- Make sure PostgreSQL is running
- Check your database credentials in `.env`

**"Port 8080 already in use"**
- Change `SERVER_PORT=8081` in `.env` file

**"No .env file found"**
- Create `.env` file manually in the `backend` directory
- Copy the content from Step 2 above

