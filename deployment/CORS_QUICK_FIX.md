# Quick CORS Fix for Render

## Immediate Steps (5 minutes)

### 1. Set Backend Environment Variables

Go to Render Dashboard → Backend Service → Environment → Add:

```bash
APP_ENV=production
FRONTEND_URL=https://medical-records-frontend.onrender.com
```

**Replace `medical-records-frontend.onrender.com` with your actual frontend URL!**

### 2. Verify Frontend Environment Variable

Go to Render Dashboard → Frontend Service → Environment → Verify:

```bash
REACT_APP_API_URL=https://medical-records-backend.onrender.com/api/v1
```

**Replace `medical-records-backend.onrender.com` with your actual backend URL!**

### 3. Wait for Redeploy

- Render will automatically restart the backend service
- Wait 2-5 minutes
- Check service logs to verify it started

### 4. Test

Open your frontend in browser and try to register/login. CORS errors should be gone.

## Still Not Working?

1. **Check Backend Logs** (Render Dashboard → Backend → Logs)
   - Look for: `CORS: Production mode with allowed origins: [...]`
   - If you see "WARNING: Allowing all origins", `FRONTEND_URL` is not set

2. **Verify URLs Match Exactly**
   - Frontend URL in `FRONTEND_URL` must match exactly what browser shows
   - Use `https://` not `http://`
   - Check for trailing slash consistency

3. **Check Browser Console**
   - Open DevTools (F12) → Console tab
   - Look for specific CORS error message
   - Check Network tab → Failed request → Response Headers

4. **Read Full Guide**: See `CORS_COMPLETE_GUIDE.md` for detailed troubleshooting

## Common Mistakes

❌ **Wrong**: `FRONTEND_URL=http://localhost:3000` (local URL on Render)  
✅ **Right**: `FRONTEND_URL=https://medical-records-frontend.onrender.com`

❌ **Wrong**: `FRONTEND_URL=medical-records-frontend.onrender.com` (missing https://)  
✅ **Right**: `FRONTEND_URL=https://medical-records-frontend.onrender.com`

❌ **Wrong**: `APP_ENV=development` (uses dev CORS rules)  
✅ **Right**: `APP_ENV=production`

