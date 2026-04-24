CREATE TABLE IF NOT EXISTS links (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    slug        VARCHAR(255) NOT NULL,
    original_url TEXT        NOT NULL,
    metadata    JSONB       NOT NULL DEFAULT '{}',
    analytics   JSONB       NOT NULL DEFAULT '[]',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_links_slug ON links (slug);
