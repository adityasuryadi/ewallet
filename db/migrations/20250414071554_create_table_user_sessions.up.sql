CREATE TABLE "user_sessions" (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    token TEXT NOT NULL,
    token_expired TIMESTAMP NOT NULL,
    refresh_token TEXT NOT NULL,
    refresh_token_expired TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
)