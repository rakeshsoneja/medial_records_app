# Step-by-Step Render Deployment Guide

## Prerequisites

- ✅ GitHub repository with your code pushed
- ✅ Render account (sign up at https://render.com)
- ✅ GitHub account connected to Render

## Step 1: Create PostgreSQL Database

1. **Go to Render Dashboard**
   - Visit https://dashboard.render.com
   - Sign in or create account

2. **Create PostgreSQL Database**
   - Click "New +" button (top right)
   - Select "PostgreSQL"
   - Fill in:
     - **Name**: `medical-records-db`
     - **Database**: `medical_records`
     - **User**: `medical_user`
     - **Region**: Choose closest to you (e.g., `Oregon (US West)`)
     - **PostgreSQL Version**: Latest (15)
     - **Plan**: Free (for testing) or Starter ($7/month for production)
   - Click "Create Database"
   - **Wait for database to be created** (takes 1-2 minutes)
   - **Copy the Internal Database URL** (you'll need it)

## Step 2: Deploy Backend Service

1. **Create Web Service**
   - Click "New +" → "Web Service"
   - Connect your GitHub account (if not already connected)
   - Select your repository: `rakeshsoneja/medial_records_app`

2. **Configure Backend Service**
   - **Name**: `medical-records-backend`
   - **Region**: Same as database
   - **Branch**: `main` (or your default branch)
   - **Root Directory**: Leave empty (or `backend` if you want)
   - **Environment**: `Go`
   - **Build Command**:
     ```bash
     cd backend && go mod download && go build -o ../bin/server ./cmd/server
     ```
   - **Start Command**:
     ```bash
     ./bin/server
     ```

3. **Add Environment Variables**
   Click "Advanced" → "Add Environment Variable" and add:

   **Database Variables** (from your PostgreSQL service):
   - `DB_HOST` = (from PostgreSQL service - Internal Database URL host)
   - `DB_PORT` = `5432`
   - `DB_USER` = `medical_user`
   - `DB_PASSWORD` = (from PostgreSQL service)
   - `DB_NAME` = `medical_records`
   - `DB_SSLMODE` = `require`

   **Application Variables**:
   - `SERVER_PORT` = `8080`
   - `PORT` = `8080` (Render uses this, but we set SERVER_PORT as fallback)
   - `JWT_SECRET` = (Generate a strong random string - use: `openssl rand -base64 32`)
   - `JWT_EXPIRATION_HOURS` = `24`
   - `APP_ENV` = `production`
   - `APP_DEBUG` = `false`

   **Optional Variables** (if you have them):
   - AWS S3 credentials (for file uploads)
   - SMTP settings (for email)
   - Twilio credentials (for SMS)

4. **Create Service**
   - Click "Create Web Service"
   - Wait for build and deployment (5-10 minutes)
   - **Note the service URL** (e.g., `https://medical-records-backend.onrender.com`)

## Step 3: Deploy Frontend Service

1. **Create Static Site** (or Web Service)
   - Click "New +" → "Static Site"
   - Connect your GitHub account (if not already)
   - Select your repository: `rakeshsoneja/medial_records_app`

2. **Configure Frontend Service**
   - **Name**: `medical-records-frontend`
   - **Region**: Same as backend
   - **Branch**: `main`
   - **Root Directory**: `frontend`
   - **Build Command**:
     ```bash
     npm install && npm run build
     ```
   - **Publish Directory**: `build`

3. **Add Environment Variables**
   - `REACT_APP_API_URL` = `https://medical-records-backend.onrender.com/api/v1`
     (Replace with your actual backend URL)

4. **Create Service**
   - Click "Create Static Site"
   - Wait for build and deployment (3-5 minutes)
   - **Note the frontend URL**

## Step 4: Post-Deployment Configuration

### Update Frontend API URL

After backend is deployed, update frontend environment variable:
1. Go to frontend service → "Environment"
2. Update `REACT_APP_API_URL` with your actual backend URL
3. Click "Save Changes"
4. Service will automatically rebuild

### Update Backend CORS

The backend needs to allow your frontend URL. Update CORS in `backend/internal/router/router.go`:

```go
r.Use(cors.New(cors.Config{
    AllowOrigins: []string{
        "http://localhost:3000",
        "http://localhost:3001",
        "https://medical-records-frontend.onrender.com", // Add your frontend URL
    },
    // ... rest of config
}))
```

Commit and push this change.

## Step 5: Verify Deployment

1. **Check Backend**
   - Health: `https://your-backend.onrender.com/health`
   - API Docs: `https://your-backend.onrender.com/swagger/index.html`

2. **Check Frontend**
   - Visit: `https://your-frontend.onrender.com`
   - Try logging in or registering

3. **Test Database Connection**
   - Backend logs should show successful database connection
   - Try creating a record through the frontend

## Step 6: Create Default User (Optional)

1. Go to backend service → "Shell" tab
2. Run:
   ```bash
   cd backend
   go run cmd/seed/main.go
   ```
3. This creates:
   - Email: `admin@medicalrecords.com`
   - Password: `admin123`

## 🔧 Troubleshooting

### Backend Won't Start

**Check logs:**
- Go to backend service → "Logs"
- Look for error messages

**Common issues:**
- Database connection failed → Check DB environment variables
- Port error → Ensure PORT or SERVER_PORT is set
- Build failed → Check Go version and dependencies

### Frontend Can't Connect to Backend

**Check:**
1. Backend URL is correct in `REACT_APP_API_URL`
2. CORS is configured in backend
3. Backend service is running (not sleeping)

### Database Connection Issues

**Check:**
1. Database is running
2. Environment variables match database credentials
3. `DB_SSLMODE=require` is set (Render requires SSL)

### Services Keep Sleeping (Free Tier)

**Solution:**
- Free tier services sleep after 15 minutes of inactivity
- First request after sleep takes 30-60 seconds
- Consider upgrading to paid plan for always-on services

## 📊 Monitoring

- **Logs**: View real-time logs in each service dashboard
- **Metrics**: Monitor CPU, memory, and request metrics
- **Alerts**: Set up email alerts for service failures

## 🔄 Updating Deployment

When you push changes to GitHub:
- Render automatically detects changes
- Rebuilds and redeploys services
- No manual intervention needed

## 💰 Cost Estimation

**Free Tier:**
- PostgreSQL: 90 days free, then $7/month
- Backend: Free (with limitations)
- Frontend: Free (with limitations)

**Paid Tier (Recommended for Production):**
- PostgreSQL Starter: $7/month
- Backend Starter: $7/month
- Frontend: Free (static sites are free)

**Total**: ~$14/month for production setup

## 🔗 Next Steps

1. Set up custom domain (optional)
2. Configure SSL certificates (automatic on Render)
3. Set up monitoring and alerts
4. Configure backups for database
5. Set up CI/CD pipeline (optional)

