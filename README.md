# Feedback System API

Go backend for collecting and managing user feedback.

## Overview

This project provides an API for working with user feedback:

- **Get feedback** with pagination, search, and sorting support
- **Submit new feedback** with rating and text

## Quick Start

```bash
# Copy configuration
cp .env.example .env

# Run with Docker Compose
docker-compose up --build
```

API available at: `http://localhost:2510`

## API

- `GET /api` — get feedback list
  - Query params: `skip`, `take`, `search`, `sortBy`, `caseSensitive`, `wholeWord`
- `POST /api` — submit new feedback
  - Body: `{ "rating": 5, "feedback_text": "..." }`
