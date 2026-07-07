-- Sprint 1: core schema (users + articles)
-- Applied automatically on first container start via docker-entrypoint-initdb.d
CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- for gen_random_uuid()

CREATE TYPE user_role AS ENUM ('journalist', 'auditor', 'reader', 'admin');
CREATE TYPE article_status AS ENUM ('draft', 'published', 'retracted');
CREATE TYPE verification_status AS ENUM ('unverified', 'pending', 'verified', 'rejected');

CREATE TABLE users (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email           VARCHAR(255) UNIQUE NOT NULL,
    password_hash   VARCHAR(255) NOT NULL,
    display_name    VARCHAR(120) NOT NULL,
    role            user_role NOT NULL DEFAULT 'reader',
    bio             TEXT DEFAULT '',
    -- FR-2a: auditor identity verification (credential linking)
    credential_url  VARCHAR(500),           -- e.g. link to press card / org profile
    verification    verification_status NOT NULL DEFAULT 'unverified',
    rank_score      DOUBLE PRECISION NOT NULL DEFAULT 0, -- populated in Sprint 3
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE articles (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    author_id           UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title               VARCHAR(300) NOT NULL,
    content             TEXT NOT NULL,
    status              article_status NOT NULL DEFAULT 'draft',
    corruption_factor   DOUBLE PRECISION NOT NULL DEFAULT 0, -- populated in Sprint 3
    
    -- FR-4c: whether this row has been synced into Neo4j yet.
    -- Sprint 1 syncs synchronously in the same request (see services/article_service.go),
    -- this flag exists so a future CDC/worker process can safely retry failures
    -- without re-reading all of Postgres.
    synced_to_graph     BOOLEAN NOT NULL DEFAULT FALSE,
    
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_articles_author_id ON articles(author_id);
CREATE INDEX idx_articles_status ON articles(status);

-- keep updated_at fresh
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_articles_updated_at BEFORE UPDATE ON articles
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();
