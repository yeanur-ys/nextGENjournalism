# nextGENjournalism

A monorepo blueprint for a transparent journalism platform with role-based frontend dashboards, a Go API backend, and a Python graph-analysis worker.

## Monorepo Architecture

```text
nextGENjournalism/
├── .github/workflows/
│   ├── frontend-ci.yml
│   └── backend-ci.yml
├── apps/
│   ├── frontend/
│   │   ├── src/app/
│   │   │   ├── layout.tsx
│   │   │   ├── page.tsx
│   │   │   ├── profile/[journalistId]/page.tsx
│   │   │   ├── journalist/{layout.tsx,dashboard/page.tsx,publish/page.tsx,appeals/page.tsx}
│   │   │   ├── auditor/{layout.tsx,dashboard/page.tsx,claims/[claimId]/page.tsx}
│   │   │   └── admin/{layout.tsx,dashboard/page.tsx,compliance/page.tsx}
│   │   ├── src/components/{ui/Button.tsx,ui/Input.tsx,LineageGraph.tsx,DashboardNav.tsx}
│   │   ├── src/graph/{sigma-config.ts,shaders/*,hooks/useSemanticZoom.ts}
│   │   ├── src/lib/{api.ts,crypto.ts}
│   │   ├── package.json
│   │   ├── next.config.js
│   │   └── tsconfig.json
│   ├── go-backend/
│   │   ├── cmd/api/main.go
│   │   ├── internal/{server,auth,articles,consensus,ranking,compliance,kafka}
│   │   ├── go.mod
│   │   └── go.sum
│   └── python-worker/
│       ├── src/{main.py,louvain.py,config.py}
│       ├── requirements.txt
│       └── Dockerfile
├── packages/
│   ├── config-eslint/index.json
│   ├── config-typescript/base.json
│   └── database/
│       ├── index.ts
│       ├── postgres/{schema.sql,migrations/0001_init.sql}
│       ├── neo4j/{schema.cypher,queries.ts}
│       └── redis/{client.ts,keys.ts}
├── infra/
│   ├── docker-compose.yml
│   ├── postgres/postgresql.conf
│   ├── kafka/server.properties
│   ├── debezium/register-postgres.json
│   └── neo4j/conf/neo4j.conf
├── package.json
├── pnpm-workspace.yaml
└── turbo.json
```

## Current Starter Implementation

- `apps/frontend` contains role-routed Next.js pages and starter graph/client utility modules.
- `apps/go-backend` contains a compilable HTTP server with starter route groups and domain packages.
- `apps/python-worker` contains a polling worker skeleton and Louvain clustering integration.
- `packages/database` and `infra` include initial SQL/Cypher/Redis and local infrastructure scaffolding.

## Development

Install dependencies:

```bash
corepack pnpm install
```

Run workspace quality checks:

```bash
corepack pnpm lint
corepack pnpm check-types
corepack pnpm build
```
