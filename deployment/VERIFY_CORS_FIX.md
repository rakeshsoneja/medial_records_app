# Verify CORS Fix is Working

## Current Status

✅ `FRONTEND_URL` is correctly set to: `https://medical-records-frontend-etnv.onrender.com`

## Why You're Still Seeing the Error

The error shows the backend is returning:
```
Access-Control-Allow-Origin: https://medical-records-minp.onrender.com
```

But it should return:
```
Access-Control-Allow-Origin: https://medical-records-frontend-etnv.onrender.com
```

This means either:
1. **Code hasn't deployed yet** - The fix I just pushed needs to deploy (2-5 minutes)
2. **Service needs restart** - Environment variable changes require restart
3. **Old code is still running** - Render might be using cached build

## Step-by-Step Fix

### Step 1: Verify Code is Deployed

1. Go to [Render Dashboard](https://dashboard.render.com)
2. Select **Backend Service** (`medical-records-backend`)
3. Go to **Events** tab
4. Look for latest deployment
5. Should show: `Deployed successfully` with recent timestamp
6. If not deployed yet, wait 2-5 minutes

### Step 2: Manually Restart Backend Service

Even if code is deployed, restart to ensure environment variables are loaded:

1. Go to **Backend Service** → **Manual Deploy** tab
2. Click **Clear build cache & deploy**
3. Wait for deployment to complete (2-5 minutes)

### Step 3: Verify Environment Variable

1. Go to **Backend Service** → **Environment** tab
2. Verify `FRONTEND_URL` shows exactly:
   ```
   https://medical-records-frontend-etnv.onrender.com
   ```
3. **No trailing slash** - should end with `.com` not `.com/`
4. **Must be `https://`** not `http://`

### Step 4: Check Backend Logs

After restart, check logs for:

**Good signs:**
```
CORS: Origin https://medical-records-frontend-etnv.onrender.com matched configured allowed origins
```
or
```
CORS: WARNING - Origin https://medical-records-frontend-etnv.onrender.com does not match configured FRONTEND_URL [...], but allowing it anyway
```

**Bad signs:**
```
CORS: Origin ... matched configured allowed origins: [https://medical-records-minp.onrender.com]
```
(This means old environment variable is still being used)

### Step 5: Test OPTIONS Request

Run this command (replace with your actual backend URL):

```bash
curl -X OPTIONS https://medical-records-backend.onrender.com/api/v1/auth/register \
  -H "Origin: https://medical-records-frontend-etnv.onrender.com" \
  -H "Access-Control-Request-Method: POST" \
  -v
```

**Expected Response:**
- Status: `204 No Content`
- Header: `Access-Control-Allow-Origin: https://medical-records-frontend-etnv.onrender.com` ✅

**If you see:**
- `Access-Control-Allow-Origin: https://medical-records-minp.onrender.com` ❌
  → Service hasn't restarted or code hasn't deployed

### Step 6: Test in Browser

1. Open: `https://medical-records-frontend-etnv.onrender.com`
2. Open DevTools (F12) → **Network** tab
3. Try to register
4. Check the failed request:
   - **Request Headers** → `Origin: https://medical-records-frontend-etnv.onrender.com`
   - **Response Headers** → `Access-Control-Allow-Origin: https://medical-records-frontend-etnv.onrender.com` ✅

## Quick Fix: Force Restart

If nothing works, force a complete restart:

1. **Backend Service** → **Settings** tab
2. Scroll to bottom
3. Click **Suspend** (wait 10 seconds)
4. Click **Resume** (wait 2-5 minutes for restart)

This ensures fresh environment variable loading.

## The Code Fix

The latest code I pushed will **always use the requesting origin**, so even if `FRONTEND_URL` is wrong, CORS will work. But you should still set it correctly.

Once the code deploys, CORS should work regardless of `FRONTEND_URL` value, but correct `FRONTEND_URL` ensures proper logging.

## Still Not Working?

If after all steps it still doesn't work:

1. **Share backend logs** (last 20 lines with CORS messages)
2. **Share OPTIONS test result** (from Step 5)
3. **Share browser Network tab** (screenshot of failed request headers)

The code fix should make it work automatically once deployed!

