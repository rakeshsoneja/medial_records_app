# Medical Records Management & Sharing App

A secure, HIPAA-compliant application for managing and sharing medical records, prescriptions, appointments, and lab reports.

## Features

- **User Authentication**: Secure sign-up/login with email/password, phone OTP, or OAuth
- **Medical Records Management**: 
  - Prescriptions with photos/PDFs
  - Appointments with reminders
  - Lab reports with tagging
- **Health Insurance Details**: Store and manage insurance information
- **Medication Management**: Track regular medicines by pharmacy with refill reminders
- **Health Check-up Reminders**: Automated reminders for scheduled check-ups
- **Secure Sharing**: Time-limited shareable links via SMS/Email
- **Dashboard**: Summary view of all records with search and filter
- **Security**: End-to-end encryption, audit logs, HIPAA-compliant practices

## Tech Stack

- **Frontend**: React.js with TypeScript
- **Backend**: Go (Golang)
- **Database**: PostgreSQL
- **Storage**: AWS S3 (or compatible) for file storage
- **Authentication**: JWT with optional OAuth support
- **API Documentation**: Swagger/OpenAPI

## Project Structure

```
DigitalRecordsProject/
├── backend/          # Go backend application
├── frontend/         # React frontend application
├── docker-compose.yml # Docker setup for local development
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.21+
- Node.js 18+
- PostgreSQL 14+
- Docker (optional, for containerized setup)

### Backend Setup

```bash
cd backend
go mod download
cp .env.example .env  # Configure your environment variables
go run cmd/server/main.go
```

### Frontend Setup

```bash
cd frontend
npm install
npm start
```

### Database Setup

```bash
# Using Docker
docker-compose up -d postgres

# Run migrations
cd backend
go run cmd/migrate/main.go
```

## Environment Variables

See `.env.example` files in backend and frontend directories for required configuration.

## API Documentation

Once the server is running, visit `http://localhost:8080/swagger/index.html` for interactive API documentation.

## Default User (For Testing)

**No default user exists by default.** To create a test user:

```bash
cd backend
go run cmd/seed/main.go
```

This creates:
- **Email**: `admin@medicalrecords.com`
- **Password**: `admin123`

⚠️ **For production**: Register users through the frontend registration page.

## Security

- All data is encrypted at rest and in transit
- Time-limited sharing links with auto-expiration
- Audit logs for all shared access
- HIPAA-compliant security practices

## License

MIT

