-- +migrate Up

CREATE TYPE telegram_access_levels_enum AS ENUM ('owner', 'admin', 'member', 'self', 'left', 'banned');

CREATE TABLE IF NOT EXISTS responses (
    id UUID PRIMARY KEY,
    status TEXT NOT NULL,
    error TEXT
--     payload JSONB,
--     created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNIQUE,
    username TEXT UNIQUE,
    telegram_id BIGINT PRIMARY KEY,
    access_hash BIGINT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    phone TEXT UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS users_id_idx ON users(id);
CREATE INDEX IF NOT EXISTS users_username_idx ON users(username);
CREATE INDEX IF NOT EXISTS users_telegramId_idx ON users(telegram_id);

CREATE TABLE IF NOT EXISTS links (
    id SERIAL PRIMARY KEY,
    link TEXT NOT NULL,
    UNIQUE(link)
);
INSERT INTO links (link) VALUES ('HELP TG API');
INSERT INTO links (link) VALUES ('WE vs. ACS');

CREATE INDEX IF NOT EXISTS links_link_idx ON links(link);

CREATE TABLE IF NOT EXISTS permissions (
    request_id TEXT NOT NULL,
    telegram_id INT NOT NULL,
    link TEXT NOT NULL,
    access_level telegram_access_levels_enum NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (telegram_id, link),
    FOREIGN KEY(telegram_id) REFERENCES users(telegram_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(link) REFERENCES links(link) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX IF NOT EXISTS permissions_telegramId_idx ON permissions(telegram_id, link);
CREATE INDEX IF NOT EXISTS permissions_link_idx ON permissions(telegram_id, link);

-- +migrate Down

DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS responses;
DROP TABLE IF EXISTS links;
DROP TABLE IF EXISTS users;

DROP INDEX IF EXISTS users_id_idx;
DROP INDEX IF EXISTS users_username_idx;
DROP INDEX IF EXISTS users_telegramId_idx;

DROP INDEX IF EXISTS links_link_idx;

DROP INDEX IF EXISTS permissions_telegramId_idx;
DROP INDEX IF EXISTS permissions_link_idx;

DROP TYPE IF EXISTS telegram_access_levels_enum;