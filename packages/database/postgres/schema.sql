CREATE TABLE IF NOT EXISTS journalists (
  id UUID PRIMARY KEY,
  handle TEXT NOT NULL UNIQUE,
  reputation NUMERIC(8, 4) NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS articles (
  id UUID PRIMARY KEY,
  journalist_id UUID NOT NULL REFERENCES journalists(id),
  parent_article_id UUID NULL REFERENCES articles(id),
  title TEXT NOT NULL,
  body TEXT NOT NULL,
  published_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
