# CORS Debugging Steps - Step by Step

If you're still getting CORS errors, follow these steps in order:

## Step 1: Verify Environment Variables on Render

### Backend Service:
1. Go to [Render Dashboard](https://dashboard.render.com)
2. Select **Backend Service** (`medical-records-backend`)
3. Click **Environment** tab
4. **VERIFY these are set:**
   ```
   APP_ENV=production
   FRONTEND_URL=https://medical-records-frontend.onrender.com
   ```
   ⚠️ **Replace with your actual frontend URL!**

5. **If missing, add them:**
   - Click "Add Environment Variable"
   - Add `APP_ENV` = `production`
   - Add `FRONTEND_URL` = `https://your-frontend-url.onrender.com`
   - Click "Save Changes"
   - Service will auto-restart (wait 2-5 minutes)

### Frontend Service:
1. Select **Frontend Service** (`medical-records-frontend`)
2. Click **Environment** tab
3. **VERIFY:**
   ```
   REACT_APP_API_URL=https://medical-records-backend.onrender.com/api/v1
   ```
   ⚠️ **Replace with your actual backend URL!**

## Step 2: Check Backend Logs

1. Go to Backend Service → **Logs** tab
2. Look for CORS-related messages:
   ```
   CORS: Request from origin: https://...
   CORS: Handling OPTIONS preflight
   CORS: Production mode with allowed origins: [...]
   ```

3. **If you see:**
   - `CORS: No allowed origins configured` → `FRONTEND_URL` is not set
   - `CORS: Origin ... not in allowed list` → URL mismatch
   - No CORS logs at all → Middleware not running

## Step 3: Test OPTIONS Request Manually

Open terminal/command prompt and run:

```bash
curl -X OPTIONS https://medical-records-backend.onrender.com/api/v1/auth/register \
  -H "Origin: https://medical-records-frontend.onrender.com" \
  -H "Access-Control-Request-Method: POST" \
  -H "Access-Control-Request-Headers: Content-Type" \
  -v
```

**Expected Response:**
- Status: `204 No Content`
- Headers should include:
  ```
  Access-Control-Allow-Origin: https://medical-records-frontend.onrender.com
  Access-Control-Allow-Methods: GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD
  Access-Control-Allow-Headers: Origin, Content-Type, ...
  Access-Control-Allow-Credentials: true
  ```

**If you get:**
- `404 Not Found` → Route doesn't exist
- `405 Method Not Allowed` → OPTIONS not handled
- No `Access-Control-Allow-Origin` header → CORS middleware not working

## Step 4: Check Browser Console

1. Open your frontend: `https://medical-records-frontend.onrender.com`
2. Open **Browser DevTools** (F12)
3. Go to **Console** tab
4. Try to register/login
5. **Look for CORS error message:**
   - `Access to XMLHttpRequest ... has been blocked by CORS policy`
   - `No 'Access-Control-Allow-Origin' header is present`
   - `Preflight request doesn't pass access control check`

## Step 5: Check Network Tab

1. In DevTools, go to **Network** tab
2. Try to register/login again
3. Find the failed request (usually red)
4. Click on it
5. Check **Request Headers:**
   - `Origin: https://medical-records-frontend.onrender.com` (should match your frontend)
6. Check **Response Headers:**
   - `Access-Control-Allow-Origin: https://medical-records-frontend.onrender.com` (should be present)
   - If missing → CORS middleware not working

## Step 6: Verify Code is Deployed

1. Check GitHub repository
2. Verify these files exist:
   - `backend/internal/middleware/cors_simple.go`
   - `backend/internal/router/router.go` (should use `SimpleCORSMiddleware()`)
3. Check Render deployment logs:
   - Backend Service → **Events** tab
   - Look for latest deployment
   - Should show "Deployed successfully"

## Step 7: Common Issues & Fixes

### Issue: "No 'Access-Control-Allow-Origin' header"

**Cause:** CORS middleware not running or not configured

**Fix:**
1. Verify `SimpleCORSMiddleware()` is used in `router.go`
2. Check backend logs for CORS messages
3. Ensure service restarted after code changes

### Issue: "Preflight request doesn't pass"

**Cause:** OPTIONS request not handled correctly

**Fix:**
1. Our middleware handles OPTIONS automatically
2. Check backend logs for "Handling OPTIONS preflight"
3. Test OPTIONS request manually (Step 3)

### Issue: "Origin mismatch"

**Cause:** `FRONTEND_URL` doesn't match actual frontend URL

**Fix:**
1. Check actual frontend URL in browser
2. Update `FRONTEND_URL` in backend environment to match exactly
3. Include `https://` and no trailing slash (or include it consistently)

### Issue: "Credentials not allowed"

**Cause:** Using wildcard `*` with credentials

**Fix:**
1. Our middleware uses explicit origins
2. Verify `FRONTEND_URL` is set correctly
3. Check that `Access-Control-Allow-Credentials: true` is in response

## Step 8: Nuclear Option - Allow All Origins (Temporary)

If nothing else works, temporarily allow all origins to verify CORS is the issue:

1. In backend environment, remove `FRONTEND_URL`
2. The middleware will allow all origins (with logging)
3. If this works, the issue is origin matching
4. Then set `FRONTEND_URL` correctly

**⚠️ WARNING:** Only for debugging! Not secure for production.

## Step 9: Get Help

If still not working, collect this information:

1. **Backend Logs** (last 50 lines with CORS messages)
2. **Browser Console Error** (screenshot)
3. **Network Tab** (failed request details)
4. **Environment Variables** (screenshot of Render env vars - hide passwords)
5. **OPTIONS Test Result** (from Step 3)

## Quick Checklist

- [ ] `APP_ENV=production` set in backend
- [ ] `FRONTEND_URL` set to exact frontend URL (with https://)
- [ ] `REACT_APP_API_URL` set in frontend
- [ ] Backend service restarted after env var changes
- [ ] Code deployed to GitHub and Render
- [ ] Backend logs show CORS messages
- [ ] OPTIONS request returns 204 with CORS headers
- [ ] Browser Network tab shows CORS headers in response

