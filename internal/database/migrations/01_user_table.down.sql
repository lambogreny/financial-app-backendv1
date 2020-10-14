CREATE EXTENSION  IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
    user_ud UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL,
    password_hash byte,
    create_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX user_email
    ON users (email);