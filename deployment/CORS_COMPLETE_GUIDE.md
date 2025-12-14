# Complete CORS Fix Guide for Render Deployment

## Why CORS Errors Happen

**CORS (Cross-Origin Resource Sharing)** is a browser security feature that blocks requests from one domain to another unless the server explicitly allows it.

### Common CORS Scenarios on Render:

1. **Different Domains**: Frontend (`medical-records-frontend.onrender.com`) and Backend (`medical-records-backend.onrender.com`) are on different subdomains
2. **HTTPS vs HTTP**: Render uses HTTPS, but if your code expects HTTP, it can cause issues
3. **Missing Preflight Handling**: Browsers send OPTIONS requests before actual requests - these must be handled
4. **Missing Headers**: Backend must send specific CORS headers in responses

## How Our Solution Works

### 1. **Dedicated CORS Middleware** (`backend/internal/middleware/cors.go`)

- ✅ Handles OPTIONS preflight requests automatically
- ✅ Configures allowed origins based on environment variables
- ✅ Supports both development and production
- ✅ Includes logging for debugging
- ✅ Follows security best practices

### 2. **Environment-Based Configuration**

- **Development**: Allows `localhost:3000`, `localhost:3001`
- **Production**: Uses `FRONTEND_URL` environment variable
- **Fallback**: If `FRONTEND_URL` not set, allows all origins (with warning)

### 3. **Proper OPTIONS Handling**

The middleware automatically:
- Intercepts OPTIONS requests
- Returns appropriate CORS headers
- Returns 204 status code
- Allows the actual request to proceed

## Setup Instructions

### Step 1: Set Environment Variables on Render

#### Backend Service Environment Variables:

