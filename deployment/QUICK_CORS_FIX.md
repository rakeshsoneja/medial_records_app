# Quick CORS Fix for Register/Login

If you're getting CORS errors when calling `/api/v1/auth/register` from the frontend, follow these steps:

## 🔧 Immediate Fix: Add FRONTEND_URL Environment Variable

### Step 1: Get Your Frontend URL

Your frontend URL is: `https://medical-records-frontend-etnv.onrender.com`

### Step 2: Add FRONTEND_URL to Backend

1. Go to [Render Dashboard](https://dashboard.render.com)
2. Click on `medical-records-backend` service
3. Click **"Environment"** tab
4. Click **"Add Environment Variable"**
5. Add:
   - **Key**: `FRONTEND_URL`
   - **Value**: `https://medical-records-frontend-etnv.onrender.com`
   - **Important**: No trailing slash!
6. Click **"Save Changes"**

### Step 3: Verify APP_ENV is Set

1. Still in **"Environment"** tab
2. Check if `APP_ENV` exists
3. If not, add it:
   - **Key**: `APP_ENV`
   - **Value**: `production`
4. Click **"Save Changes"**

### Step 4: Wait for Redeploy

- Backend will automatically redeploy (2-5 minutes)
- Check **"Events"** or **"Logs"** tab
- Wait for "Live" status

### Step 5: Test Again

1. Go to your frontend: `https://medical-records-frontend-etnv.onrender.com`
2. Try to register
3. CORS error should be gone

## 🔍 Verify CORS is Working

After backend redeploys, test the CORS headers:

1. Open browser DevTools (F12)
2. Go to **"Network"** tab
3. Try to register
4. Click on the `/auth/register` request
5. Check **"Response Headers"**:
   - Should see: `Access-Control-Allow-Origin: https://medical-records-frontend-etnv.onrender.com`
   - Should see: `Access-Control-Allow-Credentials: true`

## 🐛 If Still Getting CORS Errors

### Check 1: Verify Environment Variables

1. Go to backend → **"Environment"** tab
2. Verify:
   - `FRONTEND_URL` = `https://medical-records-frontend-etnv.onrender.com` (no trailing slash)
   - `APP_ENV` = `production`
3. Refresh the page to make sure they're saved

### Check 2: Check Backend Logs

1. Go to backend → **"Logs"** tab
2. Look for CORS-related errors
3. Check if backend is actually running

### Check 3: Verify Backend is Live

1. Test backend health: `https://medical-records-backend.onrender.com/health`
2. Should return: `{"status":"ok"}`
3. If not, backend might not be running

### Check 4: Clear Browser Cache

1. Clear browser cache (Ctrl+Shift+Delete)
2. Or try in incognito/private mode
3. Try registering again

## 📋 Required Environment Variables

Your backend should have:

- [ ] `FRONTEND_URL` = `https://medical-records-frontend-etnv.onrender.com`
- [ ] `APP_ENV` = `production`
- [ ] `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_SSLMODE`
- [ ] `JWT_SECRET`
- [ ] Other required variables

## 💡 Alternative: If FRONTEND_URL Doesn't Work

The code has a fallback that allows all origins in production if `FRONTEND_URL` is not set. But this requires:

1. `APP_ENV` = `production` (must be set)
2. `FRONTEND_URL` = (empty or not set)

However, **it's better to set FRONTEND_URL explicitly** for security.

## 🎯 Quick Test Command

You can test CORS with curl:

```bash
curl -X OPTIONS https://medical-records-backend.onrender.com/api/v1/auth/register \
  -H "Origin: https://medical-records-frontend-etnv.onrender.com" \
  -H "Access-Control-Request-Method: POST" \
  -H "Access-Control-Request-Headers: Content-Type" \
  -v
```

You should see `Access-Control-Allow-Origin` in the response headers.

## ✅ After Fix

Once CORS is working:
- Registration should work
- Login should work
- All API calls from frontend should work

