-- Схема для практики занятия 6.
CREATE TABLE IF NOT EXISTS users (
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS accounts (
    id      SERIAL PRIMARY KEY,
    balance INTEGER NOT NULL DEFAULT 0 CHECK (balance >= 0)
);

INSERT INTO accounts (balance) VALUES (100), (0);
