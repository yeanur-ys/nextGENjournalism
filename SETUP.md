# Local Development Setup

## 1) Environment
```bash
cp .env.example .env
```

## 2) Start infrastructure
```bash
docker compose up -d --build
```

Services:
- Web: `http://localhost:3000`
- API: `http://localhost:8080`
- Neo4j Browser: `http://localhost:7474`
- Debezium Connect: `http://localhost:8083`

## 3) Backend only (optional)
```bash
cd apps/api
cp .env.example .env
go mod tidy
go run ./cmd/api
```

## 4) Frontend only (optional)
```bash
cd apps/web
npm install
npm run dev
```

## Notes
- PostgreSQL migrations in `database/postgres/migrations` are mounted to `/docker-entrypoint-initdb.d`.
- Neo4j constraints are initialized at API startup and in `database/neo4j/constraints.cypher`.
