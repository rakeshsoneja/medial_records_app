# Force CORS Fix - Complete Reset

Since you can't find `medical-records-minp.onrender.com` in environment variables, the old code is likely still running. Let's force a complete reset.

## Step 1: Verify Code is Deployed

1. Go to [Render Dashboard](https://dashboard.render.com)
2. Backend Service → **Events** tab
3. Check latest deployment timestamp
4. Should show recent deployment with message: "Fix CORS: Always use requesting origin..."

**If not deployed:**
- Wait 2-5 minutes for auto-deploy, OR
- Go to **Manual Deploy** → **Deploy latest commit**

## Step 2: Force Complete Restart

1. Backend Service → **Manual Deploy** tab
2. Click **"Clear build cache & deploy"**
3. Wait 2-5 minutes for deployment

This ensures:
- Fresh code build
- Fresh environment variable loading
- No cached values

## Step 3: Check Backend Logs After Restart

After restart, look for these log messages:

**Good (new code running):**
```
CORS DEBUG: FRONTEND_URL=https://medical-records-frontend-etnv.onrender.com, ALLOWED_ORIGINS=, Request Origin=https://medical-records-frontend-etnv.onrender.com
CORS DEBUG: Setting Access-Control-Allow-Origin to: https://medical-records-frontend-etnv.onrender.com
```

**Bad (old code still running):**
- No "CORS DEBUG" messages
- Shows old URL in logs

## Step 4: Test OPTIONS Request

After restart, test:

```bash
curl -X OPTIONS https://medical-records-backend.onrender.com/api/v1/auth/register \
  -H "Origin: https://medical-records-frontend-etnv.onrender.com" \
  -H "Access-Control-Request-Method: POST" \
  -v 2>&1 | grep -i "access-control"
```

**Expected output:**
```
< access-control-allow-origin: https://medical-records-frontend-etnv.onrender.com
```

**If you see:**
```
< access-control-allow-origin: https://medical-records-minp.onrender.com
```
→ Old code is still running, need to force restart again

## Step 5: Nuclear Option - Suspend and Resume

If still not working:

1. Backend Service → **Settings** tab
2. Scroll to bottom
3. Click **Suspend** (wait 30 seconds)
4. Click **Resume** (wait 2-5 minutes)
5. This forces a complete service restart

## Step 6: Verify in Browser

1. Open: `https://medical-records-frontend-etnv.onrender.com`
2. Open DevTools (F12) → **Network** tab
3. Try to register
4. Check failed request → **Response Headers**
5. Should show: `Access-Control-Allow-Origin: https://medical-records-frontend-etnv.onrender.com`

## Why This Should Work

The latest code **always uses the requesting origin**, so:
- Even if environment variables are wrong, it will work
- The code doesn't care about `FRONTEND_URL` for the actual response
- It only uses `FRONTEND_URL` for logging

The debug logs will show exactly what's happening.

## If Still Not Working

Share:
1. **Backend logs** (last 30 lines with "CORS" in them)
2. **OPTIONS test result** (from Step 4)
3. **Browser Network tab** (screenshot of failed request)

The debug logs will reveal what's actually happening.

