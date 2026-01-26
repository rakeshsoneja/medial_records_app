# How to Start the Frontend

## Prerequisites

- Node.js 18 or higher installed
- Backend server running on `http://localhost:8080`

## Quick Start

### Step 1: Navigate to Frontend Directory

```bash
cd frontend
```

### Step 2: Install Dependencies

```bash
npm install
```

This will install all required packages listed in `package.json`.

**Note:** This only needs to be done once, or when dependencies change.

### Step 3: Create .env File (Optional)

Create a `.env` file in the `frontend` directory if you want to customize the API URL:

```env
REACT_APP_API_URL=http://localhost:8080/api/v1
```

**Note:** If you don't create this file, it will default to `http://localhost:8080/api/v1`

### Step 4: Start the Development Server

```bash
npm start
```

The frontend will:
- Start on `http://localhost:3000`
- Automatically open in your browser
- Hot-reload when you make changes to the code

## Verify It's Working

1. Open your browser to `http://localhost:3000`
2. You should see the login page
3. The frontend will connect to the backend API automatically

## Common Commands

### Start Development Server
```bash
npm start
```

### Build for Production
```bash
npm run build
```

### Run Tests
```bash
npm test
```

## Troubleshooting

### "Port 3000 already in use"
- Close the other application using port 3000
- Or set a different port: `PORT=3001 npm start`

### "Cannot connect to backend"
- Make sure the backend server is running on `http://localhost:8080`
- Check the `REACT_APP_API_URL` in `.env` file
- Verify CORS is configured in the backend

### "Module not found" errors
- Run `npm install` again
- Delete `node_modules` folder and `package-lock.json`, then run `npm install`

### "npm: command not found"
- Install Node.js from https://nodejs.org/
- Make sure Node.js is in your PATH

## Development Tips

- The app uses hot-reload, so changes will appear automatically
- Check the browser console for any errors
- Network requests can be viewed in browser DevTools (F12)

## Production Build

To create an optimized production build:

```bash
npm run build
```

This creates a `build` folder with optimized static files that can be served by any web server.




