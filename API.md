# API Documentation (Sprint 1)

Base URL: `/api/v1`

## Health
- `GET /healthz`

## Auth
- `POST /api/v1/auth/register`
  - Body: `{ "email", "password", "displayName", "role" }`
- `POST /api/v1/auth/login`
  - Body: `{ "email", "password" }`
  - Returns JWT token and user profile

## Profile
- `GET /api/v1/profile`
  - Auth: `Authorization` header with a JWT bearer token

## Articles
- `POST /api/v1/articles`
  - Auth required; roles: `journalist`, `admin`
  - Body: `{ "title", "content", "status" }`
- `GET /api/v1/articles/{articleID}`
  - Public article read
- `PUT /api/v1/articles/{articleID}`
  - Auth required; roles: `journalist`, `admin`
  - Body: `{ "title", "content", "status" }`
