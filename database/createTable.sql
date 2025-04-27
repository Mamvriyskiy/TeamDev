-- Создание схемы warehouse для history_item
CREATE SCHEMA IF NOT EXISTS warehouse;

-- Таблица warehouse.history_item
CREATE TABLE warehouse.history_item (
    id SERIAL PRIMARY KEY,
    action TEXT NOT NULL,
    item_type TEXT NOT NULL,
    item_id INTEGER NOT NULL,
    item JSONB NOT NULL
);
    
-- Таблица data.user
CREATE TABLE user_account (
    id SERIAL PRIMARY KEY,
    telegram TEXT,
    balance NUMERIC(10, 2) DEFAULT 0,
    blocked BOOLEAN DEFAULT FALSE
);

-- Таблица data.account
CREATE TABLE account (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL,
    social_network TEXT NOT NULL,
    url TEXT NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES user_account(id)
);

-- Таблица data.request
CREATE TABLE request (
    id SERIAL PRIMARY KEY,
    account_id INTEGER NOT NULL,
    requested_followers INTEGER NOT NULL,
    requested_duration INTERVAL NOT NULL,
    prompt TEXT,
    follow_reward NUMERIC(10, 2),
    cost_per_follower NUMERIC(10, 2),
    status TEXT NOT NULL,
    last_changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES account(id)
);

-- Таблица data.offer
CREATE TABLE offer (
    id SERIAL PRIMARY KEY,
    recipient_id INTEGER NOT NULL,
    request_id INTEGER NOT NULL,
    status TEXT NOT NULL,
    last_changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (recipient_id) REFERENCES user_account(id),
    FOREIGN KEY (request_id) REFERENCES request(id)
);

-- Таблица data.subscription
CREATE TABLE subscription (
    id SERIAL PRIMARY KEY,
    offer_id INTEGER NOT NULL,
    followee_id INTEGER NOT NULL,
    last_verified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (offer_id) REFERENCES offer(id),
    FOREIGN KEY (followee_id) REFERENCES user_account(id)
);