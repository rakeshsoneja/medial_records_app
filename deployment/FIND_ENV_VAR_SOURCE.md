# Finding Where `medical-records-minp.onrender.com` is Set

Since you can't find it in the environment variables UI, it's likely set in one of these places:

## Places to Check

### 1. Render Blueprint (render.yaml)

The `render.yaml` file might have it hardcoded:

1. Check `render.yaml` in root directory
2. Check `deployment/render.yaml`
3. Look for `FRONTEND_URL` or `ALLOWED_ORIGINS` in `envVars` section

**If found:** Remove or update it, then redeploy the blueprint.

### 2. Render Service Settings - Hidden/System Variables

Some variables might be set at the service level:

1. Backend Service → **Settings** tab
2. Look for any "System" or "Internal" environment variables
3. Check if there's a "Show all variables" option

### 3. Render Blueprint Deployment

If you deployed using a Blueprint:

1. Go to Render Dashboard
2. Find your Blueprint (if you used one)
3. Check the Blueprint's environment variables
4. The Blueprint might override service-level variables

### 4. Previous Service/Deleted Service

The URL might be from a previous service:

1. Check **all services** in Render (including suspended/deleted)
2. Look for any service with "minp" in the name
3. Check if that service's environment variables are somehow linked

### 5. Render's Internal Cache

Render might be caching old environment variables:

**Solution:** Force a complete rebuild:
1. Backend Service → **Manual Deploy**
2. Click **"Clear build cache & deploy"**
3. This forces fresh environment variable loading

### 6. Database-Linked Variables

If you linked a database, check:

1. Database Service → **Settings**
2. Look for any environment variables set there
3. Linked databases can sometimes pass variables to services

### 7. Render API/CLI

If you used Render API or CLI:

1. Check if you have any Render CLI configuration files
2. Check `.render.yaml` or similar files
3. Check if you used `render.yaml` from a different branch

## How to Find It - Step by Step

### Step 1: Check render.yaml Files

```bash
# In your local repository
grep -r "minp" render.yaml deployment/render.yaml
grep -r "FRONTEND_URL" render.yaml deployment/render.yaml
grep -r "ALLOWED_ORIGINS" render.yaml deployment/render.yaml
```

### Step 2: Check Render Dashboard - All Services

1. Go to Render Dashboard
2. List **ALL** services (active, suspended, deleted)
3. For each service, check Environment variables
4. Search for "minp" or "medical-records-minp"

### Step 3: Check Blueprint (if used)

1. Render Dashboard → **Blueprints** (if you used one)
2. Check the Blueprint's configuration
3. Look for environment variables in the Blueprint

### Step 4: Export All Environment Variables

1. Backend Service → **Environment** tab
2. Take a screenshot or copy ALL variables
3. Check if there's a "Show system variables" option
4. Look for any variable containing "minp"

### Step 5: Check Service Logs at Startup

When the service starts, it logs environment variables:

1. Backend Service → **Logs** tab
2. Look at the very beginning of the logs (when service starts)
3. Look for any log messages showing environment variables
4. Search for "minp" in the logs

## Nuclear Option: Reset All CORS Variables

If you can't find it, reset everything:

1. **Remove** `FRONTEND_URL` (temporarily)
2. **Remove** `ALLOWED_ORIGINS` (if it exists)
3. **Clear build cache & deploy**
4. The new code will allow all origins (with warning)
5. Then add back `FRONTEND_URL` with correct value

## The Code Solution

The latest code **always uses the requesting origin**, so:
- Even if `FRONTEND_URL` is wrong, CORS will work
- The code doesn't care about environment variables for the response
- It only uses them for logging

Once the new code deploys, CORS should work regardless of what's in environment variables.

## Quick Test

After deploying the new code with debug logging, check the logs:

```
CORS DEBUG: FRONTEND_URL=..., ALLOWED_ORIGINS=..., Request Origin=...
```

This will show you exactly what environment variables are being read, even if they're not visible in the UI.

