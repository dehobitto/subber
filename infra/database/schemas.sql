CREATE TABLE IF NOT EXISTS subscriptions (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    repository VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',
    token VARCHAR(255),
    last_seen_tag VARCHAR(100)
);