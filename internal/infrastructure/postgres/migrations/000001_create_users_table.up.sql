CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255),
    nickname VARCHAR(64) NOT NULL,
    pass_hash VARCHAR(64) NOT NULL
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_nickname ON users(nickname);