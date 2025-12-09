CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE IF NOT EXISTS users.user(
    id         UUID PRIMARY KEY,
    created_by VARCHAR   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by VARCHAR   NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    sub        VARCHAR   NOT NULL,
    email      VARCHAR   NOT NULL,
    full_name  VARCHAR   NOT NULL,
    username   VARCHAR
);