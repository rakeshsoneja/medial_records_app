# Setup Guide

## Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher
- PostgreSQL 14 or higher
- Docker (optional, for containerized setup)

## Quick Start

### 1. Database Setup

Using Docker:
```bash
docker-compose up -d postgres
```

Or install PostgreSQL locally and create a database:
```sql
CREATE DATABASE medical_records;
CREATE USER medical_user WITH PASSWORD 'medical_password';
GRANT ALL PRIVILEGES ON DATABASE medical_records TO medical_user;
```

### 2. Backend Setup

```bash
cd backend

# Copy environment file
cp .env.example .env

# Edit .env with your configuration
# Update database credentials, JWT secret, AWS S3 credentials, etc.

# Install dependencies
go mod download

# Run migrations (automatic on startup)
# Or manually: go run cmd/migrate/main.go

# Start server
go run cmd/server/main.go
```

The backend will start on `http://localhost:8080`

### 3. Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Copy environment file
cp .env.example .env

# Start development server
npm start
```

The frontend will start on `http://localhost:3000`

### 4. Access the Application

- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- Swagger Documentation: http://localhost:8080/swagger/index.html

## Environment Variables

### Backend (.env)

Required:
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`
- `JWT_SECRET` (use a strong secret in production)
- `SERVER_PORT` (default: 8080)

Optional (for file uploads):
- AWS S3 credentials for storing prescription PDFs and lab reports

Optional (for email/SMS sharing):
- SMTP configuration for email
- Twilio credentials for SMS

### Frontend (.env)

- `REACT_APP_API_URL` (default: http://localhost:8080/api/v1)

## Features

### Implemented

✅ User Authentication (Email/Password)
✅ Medical Records Management:
   - Prescriptions
   - Appointments
   - Lab Reports
   - Health Insurance
✅ Medication Tracking with Pharmacy Info
✅ Health Check-up Reminders
✅ Secure Sharing with Time-limited Links
✅ Dashboard with Summary View
✅ API Documentation (Swagger)

### Security Features

- JWT-based authentication
- Password hashing with bcrypt
- Time-limited share links
- Access count limits
- Audit logs for shared record access
- User-specific data isolation

### Future Enhancements

- Phone OTP authentication
- OAuth integration (Google, Facebook)
- File upload to S3
- Email/SMS notifications
- OCR for prescription extraction
- Doctor portal
- Multi-device sync
- Vehicle Insurance
- Investments tracking

## API Endpoints

### Authentication
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - Login user
- `GET /api/v1/auth/profile` - Get user profile

### Prescriptions
- `POST /api/v1/prescriptions` - Create prescription
- `GET /api/v1/prescriptions` - List prescriptions
- `GET /api/v1/prescriptions/:id` - Get prescription
- `PUT /api/v1/prescriptions/:id` - Update prescription
- `DELETE /api/v1/prescriptions/:id` - Delete prescription

### Appointments
- `POST /api/v1/appointments` - Create appointment
- `GET /api/v1/appointments` - List appointments

### Lab Reports
- `POST /api/v1/lab-reports` - Create lab report
- `GET /api/v1/lab-reports` - List lab reports

### Medications
- `POST /api/v1/medications` - Create medication
- `GET /api/v1/medications` - List medications
- `GET /api/v1/medications/refill-needed` - Get medications needing refill

### Reminders
- `POST /api/v1/reminders` - Create reminder
- `GET /api/v1/reminders` - List reminders
- `GET /api/v1/reminders/upcoming` - Get upcoming reminders

### Insurance
- `POST /api/v1/insurance` - Create insurance record
- `GET /api/v1/insurance` - List insurance records

### Sharing
- `POST /api/v1/sharing/create` - Create share link
- `GET /api/v1/sharing/my-shares` - Get my shared records
- `POST /api/v1/sharing/:id/revoke` - Revoke share link
- `GET /api/v1/share/:token` - Access shared record (public)

### Dashboard
- `GET /api/v1/dashboard` - Get dashboard summary

## Development

### Running Tests
```bash
cd backend
go test ./...
```

### Generating Swagger Docs
```bash
cd backend
swag init -g cmd/server/main.go -o ./docs
```

### Building for Production

Backend:
```bash
cd backend
go build -o bin/server cmd/server/main.go
```

Frontend:
```bash
cd frontend
npm run build
```

## Troubleshooting

### Database Connection Issues
- Ensure PostgreSQL is running
- Check database credentials in `.env`
- Verify database exists and user has permissions

### Port Already in Use
- Change `SERVER_PORT` in backend `.env`
- Update `REACT_APP_API_URL` in frontend `.env`

### CORS Issues
- Update CORS configuration in `backend/internal/router/router.go`
- Add your frontend URL to allowed origins

## Architecture

### Backend Structure
```
backend/
├── cmd/
│   └── server/        # Application entry point
├── internal/
│   ├── auth/          # Authentication utilities
│   ├── config/        # Configuration management
│   ├── database/      # Database models and migrations
│   ├── handlers/      # HTTP request handlers
│   ├── middleware/    # HTTP middleware
│   ├── router/        # Route definitions
│   ├── services/      # Business logic
│   └── utils/         # Utility functions
└── go.mod             # Go dependencies
```

### Frontend Structure
```
frontend/
├── public/            # Static files
├── src/
│   ├── components/   # Reusable components
│   ├── contexts/     # React contexts
│   ├── pages/        # Page components
│   ├── services/     # API services
│   └── App.js        # Main app component
└── package.json      # Node dependencies
```

## License

MIT

