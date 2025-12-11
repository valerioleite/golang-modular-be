CREATE SCHEMA IF NOT EXISTS settings;

CREATE TABLE IF NOT EXISTS settings.tenant(
    id UUID PRIMARY KEY,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by VARCHAR NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name VARCHAR(255) NOT NULL,
    logo VARCHAR(255),
    banner VARCHAR(255),
    user_id UUID NOT NULL
);