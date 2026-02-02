-- internal/adapters/storage/postgres/sql/schema.sql

CREATE TABLE users (
    id UUID PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role TEXT NOT NULL -- ADMIN, FIELD, OBSERVER, VIEWER
);

CREATE TABLE audits (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    norm TEXT NOT NULL, -- ISO 9001, etc.
    status TEXT NOT NULL, -- PLANNED, IN_PROGRESS, FINISHED
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE assignments (
    user_id UUID REFERENCES users(id),
    audit_id UUID REFERENCES audits(id),
    sector_id TEXT NOT NULL,
    PRIMARY KEY (user_id, audit_id, sector_id)
);