# Fix CORS Error in Register/Login

If you're getting a CORS (Cross-Origin Resource Sharing) error when trying to register or login, follow these steps.

## 🔧 Quick Fix: Add FRONTEND_URL to Backend

### Step 1: Get Your Frontend URL

Your frontend URL is: `https://medical-records-frontend-etnv.onrender.com`

### Step 2: Add FRONTEND_URL to Backend Environment

1. Go to [Render Dashboard](https://dashboard.render.com)
2. Click on `medical-records-backend` service
3. Click **"Environment"** tab
4. Click **"Add Environment Variable"**
5. Add:
   - **Key**: `FRONTEND_URL`
   - **Value**: `https://medical-records-frontend-etnv.onrender.com`
6. Click **"Save Changes"**
7. Backend will automatically redeploy (takes 2-5 minutes)

### Step 3: Verify CORS is Working

After backend redeploys:

1. **Test Registration**:
   - Go to your frontend: `https://medical-records-frontend-etnv.onrender.com`
   - Try to register a new user
   - CORS error should be gone

2. **Check Browser Console**:
   - Open browser DevTools (F12)
   - Go to "Console" tab
   - Should see no CORS errors

## 🔍 Verify Backend CORS Configuration

The backend code already supports CORS with the `FRONTEND_URL` environment variable. After adding it, the backend will:

- Allow requests from your frontend URL
- Allow credentials (cookies, auth headers)
- Support all necessary HTTP methods (GET, POST, PUT, DELETE, OPTIONS)

## 🐛 If CORS Error Persists

### Check 1: Verify Environment Variable

1. Go to backend service → "Environment" tab
2. Verify `FRONTEND_URL` is set correctly
3. Make sure there are no extra spaces or quotes
4. Value should be: `https://medical-records-frontend-etnv.onrender.com` (no trailing slash)

### Check 2: Verify Backend is Running

1. Check backend service status (should be "Live")
2. Test backend health: `https://medical-records-backend.onrender.com/health`
3. Should return: `{"status":"ok"}`

### Check 3: Check Browser Console

1. Open browser DevTools (F12)
2. Go to "Network" tab
3. Try to register/login
4. Look for the failed request
5. Check the "Response Headers" for CORS headers:
   - `Access-Control-Allow-Origin` should include your frontend URL
   - `Access-Control-Allow-Credentials: true`

### Check 4: Verify Frontend API URL

1. Go to frontend service → "Environment" tab
2. Verify `REACT_APP_API_URL` is set to:
   ```
   https://medical-records-backend.onrender.com/api/v1
   ```
3. Make sure it matches your actual backend URL

## 📋 Complete Environment Variables Checklist

### Backend Service Should Have:

- [ ] `FRONTEND_URL` = `https://medical-records-frontend-etnv.onrender.com`
- [ ] `DB_HOST` (from database)
- [ ] `DB_PORT` (from database)
- [ ] `DB_USER` (from database)
- [ ] `DB_PASSWORD` (from database)
- [ ] `DB_NAME` (from database)
- [ ] `DB_SSLMODE` = `require`
- [ ] `JWT_SECRET` (auto-generated or set manually)
- [ ] `APP_ENV` = `production`
- [ ] `SERVER_PORT` = `8080` (or leave empty)

### Frontend Service Should Have:

- [ ] `REACT_APP_API_URL` = `https://medical-records-backend.onrender.com/api/v1`

## 🔄 After Fixing

1. **Wait for Backend to Redeploy** (2-5 minutes)
2. **Clear Browser Cache** (Ctrl+Shift+Delete or Cmd+Shift+Delete)
3. **Try Registration Again**
4. **Check Browser Console** for any remaining errors

## 💡 Why This Happens

CORS errors occur when:
- Frontend (on one domain) tries to access backend (on another domain)
- Browser blocks the request for security
- Backend must explicitly allow the frontend domain

By setting `FRONTEND_URL`, the backend knows which frontend domain to allow.

## 🎯 Quick Test

After adding `FRONTEND_URL` and backend redeploys:

```bash
# Test CORS headers
curl -H "Origin: https://medical-records-frontend-etnv.onrender.com" \
     -H "Access-Control-Request-Method: POST" \
     -H "Access-Control-Request-Headers: Content-Type" \
     -X OPTIONS \
     https://medical-records-backend.onrender.com/api/v1/auth/register \
     -v
```

You should see `Access-Control-Allow-Origin` header in the response.

