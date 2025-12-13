# Quick Deploy to Render

## Fastest Method: Using render.yaml

1. **Go to Render Dashboard**
   - Visit https://dashboard.render.com
   - Sign in

2. **Create Blueprint**
   - Click "New +" → "Blueprint"
   - Connect your GitHub repository: `rakeshsoneja/medial_records_app`
   - **Blueprint File Path**: `deployment/render.yaml`
   - Click "Apply"

3. **Wait for Services to Deploy**
   - Render will create:
     - PostgreSQL database
     - Backend service
     - Frontend service
   - Takes 5-10 minutes

4. **Configure Environment Variables**
   After services are created:
   - Go to backend service → "Environment"
   - Set `JWT_SECRET` (generate: `openssl rand -base64 32`)
   - Go to frontend service → "Environment"
   - Update `REACT_APP_API_URL` with your backend URL

5. **Done!**
   - Backend: `https://medical-records-backend.onrender.com`
   - Frontend: `https://medical-records-frontend.onrender.com`

## Manual Method

If render.yaml doesn't work, follow `DEPLOYMENT_STEPS.md` for manual setup.

## Important Notes

- **Free tier**: Services sleep after 15 min inactivity
- **First request**: May take 30-60 seconds to wake up
- **Database**: Free tier expires after 90 days
- **CORS**: Backend automatically allows frontend URL if `FRONTEND_URL` env var is set

