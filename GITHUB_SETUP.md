# GitHub Repository Setup Guide

This guide will help you commit and push this project to your GitHub repository.

## Prerequisites

- Git installed on your system
- A GitHub account
- GitHub repository created (or we'll create one)

## Step 1: Initialize Git Repository

If git is not already initialized:

```bash
# From project root directory
git init
```

## Step 2: Verify .gitignore File

Make sure `.gitignore` exists in the root directory. It should exclude:
- Environment files (`.env`)
- Dependencies (`node_modules/`, `vendor/`)
- Build outputs
- IDE files
- Database files

## Step 3: Create GitHub Repository (if not created)

1. Go to https://github.com/new
2. Create a new repository (e.g., `medical-records-app`)
3. **DO NOT** initialize with README, .gitignore, or license (we already have these)
4. Copy the repository URL (e.g., `https://github.com/yourusername/medical-records-app.git`)

## Step 4: Add All Files to Git

```bash
# Add all files to staging
git add .

# Check what will be committed
git status
```

## Step 5: Make Initial Commit

```bash
git commit -m "Initial commit: Medical Records Management App

- Backend: Go API with PostgreSQL
- Frontend: React application
- Features: Prescriptions, Appointments, Lab Reports, Medications, Reminders, Insurance
- Secure sharing with time-limited links
- JWT authentication
- Swagger API documentation"
```

## Step 6: Connect to GitHub Repository

```bash
# Add remote repository (replace with your actual repository URL)
git remote add origin https://github.com/yourusername/your-repo-name.git

# Verify remote was added
git remote -v
```

## Step 7: Push to GitHub

```bash
# Push to main branch (or master if that's your default)
git branch -M main
git push -u origin main
```

If you get authentication errors, you may need to:
- Use a Personal Access Token instead of password
- Set up SSH keys
- Use GitHub CLI

## Step 8: Verify on GitHub

1. Go to your GitHub repository
2. Verify all files are present
3. Check that `.env` files are NOT included (they should be in .gitignore)

## Important Files to Check Before Committing

### ✅ Should be committed:
- All source code files
- `package.json`, `go.mod`, `go.sum`
- `README.md`, `SETUP.md`, documentation files
- `.gitignore`
- `docker-compose.yml`
- Configuration examples (`.env.example`)

### ❌ Should NOT be committed (already in .gitignore):
- `.env` files (contain secrets)
- `node_modules/` (dependencies)
- `vendor/` (Go dependencies)
- Build outputs
- Database files
- IDE files

## Quick Setup Script

You can run these commands in sequence:

```bash
# 1. Initialize git
git init

# 2. Add all files
git add .

# 3. Initial commit
git commit -m "Initial commit: Medical Records Management App"

# 4. Add remote (replace with your repo URL)
git remote add origin https://github.com/yourusername/your-repo-name.git

# 5. Push to GitHub
git branch -M main
git push -u origin main
```

## Troubleshooting

### "Repository already exists" error
- The directory might already be a git repository
- Check: `git status`
- If it's already initialized, skip `git init`

### "Authentication failed" error
- Use Personal Access Token instead of password
- Or set up SSH keys: https://docs.github.com/en/authentication/connecting-to-github-with-ssh

### "Remote origin already exists"
- Remove existing remote: `git remote remove origin`
- Then add your new remote

### Want to use a different branch name?
```bash
git branch -M master  # or any other branch name
git push -u origin master
```

## Next Steps After Pushing

1. Add repository description on GitHub
2. Add topics/tags (e.g., `go`, `react`, `medical-records`, `postgresql`)
3. Create a README badge if desired
4. Set up branch protection rules (optional)
5. Add collaborators (optional)

## Security Reminders

⚠️ **IMPORTANT**: Before pushing, verify:
- No `.env` files are committed
- No API keys or secrets in code
- Database passwords are in `.env.example` only (with placeholder values)
- JWT secrets are placeholders in example files

