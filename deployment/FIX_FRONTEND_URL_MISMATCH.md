# Fix: FRONTEND_URL Mismatch Error

## The Error

```
Access-Control-Allow-Origin header has a value 'https://medical-records-minp.onrender.com' 
that is not equal to the supplied origin 'https://medical-records-frontend-etnv.onrender.com'
```

## The Problem

Your `FRONTEND_URL` environment variable on Render is set to the **wrong URL**:
- **Configured:** `https://medical-records-minp.onrender.com` ❌
- **Actual Frontend:** `https://medical-records-frontend-etnv.onrender.com` ✅

The browser requires an **exact match** between:
1. The `Origin` header sent by the frontend
2. The `Access-Control-Allow-Origin` header returned by the backend

## The Fix

### Option 1: Update FRONTEND_URL (Recommended)

1. Go to [Render Dashboard](https://dashboard.render.com)
2. Select **Backend Service** (`medical-records-backend`)
3. Click **Environment** tab
4. Find `FRONTEND_URL`
5. **Update it to:** `https://medical-records-frontend-etnv.onrender.com`
6. Click **Save Changes**
7. Wait 2-5 minutes for service to restart

### Option 2: Use Code Fix (Already Deployed)

The code has been updated to **always use the requesting origin**, so even if `FRONTEND_URL` is wrong, CORS will work. However, it's still best to set it correctly.

## Verify the Fix

After updating `FRONTEND_URL`:

1. **Check Backend Logs:**
   ```
   CORS: Origin https://medical-records-frontend-etnv.onrender.com matched configured allowed origins
   ```

2. **Test in Browser:**
   - Open your frontend
   - Try to register
   - Should work without CORS errors

3. **Check Network Tab:**
   - Response header should show:
   ```
   Access-Control-Allow-Origin: https://medical-records-frontend-etnv.onrender.com
   ```

## How to Find Your Actual Frontend URL

1. Open your frontend in browser
2. Look at the address bar - that's your frontend URL
3. Copy it exactly (including `https://`)
4. Use it as `FRONTEND_URL` in backend environment

## Prevention

Always ensure:
- `FRONTEND_URL` in backend matches the **actual frontend URL** shown in browser
- No trailing slash (or include it consistently)
- Use `https://` not `http://`
- Check after creating new services (Render URLs can change)

