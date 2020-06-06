DROP TYPE IF EXISTS user_status CASCADE; CREATE TYPE user_status AS ENUM ('pending', 'active', 'inactive', 'blocked');

DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
    id              	SERIAL PRIMARY KEY,
    guid	       		TEXT NOT NULL UNIQUE,
    identifier          TEXT NOT NULL UNIQUE,
    identifier_type     TEXT NOT NULL UNIQUE,
    password           	TEXT NOT NULL UNIQUE,
    require_password    BOOLEAN NOT NULL UNIQUE,
    email           	TEXT NOT NULL UNIQUE,
    display_name        TEXT NOT NULL,
    permissions         TEXT[] NOT NULL DEFAULT '{}',
    status          	user_status NOT NULL DEFAULT 'pending',
    confirmation_code   TEXT NOT NULL DEFAULT '',
    confirmation_expiry TIMESTAMP WITH TIME ZONE NOT NULL,

    created_at      	TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at      	TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

DROP INDEX IF EXISTS idx_user_guid; CREATE UNIQUE INDEX idx_user_guid ON users(guid);
DROP INDEX IF EXISTS idx_user_identifier; CREATE UNIQUE INDEX idx_user_identifier ON users(identifier, identifier_type);
