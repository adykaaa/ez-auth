CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE Users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT,
    email TEXT UNIQUE,
    email_verified TIMESTAMP,
    image TEXT
);

CREATE TABLE Accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES Users(id) ON DELETE CASCADE,
    type TEXT,
    provider TEXT,
    provider_account_id TEXT,
    refresh_token TEXT,
    access_token TEXT,
    expires_at TIMESTAMP,
    token_type TEXT,
    scope TEXT,
    id_token TEXT,
    session_state TEXT
);

CREATE TABLE UserRefreshTokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES Users(id) ON DELETE CASCADE,
    refresh_token TEXT
);
