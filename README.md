# nextGENjournalism

Sprint 1 foundation for the MVC-based nextGENjournalism platform.

## Stack
- **Frontend**: Next.js 16 (App Router) + TypeScript (`apps/web`)
- **Backend**: Go 1.22 + chi router + JWT + bcrypt (`apps/api`)
- **Databases**: PostgreSQL 15 + Neo4j 5
- **Infra**: Docker Compose with PostgreSQL, Neo4j, Redis, Zookeeper, Kafka, Debezium, API, and Web

## Repository Structure
- `apps/web` — frontend pages/components for auth, profile, dashboards, and article flows
- `apps/api` — MVC-style Go API (controllers, services, models, middleware, router)
- `database/postgres/migrations` — PostgreSQL schema/migrations
- `database/neo4j` — Neo4j constraints

## Quick Start
See [SETUP.md](./SETUP.md) for local development and [API.md](./API.md) for endpoint docs.
