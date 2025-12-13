# Verify Database Environment Variables on Render

If you're still getting "localhost" errors, the database environment variables are not set. Follow these steps to verify and fix.

## 🔍 Step 1: Verify Database Exists

1. Go to [Render Dashboard](https://dashboard.render.com)
2. Check if you have a PostgreSQL database service
3. It should be named something like `medical-records-db`
4. Status should be **"Available"** (green)

**If you don't have a database:**
- Click "New +" → "PostgreSQL"
- Create it first, then continue

## 🔍 Step 2: Check Backend Environment Variables

1. Go to `medical-records-backend` service
2. Click **"Environment"** tab
3. **Look for these variables** (scroll through the list):

### Required Variables:
- [ ] `DB_HOST` - Should NOT be empty, should NOT be "localhost"
- [ ] `DB_PORT` - Usually "5432"
- [ ] `DB_USER` - Your database username
- [ ] `DB_PASSWORD` - Your database password (hidden)
- [ ] `DB_NAME` - Your database name
- [ ] `DB_SSLMODE` - Should be "require"

### Alternative Variable:
- [ ] `DATABASE_URL` - Full connection string (if Render provides this)

## ❌ If Variables Are Missing

### Option A: Use "Link Database" (Easiest)

1. In the **"Environment"** tab
2. Scroll to **"Add from Database"** section (usually at the bottom)
3. Or look for a **"Link Database"** button
4. Click it
5. Select your `medical-records-db` database
6. Click **"Link"** or **"Add"**
7. Render will automatically add all database variables
8. **Save Changes**

### Option B: Add Variables Manually

1. Go to your database service (`medical-records-db`)
2. Find the **"Connections"** section
3. Copy the connection details:
   - **Internal Database URL** (if backend is in same region)
   - Or individual values: Host, Port, Database, User, Password

4. Go back to `medical-records-backend` → **"Environment"** tab
5. Click **"Add Environment Variable"** for each:

   **Variable 1:**
   - Key: `DB_HOST`
   - Value: (copy from database, e.g., `dpg-xxxxx-a.oregon-postgres.render.com`)

   **Variable 2:**
   - Key: `DB_PORT`
   - Value: `5432`

   **Variable 3:**
   - Key: `DB_USER`
   - Value: (copy from database)

   **Variable 4:**
   - Key: `DB_PASSWORD`
   - Value: (copy from database - click "Show" to reveal)

   **Variable 5:**
   - Key: `DB_NAME`
   - Value: (copy from database, usually `medical_records` or default name)

   **Variable 6:**
   - Key: `DB_SSLMODE`
   - Value: `require`

6. Click **"Save Changes"** after adding all variables

## ✅ Step 3: Verify Variables Are Saved

1. Refresh the page
2. Go back to **"Environment"** tab
3. Verify all variables are still there
4. Check that `DB_HOST` is NOT "localhost"

## 🔄 Step 4: Wait for Redeploy

1. After saving, backend will automatically redeploy
2. Go to **"Events"** or **"Logs"** tab
3. Watch for deployment progress
4. Look for the new debug log that shows database config:
   ```
   Database config - Host: dpg-xxxxx..., Port: 5432, User: ..., DB: ..., SSLMode: require
   ```
5. Should NOT see "localhost" in the logs

## 🐛 Step 5: Check Logs After Redeploy

After backend redeploys, check the logs:

### ✅ Success Looks Like:
```
Database config - Host: dpg-xxxxx-a.oregon-postgres.render.com, Port: 5432, User: medical_user, DB: medical_records, SSLMode: require
Server starting on port 8080
```

### ❌ Still Failing Looks Like:
```
Database config - Host: localhost, Port: 5432, User: medical_user, DB: medical_records, SSLMode: disable
Database connection failed: failed to connect to `host=localhost...
```

**If you still see "localhost":**
- Environment variables are not being read
- Check variable names are exactly: `DB_HOST`, `DB_PORT`, etc. (case-sensitive)
- Make sure you clicked "Save Changes"
- Try removing and re-adding the variables

## 🔍 Alternative: Check DATABASE_URL

Some Render setups provide `DATABASE_URL` instead of individual variables:

1. Go to backend → **"Environment"** tab
2. Look for `DATABASE_URL`
3. If it exists, it should look like:
   ```
   postgres://user:password@host:port/dbname?sslmode=require
   ```
4. The code now supports this format automatically

## 📋 Quick Checklist

- [ ] Database service exists and is "Available"
- [ ] Backend service has `DB_HOST` (not empty, not "localhost")
- [ ] Backend service has `DB_PORT`
- [ ] Backend service has `DB_USER`
- [ ] Backend service has `DB_PASSWORD`
- [ ] Backend service has `DB_NAME`
- [ ] Backend service has `DB_SSLMODE` = `require`
- [ ] All variables saved (clicked "Save Changes")
- [ ] Backend redeployed after adding variables
- [ ] Logs show correct database host (not "localhost")

## 💡 Pro Tip

**Use "Link Database"** - It's the most reliable method. Render automatically:
- Adds all required variables with correct values
- Uses internal URLs for same-region services
- Handles SSL configuration
- Updates if database credentials change

## 🆘 Still Not Working?

If you've verified all variables are set correctly but still getting errors:

1. **Check variable names** - Must be exactly: `DB_HOST`, `DB_PORT`, etc. (uppercase)
2. **Check for typos** - No extra spaces, correct values
3. **Try removing and re-adding** - Sometimes Render needs a refresh
4. **Check database is running** - Database service should be "Available"
5. **Check regions match** - Backend and database should be in same region for internal connections

## 📞 Next Steps

After database connects successfully:
- Backend will run migrations automatically
- Create default user (run seed script)
- Test your app!

See [NEXT_STEPS.md](./NEXT_STEPS.md) for complete setup.