1. Go to [Render Dashboard](https://dashboard.render.com)
2. Select your **Backend Service** (`medical-records-backend`)
3. Go to **Environment** tab
4. Add/Update these variables:

```bash
# Required for CORS
APP_ENV=production
FRONTEND_URL=https://medical-records-frontend.onrender.com

# Optional: If you have multiple frontend URLs
ALLOWED_ORIGINS=https://medical-records-frontend.onrender.com,https://another-frontend.onrender.com
```

**Important Notes:**
- Use `https://` (not `http://`) for Render URLs
- Don't include trailing slash in `FRONTEND_URL` (or include it consistently)
- `APP_ENV=production` ensures production CORS rules are used

#### Frontend Service Environment Variables:

```bash
REACT_APP_API_URL=https://medical-records-backend.onrender.com/api/v1
```

### Step 2: Verify the Code is Deployed

The CORS middleware is in:
- `backend/internal/middleware/cors.go` (new file)
- `backend/internal/router/router.go` (updated to use new middleware)

Make sure these files are committed and pushed to GitHub. Render will auto-deploy.

### Step 3: Test CORS

#### Test 1: Health Endpoint (No CORS needed, but good for testing)

```bash
curl https://medical-records-backend.onrender.com/health
```

Should return: `{"status":"ok"}`

#### Test 2: OPTIONS Preflight Request

```bash
curl -X OPTIONS https://medical-records-backend.onrender.com/api/v1/auth/register \
  -H "Origin: https://medical-records-frontend.onrender.com" \
  -H "Access-Control-Request-Method: POST" \
  -H "Access-Control-Request-Headers: Content-Type" \
  -v
```

**Expected Response Headers:**
```
Access-Control-Allow-Origin: https://medical-records-frontend.onrender.com
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD
Access-Control-Allow-Headers: Origin, Content-Type, Content-Length, Accept, Authorization, X-Requested-With, X-CSRF-Token, Cache-Control
Access-Control-Allow-Credentials: true
Access-Control-Max-Age: 43200
```

#### Test 3: Actual API Request

```bash
curl -X POST https://medical-records-backend.onrender.com/api/v1/auth/register \
  -H "Origin: https://medical-records-frontend.onrender.com" \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"test123","name":"Test User"}' \
  -v
```

**Expected Response Headers:**
```
Access-Control-Allow-Origin: https://medical-records-frontend.onrender.com
Access-Control-Allow-Credentials: true
```

### Step 4: Check Browser Console

1. Open your frontend: `https://medical-records-frontend.onrender.com`
2. Open Browser DevTools (F12)
3. Go to **Console** tab
4. Try to register/login
5. Check for CORS errors

**If you still see CORS errors:**
- Check **Network** tab → Look at the failed request
- Check **Response Headers** - Are CORS headers present?
- Check **Request Headers** - Is `Origin` header correct?

## Troubleshooting

### Issue 1: "No 'Access-Control-Allow-Origin' header"

**Cause**: Backend not sending CORS headers

**Fix**:
1. Verify `APP_ENV=production` is set in backend environment
2. Verify `FRONTEND_URL` is set correctly
3. Check backend logs for CORS configuration messages
4. Ensure code is deployed (check Render deployment logs)

### Issue 2: "Preflight request doesn't pass access control check"

**Cause**: OPTIONS request failing

**Fix**:
1. Our middleware handles OPTIONS automatically
2. Verify the middleware is applied (check `router.go`)
3. Check backend logs for OPTIONS request handling

### Issue 3: "Credentials flag is true, but Access-Control-Allow-Credentials is not 'true'"

**Cause**: Credentials not allowed

**Fix**:
- Our middleware sets `AllowCredentials: true`
- Verify frontend is not sending credentials incorrectly

### Issue 4: CORS works in Postman/curl but not in browser

**Cause**: Browser enforces CORS, tools don't

**Fix**:
- This is expected - browser CORS is stricter
- Use browser DevTools to debug
- Check Network tab for actual request/response headers

### Issue 5: CORS works locally but not on Render

**Cause**: Different URLs/environments

**Fix**:
1. Verify `APP_ENV=production` on Render
2. Verify `FRONTEND_URL` matches actual Render frontend URL
3. Check that URLs use `https://` (not `http://`)
4. Ensure no trailing slash mismatches

## Best Practices

### ✅ DO:

1. **Set `FRONTEND_URL` explicitly** - Don't rely on "allow all origins" fallback
2. **Use HTTPS** - Render uses HTTPS, so always use `https://` in URLs
3. **Test preflight requests** - Use browser DevTools to verify OPTIONS handling
4. **Monitor logs** - Check Render logs for CORS configuration messages
5. **Use environment variables** - Don't hardcode URLs in code

### ❌ DON'T:

1. **Don't disable CORS** - It's a security feature
2. **Don't use wildcard origins with credentials** - Browser blocks this
3. **Don't hardcode URLs** - Use environment variables
4. **Don't ignore CORS errors** - Fix them properly
5. **Don't use `*` for `Access-Control-Allow-Origin` with credentials** - Browser blocks this

## Render-Specific Considerations

### 1. **HTTPS Only**
- Render uses HTTPS for all services
- Always use `https://` in URLs
- Don't mix `http://` and `https://`

### 2. **Dynamic URLs**
- Render URLs can change (if you delete/recreate service)
- Use environment variables, not hardcoded URLs
- Update `FRONTEND_URL` if frontend URL changes

### 3. **Service Communication**
- Services on Render can communicate internally
- But browser → backend requests still need CORS
- CORS is enforced by the browser, not the server

### 4. **Environment Variables**
- Set in Render Dashboard → Service → Environment
- Changes require service restart
- Check logs to verify variables are loaded

## Verification Checklist

- [ ] `APP_ENV=production` set in backend environment
- [ ] `FRONTEND_URL` set to actual frontend URL (with `https://`)
- [ ] Code deployed to GitHub and Render auto-deployed
- [ ] Backend logs show CORS configuration
- [ ] OPTIONS requests return 204 with CORS headers
- [ ] Actual API requests include CORS headers in response
- [ ] Browser DevTools shows no CORS errors
- [ ] Frontend can successfully call backend APIs

## Additional Resources

- [MDN CORS Documentation](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
- [Render Environment Variables](https://render.com/docs/environment-variables)
- [Gin CORS Middleware](https://github.com/gin-contrib/cors)

## Quick Fix Command

If you need to quickly test CORS configuration:

```bash
# Test from your local machine
curl -X OPTIONS https://medical-records-backend.onrender.com/api/v1/auth/register \
  -H "Origin: https://medical-records-frontend.onrender.com" \
  -H "Access-Control-Request-Method: POST" \
  -v
```

Look for `Access-Control-Allow-Origin` in the response headers.

