# Fix: Database Connection Failed on Render

The backend is trying to connect to `localhost` instead of your Render database. This means the database environment variables are not set.

## 🔧 Quick Fix: Link Database to Backend

### Step 1: Create PostgreSQL Database (if not done)

1. Go to [Render Dashboard](https://dashboard.render.com)
2. Click **"New +"** → **"PostgreSQL"**
3. Configure:
   - **Name**: `medical-records-db`
   - **Database**: `medical_records` (or leave default)
   - **User**: Leave default (or `medical_user`)
   - **Region**: Same as your backend (e.g., Oregon)
   - **Plan**: Free (or paid)
4. Click **"Create Database"**
5. **Wait for database to be created** (takes 1-2 minutes)

### Step 2: Link Database to Backend Service

**This is the easiest way - Render will auto-add all database environment variables:**

1. Go to `medical-records-backend` service
2. Click **"Environment"** tab
3. Scroll down to **"Add from Database"** section (or look for **"Link Database"** button)
4. Click **"Link Database"** or **"Add from Database"**
5. Select your `medical-records-db` database
6. Render will automatically add these environment variables:
   - `DB_HOST`
   - `DB_PORT`
   - `DB_USER`
   - `DB_PASSWORD`
   - `DB_NAME`
   - `DATABASE_URL` (optional, some services use this)

### Step 3: Add DB_SSLMODE

1. Still in the **"Environment"** tab
2. Click **"Add Environment Variable"**
3. Add:
   - **Key**: `DB_SSLMODE`
   - **Value**: `require`
4. Click **"Save Changes"**

### Step 4: Wait for Redeploy

- Backend will automatically redeploy (takes 2-5 minutes)
- Watch the "Events" or "Logs" tab
- Should see "Live" status when done

## 🔍 Manual Method (If Link Database Doesn't Work)

If you can't use "Link Database", add these manually:

### Get Database Connection Details

1. Go to your `medical-records-db` database service
2. Find the **"Connections"** section
3. Copy these values:
   - **Internal Database URL** (if backend is in same region)
   - **External Database URL** (if backend is in different region)
   - Or get individual values: Host, Port, Database, User, Password

### Add Environment Variables Manually

1. Go to `medical-records-backend` service
2. Click **"Environment"** tab
3. Add each variable:

```
DB_HOST=<your-database-host>
DB_PORT=<your-database-port>
DB_USER=<your-database-user>
DB_PASSWORD=<your-database-password>
DB_NAME=<your-database-name>
DB_SSLMODE=require
```

**Example values** (yours will be different):
```
DB_HOST=dpg-xxxxx-a.oregon-postgres.render.com
DB_PORT=5432
DB_USER=medical_user
DB_PASSWORD=your-actual-password
DB_NAME=medical_records
DB_SSLMODE=require
```

4. Click **"Save Changes"**

## ✅ Verify Environment Variables

After adding, verify these are set in backend service:

- [ ] `DB_HOST` - Should NOT be `localhost`
- [ ] `DB_PORT` - Usually `5432`
- [ ] `DB_USER` - Your database user
- [ ] `DB_PASSWORD` - Your database password
- [ ] `DB_NAME` - Your database name
- [ ] `DB_SSLMODE` - Should be `require`

## 🐛 Troubleshooting

### Still Getting "localhost" Error

**Check:**
1. Environment variables are actually saved (refresh the page)
2. Backend has redeployed after adding variables
3. No typos in variable names (case-sensitive: `DB_HOST` not `db_host`)

### "Connection Refused" Error

**Check:**
1. Database service is running (should show "Available")
2. Database and backend are in the same region (for internal connections)
3. Using correct host (internal vs external URL)
4. `DB_SSLMODE=require` is set

### "Password Authentication Failed"

**Check:**
1. `DB_PASSWORD` is correct (copy from database service)
2. `DB_USER` matches the database user
3. No extra spaces in environment variable values

### "Database Does Not Exist"

**Check:**
1. `DB_NAME` matches the actual database name
2. Database was created successfully
3. User has permissions on the database

## 📋 Complete Environment Variables Checklist

Your backend service should have:

### Database Variables (Required):
- [ ] `DB_HOST` - Database hostname
- [ ] `DB_PORT` - Database port (usually 5432)
- [ ] `DB_USER` - Database username
- [ ] `DB_PASSWORD` - Database password
- [ ] `DB_NAME` - Database name
- [ ] `DB_SSLMODE` - Set to `require`

### Other Variables:
- [ ] `JWT_SECRET` - Should be auto-generated or set manually
- [ ] `APP_ENV` - Set to `production`
- [ ] `FRONTEND_URL` - Your frontend URL (for CORS)
- [ ] `SERVER_PORT` - Can be empty (Render uses `PORT`)

## 🎯 Test Database Connection

After backend redeploys, check the logs:

1. Go to backend service → **"Logs"** tab
2. Look for:
   - ✅ "Server starting on port..."
   - ✅ No database connection errors
   - ❌ If you see "localhost" in error, variables aren't set

## 💡 Pro Tip

**Use "Link Database" feature** - It's the easiest and most reliable way. Render automatically:
- Adds all required variables
- Uses correct internal/external URLs
- Handles SSL configuration
- Updates if database credentials change

## 🔗 Next Steps

After database is connected:
1. Backend should start successfully
2. Run migrations automatically
3. Create default user (run seed script)
4. Test your app!

See [NEXT_STEPS.md](./NEXT_STEPS.md) for complete setup guide.

