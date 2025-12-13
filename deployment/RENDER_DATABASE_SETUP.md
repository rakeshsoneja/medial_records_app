# Render Database Setup Guide

## Important: Database Must Be Created Manually

Render Blueprints **do not support creating databases directly**. You must create the PostgreSQL database manually in the Render Dashboard before deploying your services.

## Step-by-Step Database Setup

### 1. Create PostgreSQL Database

1. Go to [Render Dashboard](https://dashboard.render.com)
2. Click **"New +"** → **"PostgreSQL"**
3. Configure:
   - **Name**: `medical-records-db`
   - **Database**: `medical_records`
   - **User**: `medical_user` (or leave default)
   - **Region**: Choose closest to you (e.g., Oregon)
   - **Plan**: Free (or paid for production)
4. Click **"Create Database"**
5. **Save the connection details** - you'll need them in the next step

### 2. Get Database Connection Details

After creating the database, go to the database service page and find:
- **Internal Database URL** (for services in the same region)
- **External Database URL** (for external connections)
- **Host**
- **Port**
- **Database Name**
- **User**
- **Password**

### 3. Add Database Environment Variables to Backend Service

After deploying the backend service using the Blueprint:

1. Go to your `medical-records-backend` service in Render Dashboard
2. Click on **"Environment"** tab
3. Add these environment variables:

```
DB_HOST=<your-database-host>
DB_PORT=<your-database-port>
DB_USER=<your-database-user>
DB_PASSWORD=<your-database-password>
DB_NAME=medical_records
DB_SSLMODE=require
```

**Important**: Use the **Internal Database URL** values if your backend is in the same region, or **External Database URL** if in a different region.

### 4. Alternative: Use Render's Database Connection

If your backend service is in the same region as the database, Render automatically provides connection details. You can:

1. Go to your backend service → **"Environment"** tab
2. Click **"Link Database"** or **"Add from Database"**
3. Select your `medical-records-db` database
4. Render will automatically add the connection environment variables

## Quick Reference

After creating the database, your backend service needs these environment variables:

| Variable | Example Value | Source |
|----------|---------------|--------|
| `DB_HOST` | `dpg-xxxxx-a.oregon-postgres.render.com` | From database service |
| `DB_PORT` | `5432` | From database service |
| `DB_USER` | `medical_user` | From database service |
| `DB_PASSWORD` | `xxxxx` | From database service |
| `DB_NAME` | `medical_records` | From database service |
| `DB_SSLMODE` | `require` | Always use `require` for Render |

## Troubleshooting

### "Database connection failed"

- Verify all environment variables are set correctly
- Check that the database service is running
- Ensure `DB_SSLMODE=require` is set
- Verify the backend service is in the same region as the database (for internal connections)

### "Password authentication failed"

- Double-check the `DB_PASSWORD` value
- Ensure `DB_USER` matches the database user
- Try regenerating the database password in Render Dashboard

