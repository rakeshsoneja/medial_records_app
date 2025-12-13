# Deployment Guide for Render

This directory contains all deployment configuration files for deploying the Medical Records App to Render.

## 📁 Directory Structure

```
deployment/
├── README.md              # This file - deployment instructions
├── render.yaml            # Render service configuration
├── backend-env.template   # Backend environment variables template
├── frontend-env.template  # Frontend environment variables template
├── deploy-backend.sh      # Backend deployment script
└── deploy-frontend.sh     # Frontend deployment script
```

## 🚀 Quick Start

### Option 1: Using render.yaml (Recommended)

**⚠️ IMPORTANT: Create Database First!**

Render Blueprints do not support creating databases. You must create the PostgreSQL database manually before deploying services.

1. **Create PostgreSQL Database** (See [RENDER_DATABASE_SETUP.md](./RENDER_DATABASE_SETUP.md))
   - Go to Render Dashboard → "New +" → "PostgreSQL"
   - Create database: `medical-records-db`
   - Save connection details

2. **Deploy Services with Blueprint**
   - Go to https://dashboard.render.com
   - Click "New +" → "Blueprint"
   - Connect your GitHub repository
   - Select `render.yaml` (or `deployment/render.yaml`) as the blueprint file
   - Click "Apply"

3. **Configure Database Connection**
   - After backend service is created, go to `medical-records-backend` service
   - Click "Environment" tab
   - Add database connection variables (see [RENDER_DATABASE_SETUP.md](./RENDER_DATABASE_SETUP.md))
   - Or use "Link Database" feature if in same region

4. **Set Other Environment Variables**
   - Add any missing environment variables from the templates

### Option 2: Manual Service Setup

Follow the step-by-step guide below.

## 📋 Step-by-Step Deployment

### Step 1: Create PostgreSQL Database

1. Go to Render Dashboard → "New +" → "PostgreSQL"
2. Configure:
   - **Name**: `medical-records-db`
   - **Database**: `medical_records`
   - **User**: `medical_user`
   - **Region**: Choose closest to you
   - **Plan**: Free (or paid for production)
3. Click "Create Database"
4. **Save the connection details** (you'll need them)

### Step 2: Deploy Backend Service

1. Go to Render Dashboard → "New +" → "Web Service"
2. Connect your GitHub repository
3. Configure:
   - **Name**: `medical-records-backend`
   - **Region**: Same as database
   - **Branch**: `main` (or your default branch)
   - **Root Directory**: Leave empty (or `backend` if deploying separately)
   - **Environment**: `Go`
   - **Build Command**: 
     ```bash
     cd backend && go mod download && go build -o ../bin/server ./cmd/server
     ```
   - **Start Command**: 
     ```bash
     ./bin/server
     ```
4. **Add Environment Variables** (see `backend-env.template`)
5. Click "Create Web Service"

### Step 3: Deploy Frontend Service

1. Go to Render Dashboard → "New +" → "Static Site" (or "Web Service")
2. Connect your GitHub repository
3. Configure:
   - **Name**: `medical-records-frontend`
   - **Region**: Same as backend
   - **Branch**: `main`
   - **Root Directory**: `frontend`
   - **Build Command**: 
     ```bash
     npm install && npm run build
     ```
   - **Publish Directory**: `build`
4. **Add Environment Variables**:
   - `REACT_APP_API_URL`: Your backend URL (e.g., `https://medical-records-backend.onrender.com/api/v1`)
5. Click "Create Static Site"

## 🔐 Environment Variables

### Backend Environment Variables

See `backend-env.template` for all required variables.

**Required:**
- Database connection (auto-filled from PostgreSQL service)
- `JWT_SECRET` (generate a strong random string)
- `SERVER_PORT` (usually 8080)

**Optional:**
- AWS S3 credentials (for file uploads)
- SMTP settings (for email)
- Twilio credentials (for SMS)

### Frontend Environment Variables

See `frontend-env.template` for all required variables.

**Required:**
- `REACT_APP_API_URL`: Your backend service URL

## 🌐 Accessing Your Deployed App

**See [ACCESS_APP.md](./ACCESS_APP.md) for complete instructions on accessing your app after deployment.**

Quick steps:
1. Get your service URLs from Render Dashboard
2. Configure database connection (see [RENDER_DATABASE_SETUP.md](./RENDER_DATABASE_SETUP.md))
3. Update frontend API URL
4. Create default user (run seed script)
5. Access your app at the frontend URL

## 🔄 Post-Deployment Steps

### 1. Run Database Migrations

After backend is deployed, migrations run automatically on startup. If they don't:

1. Go to backend service → "Shell"
2. Run:
   ```bash
   cd backend
   go run cmd/server/main.go
   ```
   (Migrations run automatically on server start)

### 2. Create Default User (Optional)

1. Go to backend service → "Shell"
2. Run:
   ```bash
   cd backend
   go run cmd/seed/main.go
   ```

### 3. Verify Deployment

- Backend Health: `https://your-backend.onrender.com/health`
- Backend API Docs: `https://your-backend.onrender.com/swagger/index.html`
- Frontend: `https://your-frontend.onrender.com`

## 🛠️ Troubleshooting

### Backend Issues

**"Failed to connect to database"**
- Check database is running
- Verify environment variables are set correctly
- Check database connection string

**"Port already in use"**
- Render sets PORT automatically, use `os.Getenv("PORT")` in code
- Update backend to use `PORT` environment variable

**"Build failed"**
- Check Go version (Render uses Go 1.21+)
- Verify `go.mod` is correct
- Check build logs for errors

### Frontend Issues

**"Cannot connect to backend"**
- Verify `REACT_APP_API_URL` is set correctly
- Check CORS settings in backend
- Ensure backend URL includes `/api/v1`

**"Build failed"**
- Check Node.js version
- Verify `package.json` is correct
- Check build logs for errors

## 📝 Notes

- **Free tier limitations**: Services spin down after 15 minutes of inactivity
- **First request**: May take 30-60 seconds to wake up
- **Database**: Free tier has connection limits
- **Custom domains**: Available on paid plans

## 🔗 Useful Links

- Render Dashboard: https://dashboard.render.com
- Render Docs: https://render.com/docs
- Render Status: https://status.render.com

