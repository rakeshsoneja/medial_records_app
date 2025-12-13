# Next Steps After Frontend is Live

Your frontend service is now live! 🎉 Here's what to do next to get your full app working.

## ✅ Current Status

- ✅ Frontend Service: **LIVE** at `https://medical-records-frontend-etnv.onrender.com`
- ⚠️ Backend Service: Needs configuration
- ⚠️ Database: Needs to be created and linked

## 🔧 Required Next Steps

### Step 1: Create and Link Database

**⚠️ IMPORTANT**: If you see "localhost" connection errors, see [FIX_DATABASE_CONNECTION.md](./FIX_DATABASE_CONNECTION.md) for detailed help.

**This is REQUIRED for the backend to work.**

1. **Create PostgreSQL Database**:
   - Go to [Render Dashboard](https://dashboard.render.com)
   - Click "New +" → "PostgreSQL"
   - Name: `medical-records-db`
   - Region: Same as your backend (e.g., Oregon)
   - Plan: Free (or paid)
   - Click "Create Database"
   - **Save the connection details**

2. **Link Database to Backend**:
   - Go to `medical-records-backend` service
   - Click "Environment" tab
   - Click "Link Database" or "Add from Database"
   - Select your `medical-records-db` database
   - Render will automatically add:
     - `DB_HOST`
     - `DB_PORT`
     - `DB_USER`
     - `DB_PASSWORD`
     - `DB_NAME`

3. **Add Additional Environment Variables**:
   - `DB_SSLMODE` = `require` (add this manually)
   - `JWT_SECRET` = (should be auto-generated, or generate a strong random string)
   - `SERVER_PORT` = Leave empty (Render uses `PORT` automatically)

4. **Save Changes** - Backend will automatically redeploy

### Step 2: Update Frontend API URL

**This is REQUIRED for the frontend to connect to the backend.**

1. Go to `medical-records-frontend` service
2. Click "Environment" tab
3. Find or add `REACT_APP_API_URL`
4. Set it to your backend URL:
   ```
   https://medical-records-backend.onrender.com/api/v1
   ```
   (Replace with your actual backend URL if different)
5. **Save Changes** - Frontend will automatically rebuild

### Step 3: Wait for Backend to Deploy

1. Go to `medical-records-backend` service
2. Check "Events" or "Logs" tab
3. Wait for status to show "Live" (green)
4. This usually takes 2-5 minutes

### Step 4: Create Default User

Once backend is live:

1. Go to `medical-records-backend` service
2. Click "Shell" tab
3. Run:
   ```bash
   cd backend
   go run cmd/seed/main.go
   ```
4. This creates default login:
   - **Email**: `admin@medicalrecords.com`
   - **Password**: `admin123`

### Step 5: Test Your App

1. **Open Frontend**: `https://medical-records-frontend-etnv.onrender.com`
2. **Try to Login**: Use the default credentials above
3. **Check Backend Health**: `https://medical-records-backend.onrender.com/health`
4. **View API Docs**: `https://medical-records-backend.onrender.com/swagger/index.html`

## 📋 Checklist

- [ ] Database created (`medical-records-db`)
- [ ] Database linked to backend service
- [ ] `DB_SSLMODE=require` added to backend
- [ ] Backend service shows "Live" status
- [ ] `REACT_APP_API_URL` set in frontend
- [ ] Frontend rebuilt successfully
- [ ] Default user created (via seed script)
- [ ] Can access frontend URL
- [ ] Can login with default credentials
- [ ] Backend API responds correctly

## 🔍 Verify Everything Works

### Test Backend:
```bash
# Health check
curl https://medical-records-backend.onrender.com/health

# Should return: {"status":"ok"}
```

### Test Frontend:
1. Open: `https://medical-records-frontend-etnv.onrender.com`
2. Should see login page
3. Login with default credentials
4. Should see dashboard

### Test API:
1. Open: `https://medical-records-backend.onrender.com/swagger/index.html`
2. Should see Swagger API documentation

## ⚠️ Common Issues

### Frontend Shows Errors
- **Check**: `REACT_APP_API_URL` is set correctly
- **Check**: Backend is live and accessible
- **Check**: CORS is configured (add `FRONTEND_URL` to backend env)

### Can't Login
- **Check**: Default user was created (run seed script)
- **Check**: Backend is running
- **Check**: Database is connected

### Backend Not Starting
- **Check**: Database is linked
- **Check**: All environment variables are set
- **Check**: Backend logs for errors

## 🎉 You're Almost There!

Once you complete these steps, your Medical Records App will be fully functional and publicly accessible!

## 📞 Need Help?

- See [ACCESS_APP.md](./ACCESS_APP.md) for detailed access instructions
- See [RENDER_DATABASE_SETUP.md](./RENDER_DATABASE_SETUP.md) for database setup
- Check service logs in Render Dashboard for errors

