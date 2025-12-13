# Fix Frontend Build Failures on Render

If your frontend build is failing on Render, follow these steps to diagnose and fix the issue.

## 🔍 Step 1: Check Build Logs

1. Go to [Render Dashboard](https://dashboard.render.com)
2. Click on `medical-records-frontend` service
3. Go to **"Logs"** tab
4. Scroll to see the actual error message
5. Look for lines starting with `npm ERR!` or `Error:`

## 🐛 Common Build Errors and Fixes

### Error 1: "Module not found" or "Cannot find module"

**Cause**: Missing dependencies or dependency version conflicts

**Fix**:
1. Check if `package.json` has all required dependencies
2. Try updating `package-lock.json`:
   - Delete `package-lock.json` from repository (if committed)
   - Or update dependencies in `package.json`

### Error 2: "SyntaxError" or "Unexpected token"

**Cause**: JavaScript/JSX syntax errors in the code

**Fix**:
1. Check the file mentioned in the error
2. Look for:
   - Missing semicolons
   - Unclosed brackets/parentheses
   - Incorrect JSX syntax
   - Import/export errors

### Error 3: "REACT_APP_API_URL is not defined"

**Cause**: Missing environment variable (this is OK, it has a default)

**Fix**: This shouldn't cause a build failure, but you can:
1. Add `REACT_APP_API_URL` to frontend environment variables
2. Or ensure the code handles missing env vars gracefully

### Error 4: "Out of memory" or "JavaScript heap out of memory"

**Cause**: Build process running out of memory

**Fix**:
1. Update `package.json` build script:
   ```json
   "build": "NODE_OPTIONS=--max-old-space-size=4096 react-scripts build"
   ```
2. Or use a different build approach

### Error 5: "TypeError: Cannot read property" during build

**Cause**: Code trying to access properties that don't exist at build time

**Fix**:
1. Check for code that runs at build time (not wrapped in functions)
2. Ensure environment variables are accessed correctly
3. Check for `process.env` usage

## 🔧 Quick Fixes to Try

### Fix 1: Update Build Command

1. Go to frontend service → **"Settings"** tab
2. Find **"Build Command"**
3. Try:
   ```bash
   cd frontend && npm ci && npm run build
   ```
   (Using `npm ci` for clean install)

### Fix 2: Set Node Version

1. Go to frontend service → **"Settings"** tab
2. Add **"Node Version"**: `18` or `20`
3. Save and redeploy

### Fix 3: Check Environment Variables

1. Go to frontend service → **"Environment"** tab
2. Ensure `REACT_APP_API_URL` is set (if needed)
3. Remove any invalid environment variables

### Fix 4: Update package.json Scripts

If build is failing, try updating the build script in `package.json`:

```json
{
  "scripts": {
    "build": "CI=false react-scripts build",
    "build:prod": "NODE_OPTIONS=--max-old-space-size=4096 CI=false react-scripts build"
  }
}
```

## 📋 Build Command Options

### Option 1: Standard Build
```bash
cd frontend && npm install && npm run build
```

### Option 2: Clean Install Build
```bash
cd frontend && npm ci && npm run build
```

### Option 3: Build with Memory Increase
```bash
cd frontend && npm install && NODE_OPTIONS=--max-old-space-size=4096 npm run build
```

### Option 4: Build with CI Flag
```bash
cd frontend && npm install && CI=false npm run build
```

## 🔍 Debugging Steps

### Step 1: Check package.json

Verify `package.json` is valid:
- All dependencies are listed
- No syntax errors in JSON
- Scripts are defined correctly

### Step 2: Test Build Locally

1. Clone repository locally
2. Run:
   ```bash
   cd frontend
   npm install
   npm run build
   ```
3. If it fails locally, fix the issue there first
4. Then commit and push

### Step 3: Check for Large Files

Large files in `public/` or `src/` can cause build issues:
- Check for very large images
- Check for large JSON files
- Optimize assets if needed

### Step 4: Check Node Version

1. Go to frontend service → **"Settings"**
2. Check **"Node Version"**
3. Try setting to: `18` or `20`
4. Save and redeploy

## ✅ Verify Build Configuration

Your `render.yaml` should have:

```yaml
- type: web
  name: medical-records-frontend
  env: node
  buildCommand: cd frontend && npm install && npm run build
  startCommand: cd frontend && npx serve -s build -l ${PORT:-3000}
```

## 🆘 If Nothing Works

1. **Check Render Status**: https://status.render.com
2. **Check Build Logs**: Look for the exact error message
3. **Try Manual Build**: Build locally and commit the `build/` folder (not recommended but works)
4. **Contact Support**: Render support can help with build issues

## 📝 Common Issues Checklist

- [ ] `package.json` is valid JSON
- [ ] All dependencies are listed in `package.json`
- [ ] No syntax errors in source code
- [ ] Node version is compatible (18+)
- [ ] Build command is correct
- [ ] Environment variables are set correctly
- [ ] No large files causing memory issues
- [ ] `react-scripts` version is compatible

## 💡 Pro Tips

1. **Always test builds locally** before pushing
2. **Use `npm ci`** for more reliable installs in CI/CD
3. **Set Node version explicitly** in Render settings
4. **Check build logs carefully** - the error is usually at the bottom
5. **Use `CI=false`** if ESLint warnings are causing failures

## 🔗 Next Steps

After build succeeds:
- Frontend will be accessible at your Render URL
- Update `REACT_APP_API_URL` if needed
- Test the frontend functionality

