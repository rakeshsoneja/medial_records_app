# Database Setup Guide

## Error: Password Authentication Failed

If you're getting `password authentication failed for user "medical_user"`, you need to set up the database first.

## Option 1: Using Docker (Recommended - Easiest)

### Step 1: Start PostgreSQL Container
```bash
# From project root directory
docker-compose up -d postgres
```

This will automatically:
- Create the database `medical_records`
- Create user `medical_user` with password `medical_password`
- Set up all permissions

### Step 2: Verify Database is Running
```bash
docker ps
```

You should see a container named `medical_records_db` running.

### Step 3: Start Your Go Server
```bash
cd backend
go run cmd/server/main.go
```

## Option 2: Local PostgreSQL Setup

### Step 1: Connect to PostgreSQL
```bash
# On Windows (if PostgreSQL is in PATH)
psql -U postgres

# Or use pgAdmin or another PostgreSQL client
```

### Step 2: Create Database and User
Run these SQL commands in PostgreSQL:

```sql
-- Create the database
CREATE DATABASE medical_records;

-- Create the user
CREATE USER medical_user WITH PASSWORD 'medical_password';

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE medical_records TO medical_user;

-- Connect to the database
\c medical_records

-- Grant schema privileges (PostgreSQL 15+)
GRANT ALL ON SCHEMA public TO medical_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO medical_user;
```

### Step 3: Update .env File
Make sure your `backend/.env` file has:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=medical_user
DB_PASSWORD=medical_password
DB_NAME=medical_records
DB_SSLMODE=disable
```

### Step 4: Test Connection
```bash
# Test connection (if psql is available)
psql -h localhost -U medical_user -d medical_records
# Enter password: medical_password
```

## Option 3: Use Different Database Credentials

If you want to use your existing PostgreSQL setup:

### Step 1: Update .env File
Edit `backend/.env` with your actual credentials:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database_name
DB_SSLMODE=disable
```

### Step 2: Create Database (if needed)
```sql
CREATE DATABASE your_database_name;
```

## Troubleshooting

### "Connection refused" Error
- Make sure PostgreSQL is running
- Check if PostgreSQL is listening on port 5432
- Verify firewall settings

### "Database does not exist" Error
- Create the database: `CREATE DATABASE medical_records;`

### "User does not exist" Error
- Create the user: `CREATE USER medical_user WITH PASSWORD 'medical_password';`

### "Permission denied" Error
- Grant privileges: `GRANT ALL PRIVILEGES ON DATABASE medical_records TO medical_user;`

### Reset Everything (Docker)
```bash
# Stop and remove containers
docker-compose down

# Remove volumes (this deletes all data!)
docker-compose down -v

# Start fresh
docker-compose up -d postgres
```

## Quick Check Commands

### Check if PostgreSQL is running (Docker)
```bash
docker ps | grep postgres
```

### Check database connection (Docker)
```bash
docker exec -it medical_records_db psql -U medical_user -d medical_records
```

### View database logs (Docker)
```bash
docker logs medical_records_db
```

## Default Credentials (Docker)

- **Host**: localhost
- **Port**: 5432
- **Database**: medical_records
- **User**: medical_user
- **Password**: medical_password

**⚠️ WARNING**: These are default development credentials. Change them in production!

