CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
 ( id            SERIAL
 , uuid          UUID      DEFAULT uuid_generate_v1mc()
 , created_at    TIMESTAMP NOT NULL DEFAULT NOW()
 , edited_at     TIMESTAMP NOT NULL DEFAULT NOW()
 , username      TEXT      NOT NULL
 , discord_id    TEXT      NOT NULL
 , banned        BOOLEAN   NOT NULL DEFAULT FALSE
 , banned_mod_id UUID
 , banned_reason TEXT
 );

CREATE INDEX IF NOT EXISTS users_uuid ON users(uuid);
CREATE INDEX IF NOT EXISTS users_discord_id on users(discord_id);
