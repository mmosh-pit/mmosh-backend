CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY DEFAULT gen_random_uuid()::text,
    uuid TEXT UNIQUE,
    picture TEXT,
    banner TEXT,
    name TEXT,
    display_name TEXT,
    last_name TEXT,
    username TEXT,
    websites JSONB DEFAULT '[]',
    bio TEXT,
    challenges TEXT,
    email TEXT UNIQUE NOT NULL,
    password TEXT,
    telegram JSONB,
    sessions JSONB DEFAULT '[]',
    bluesky JSONB,
    subscription JSONB,
    wallet TEXT,
    referred_by TEXT,
    onboarding_step INT DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    last_login TIMESTAMPTZ,
    profilenft TEXT,
    role TEXT,
    from_bot TEXT,
    deactivated BOOLEAN DEFAULT FALSE,
    seniority INT DEFAULT 0,
    symbol TEXT,
    link TEXT,
    following INT DEFAULT 0,
    follower INT DEFAULT 0,
    connection_nft TEXT,
    connection_badge TEXT,
    connection INT DEFAULT 0,
    is_private BOOLEAN DEFAULT FALSE,
    request BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS email_verification (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    code INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS wallets (
    id SERIAL PRIMARY KEY,
    address TEXT,
    private TEXT,
    email TEXT UNIQUE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS failed_email_attempts (
    id SERIAL PRIMARY KEY,
    email TEXT,
    keypair TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS early_access (
    id SERIAL PRIMARY KEY,
    name TEXT,
    email TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS account_deletion_requests (
    id SERIAL PRIMARY KEY,
    name TEXT,
    email TEXT UNIQUE NOT NULL,
    reason TEXT
);

CREATE TABLE IF NOT EXISTS bots (
    id TEXT PRIMARY KEY DEFAULT gen_random_uuid()::text,
    name TEXT,
    description TEXT,
    image TEXT,
    symbol TEXT,
    key TEXT UNIQUE,
    price DOUBLE PRECISION DEFAULT 0,
    presale_start_date TEXT,
    system_prompt TEXT,
    creator_username TEXT,
    type TEXT,
    default_model TEXT,
    deactivated BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    invite_image TEXT,
    lut TEXT,
    seniority INT DEFAULT 0,
    distribution JSONB,
    invitation_price DOUBLE PRECISION DEFAULT 0,
    discount DOUBLE PRECISION DEFAULT 0,
    telegram TEXT,
    twitter TEXT,
    website TEXT,
    presale_supply INT DEFAULT 0,
    min_presale_supply INT DEFAULT 0,
    presale_end_date TEXT,
    dex_listing_date TEXT,
    creator TEXT,
    code TEXT,
    privacy TEXT,
    status TEXT DEFAULT 'active'
);

CREATE TABLE IF NOT EXISTS activated_agents (
    id SERIAL PRIMARY KEY,
    agent_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    UNIQUE(agent_id, user_id)
);

CREATE TABLE IF NOT EXISTS chat_bots (
    id TEXT PRIMARY KEY,
    name TEXT,
    type TEXT,
    picture TEXT
);

CREATE TABLE IF NOT EXISTS chats (
    id TEXT PRIMARY KEY DEFAULT gen_random_uuid()::text,
    owner TEXT REFERENCES users(id) ON DELETE CASCADE,
    chat_agent JSONB,
    deactivated BOOLEAN DEFAULT FALSE,
    last_message JSONB,
    participants JSONB DEFAULT '[]'
);

CREATE TABLE IF NOT EXISTS messages (
    id TEXT PRIMARY KEY DEFAULT gen_random_uuid()::text,
    chat_id TEXT REFERENCES chats(id) ON DELETE CASCADE,
    content TEXT,
    type TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    sender TEXT,
    agent_id TEXT
);

CREATE TABLE IF NOT EXISTS coin_addresses (
    id SERIAL PRIMARY KEY,
    token TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS receipts (
    id TEXT PRIMARY KEY DEFAULT gen_random_uuid()::text,
    package_name TEXT,
    product_id TEXT,
    purchase_token TEXT UNIQUE,
    wallet TEXT,
    platform TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    expired_at TIMESTAMPTZ,
    is_canceled BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS subscriptions (
    id TEXT PRIMARY KEY DEFAULT gen_random_uuid()::text,
    name TEXT,
    tier INT DEFAULT 0,
    product_id TEXT UNIQUE,
    platform TEXT,
    benefits JSONB DEFAULT '[]'
);

CREATE TABLE IF NOT EXISTS posts (
    id TEXT PRIMARY KEY DEFAULT gen_random_uuid()::text,
    header TEXT NOT NULL,
    sub_header TEXT,
    tags JSONB DEFAULT '[]',
    authors JSONB NOT NULL DEFAULT '[]',
    body TEXT NOT NULL,
    slug TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS themes (
    id TEXT PRIMARY KEY DEFAULT gen_random_uuid()::text,
    name TEXT,
    code_name TEXT,
    background_color TEXT,
    primary_color TEXT,
    secondary_color TEXT,
    logo TEXT
);
