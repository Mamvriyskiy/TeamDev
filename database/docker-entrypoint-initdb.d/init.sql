-- Таблица data.user
CREATE TABLE user_account (
    id UUID PRIMARY KEY,
    telegram INT,
    balance INT,
    blocked BOOLEAN  -- Изменено на BOOLEAN
);

-- Таблица data.account
CREATE TABLE account (
    id UUID PRIMARY KEY,
    owner_id UUID,
    social_network VARCHAR(255),
    url VARCHAR(255),
    FOREIGN KEY (owner_id) REFERENCES user_account(id)
);

-- Таблица data.request
CREATE TABLE request (
    id UUID PRIMARY KEY,
    account_id UUID,
    requested_followers INTEGER NOT NULL,
    requested_duration INTEGER NOT NULL,
    prompt INTEGER NOT NULL,
    follow_reward INTEGER NOT NULL,  -- Исправлено: добавлен пробел
    cost_per_follower INTEGER NOT NULL,
    status TEXT,  -- Изменено на TEXT
    last_changed_at TIMESTAMP,  -- Изменено на TIMESTAMP
    FOREIGN KEY (account_id) REFERENCES account(id)
);

-- Таблица data.offer
CREATE TABLE offer (
    id UUID PRIMARY KEY,
    recipient_id UUID,
    request_id UUID,
    status TEXT,  -- Изменено на TEXT
    last_changed_at TIMESTAMP,  -- Изменено на TIMESTAMP
    FOREIGN KEY (recipient_id) REFERENCES user_account(id),
    FOREIGN KEY (request_id) REFERENCES request(id)
);

-- Таблица data.subscription
CREATE TABLE subscription (
    id UUID PRIMARY KEY,
    offer_id UUID,
    followee_id UUID,
    last_verified_at TIMESTAMP,  -- Изменено на TIMESTAMP
    FOREIGN KEY (offer_id) REFERENCES offer(id),
    FOREIGN KEY (followee_id) REFERENCES user_account(id)
);
