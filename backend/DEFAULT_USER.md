# Default User Credentials

## ⚠️ Important: No Default User Exists

**The application does NOT have a default user.** You need to either:

1. **Register a new user** through the frontend (recommended)
2. **Create a default user** using the seed script (for testing)

## Option 1: Register Through Frontend (Recommended)

1. Start the frontend: `cd frontend && npm start`
2. Go to `http://localhost:3000`
3. Click "Register" or go to `/register`
4. Fill in the registration form:
   - Email: your-email@example.com
   - Password: (minimum 8 characters)
   - First Name: Your First Name
   - Last Name: Your Last Name
   - Phone: (optional)

## Option 2: Create Default User (For Testing)

### Step 1: Run the Seed Script

```bash
cd backend
go run cmd/seed/main.go
```

This will create a default user with:
- **Email**: `admin@medicalrecords.com`
- **Password**: `admin123`

### Step 2: Login

Use these credentials to login:
- **Email**: `admin@medicalrecords.com`
- **Password**: `admin123`

### Customize Default User

You can customize the default user by setting environment variables before running the seed script:

```bash
# Windows PowerShell
$env:DEFAULT_USER_EMAIL="your-email@example.com"
$env:DEFAULT_USER_PASSWORD="your-password"
$env:DEFAULT_USER_FIRSTNAME="Your"
$env:DEFAULT_USER_LASTNAME="Name"
go run cmd/seed/main.go

# Linux/Mac
export DEFAULT_USER_EMAIL="your-email@example.com"
export DEFAULT_USER_PASSWORD="your-password"
export DEFAULT_USER_FIRSTNAME="Your"
export DEFAULT_USER_LASTNAME="Name"
go run cmd/seed/main.go
```

Or add to your `.env` file:
```env
DEFAULT_USER_EMAIL=admin@medicalrecords.com
DEFAULT_USER_PASSWORD=admin123
DEFAULT_USER_FIRSTNAME=Admin
DEFAULT_USER_LASTNAME=User
```

## Security Note

⚠️ **IMPORTANT**: The default user is only for development/testing. In production:
- Remove or disable the seed script
- Use strong passwords
- Require email verification
- Use proper user registration flow

## Reset Default User

If you want to recreate the default user:

1. Delete the user from the database, or
2. Change the email in the seed script and run it again

