# Setup with Existing PostgreSQL

## Step 1: Update .env File

Edit the `backend/.env` file with your actual PostgreSQL credentials:

```env
# Database Configuration - Use YOUR existing PostgreSQL credentials
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_postgres_username
DB_PASSWORD=your_postgres_password
DB_NAME=medical_records
DB_SSLMODE=disable

# Server Configuration
SERVER_PORT=8080

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-in-production
JWT_EXPIRATION_HOURS=24

# Application Environment
APP_ENV=development
```

**Important:** Replace:
- `your_postgres_username` with your actual PostgreSQL username
- `your_postgres_password` with your actual PostgreSQL password
- `DB_NAME` can be `medical_records` or any database name you prefer

## Step 2: Create the Database

Connect to your PostgreSQL and create the database:

```bash
# Connect to PostgreSQL (use your actual username)
psql -U your_postgres_username -d postgres
```

Or if you use pgAdmin or another GUI tool, connect and run:

```sql
CREATE DATABASE medical_records;
```

If you want to create a specific user for this app (optional):

```sql
CREATE USER medical_user WITH PASSWORD 'medical_password';
GRANT ALL PRIVILEGES ON DATABASE medical_records TO medical_user;
```

Then update your `.env` to use that user:
```env
DB_USER=medical_user
DB_PASSWORD=medical_password
```

## Step 3: Start the Server

```bash
cd backend
go run cmd/server/main.go
```

The server will automatically:
- Connect to your database
- Create all necessary tables (via GORM migrations)

## Troubleshooting

### "Password authentication failed"
- Double-check your username and password in `.env`
- Make sure the user has permission to access the database

### "Database does not exist"
- Create the database: `CREATE DATABASE medical_records;`

### "Permission denied"
- Grant privileges: `GRANT ALL PRIVILEGES ON DATABASE medical_records TO your_username;`

### Using default postgres user
If you want to use the default `postgres` superuser:
```env
DB_USER=postgres
DB_PASSWORD=your_postgres_password
```

