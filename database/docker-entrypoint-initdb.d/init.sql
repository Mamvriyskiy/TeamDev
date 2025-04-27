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
    telegram nvarchar(32),
    balance int,
    blocked bit
);

-- Таблица data.account
CREATE TABLE account (
    id SERIAL PRIMARY KEY,
    owner_id uniqueidentifier,
    social_network nvarchar(max),
    url nvarchar(max),
    FOREIGN KEY (owner_id) REFERENCES user_account(id)
);

-- Таблица data.request
CREATE TABLE request (
    id SERIAL PRIMARY KEY,
    account_id uniqueidentifier,
    requested_followers INTEGER NOT NULL,
    requested_duration INTEGER NOT NULL,
    prompt INTEGER NOT NULL,
    follow_rewardINTEGER NOT NULL,
    cost_per_follower INTEGER NOT NULL,
    status nvarchar(max),
    last_changed_at datetime(2)	,
    FOREIGN KEY (account_id) REFERENCES account(id)
);

-- Таблица data.offer
CREATE TABLE offer (
    id SERIAL PRIMARY KEY,
    recipient_id uniqueidentifier,
    request_id uniqueidentifier,
    status nvarchar(max),
    last_changed_at datetime(2),
    FOREIGN KEY (recipient_id) REFERENCES user_account(id),
    FOREIGN KEY (request_id) REFERENCES request(id)
);

-- Таблица data.subscription
CREATE TABLE subscription (
    id SERIAL PRIMARY KEY,
    offer_id uniqueidentifier,
    followee_id uniqueidentifier,
    last_verified_at datetime(2),
    FOREIGN KEY (offer_id) REFERENCES offer(id),
    FOREIGN KEY (followee_id) REFERENCES user_account(id)
);
