# How to Access Your Deployed App on Render

After your Blueprint deployment completes, follow these steps to access and run your app publicly.

## 🚀 Quick Access Steps

### Step 1: Get Your Service URLs

1. Go to [Render Dashboard](https://dashboard.render.com)
2. You'll see your services:
   - `medical-records-backend` - Backend API
   - `medical-records-frontend` - Frontend App

3. **Click on each service** to see its public URL:
   - Backend URL: `https://medical-records-backend.onrender.com`
   - Frontend URL: `https://medical-records-frontend.onrender.com`

### Step 2: Configure Database Connection (Required!)

**⚠️ IMPORTANT**: The backend won't work until you configure the database connection.

1. **Create PostgreSQL Database** (if not done yet):
   - Go to Render Dashboard → "New +" → "PostgreSQL"
   - Name: `medical-records-db`
   - Click "Create Database"
   - **Save the connection details**

2. **Link Database to Backend**:
   - Go to `medical-records-backend` service
   - Click "Environment" tab
   - Click "Link Database" or "Add from Database"
   - Select your `medical-records-db` database
   - Render will auto-add: `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`

3. **Add Missing Environment Variables**:
   - `DB_SSLMODE` = `require` (should already be set)
   - `JWT_SECRET` = (should be auto-generated, or set a strong random string)
   - `SERVER_PORT` = `8080` (or leave empty to use Render's PORT)

4. **Save Changes** - The backend will automatically redeploy

### Step 3: Update Frontend API URL

1. Go to `medical-records-frontend` service
2. Click "Environment" tab
3. Update `REACT_APP_API_URL`:
   - Get your backend URL from the backend service dashboard
   - Format: `https://your-backend-url.onrender.com/api/v1`
   - Example: `https://medical-records-backend.onrender.com/api/v1`
4. **Save Changes** - The frontend will automatically rebuild

### Step 4: Wait for Services to Deploy

- Backend: Usually takes 2-5 minutes
- Frontend: Usually takes 3-7 minutes (includes build time)

Watch the "Events" or "Logs" tab to see deployment progress.

### Step 5: Access Your App

1. **Open your frontend URL** in a browser:
   ```
   https://medical-records-frontend.onrender.com
   ```

2. **Test the backend API**:
   - Health check: `https://medical-records-backend.onrender.com/health`
   - API docs: `https://medical-records-backend.onrender.com/swagger/index.html`

## 🔐 Default Login Credentials

After deployment, you need to create a user or seed the default user:

### Option 1: Use Seed Script (Recommended)

1. Go to `medical-records-backend` service
2. Click "Shell" tab
3. Run:
   ```bash
   cd backend
   go run cmd/seed/main.go
   ```
4. This creates a default user:
   - **Email**: `admin@medicalrecords.com`
   - **Password**: `admin123`

### Option 2: Register New User

1. Go to your frontend URL
2. Click "Sign Up" or "Register"
3. Create a new account

## 📋 Post-Deployment Checklist

- [ ] Database created and linked to backend
- [ ] Backend environment variables configured
- [ ] Frontend `REACT_APP_API_URL` updated
- [ ] Backend service is "Live" (green status)
- [ ] Frontend service is "Live" (green status)
- [ ] Default user created (via seed script)
- [ ] Can access frontend URL
- [ ] Can login with default credentials
- [ ] Backend API responds at `/health` endpoint

## 🌐 Public URLs

Once deployed, your app will be accessible at:

- **Frontend**: `https://medical-records-frontend.onrender.com`
- **Backend API**: `https://medical-records-backend.onrender.com/api/v1`
- **API Health**: `https://medical-records-backend.onrender.com/health`
- **API Docs**: `https://medical-records-backend.onrender.com/swagger/index.html`

## ⚠️ Important Notes

### Free Tier Limitations

1. **Spin Down**: Services spin down after 15 minutes of inactivity
2. **First Request**: May take 30-60 seconds to wake up
3. **Cold Starts**: Subsequent requests are faster

### Database Connection

- Use **Internal Database URL** if backend and database are in the same region
- Use **External Database URL** if they're in different regions
- Always set `DB_SSLMODE=require` for Render databases

### CORS Configuration

If you get CORS errors:
1. Go to backend service → "Environment" tab
2. Add: `FRONTEND_URL` = `https://medical-records-frontend.onrender.com`
3. Save and redeploy

## 🐛 Troubleshooting

### Backend Not Starting

**Check:**
- Database connection variables are set
- `DB_SSLMODE=require` is set
- Backend logs for errors (click "Logs" tab)

**Common Issues:**
- Missing database connection → Add database environment variables
- Wrong database credentials → Verify in database service
- Port conflict → Remove `SERVER_PORT` or set to empty (Render uses `PORT`)

### Frontend Not Loading

**Check:**
- `REACT_APP_API_URL` is set correctly
- Frontend build completed successfully
- Backend is accessible

**Common Issues:**
- Wrong API URL → Update `REACT_APP_API_URL` with correct backend URL
- CORS errors → Add `FRONTEND_URL` to backend environment
- Build failed → Check build logs for errors

### Database Connection Failed

**Check:**
- Database service is running
- All database environment variables are set
- `DB_SSLMODE=require` is set
- Using correct database URL (internal vs external)

### Can't Login

**Check:**
- Default user was created (run seed script)
- Backend is running and accessible
- Frontend API URL is correct

## 🔄 Updating Your App

After making code changes:

1. **Commit and push to GitHub**:
   ```bash
   git add .
   git commit -m "Your changes"
   git push origin main
   ```

2. **Render automatically redeploys** when you push to the connected branch

3. **Monitor deployment** in Render Dashboard → "Events" tab

## 📞 Need Help?

- **Render Dashboard**: https://dashboard.render.com
- **Render Docs**: https://render.com/docs
- **Service Logs**: Click on service → "Logs" tab
- **Service Events**: Click on service → "Events" tab

