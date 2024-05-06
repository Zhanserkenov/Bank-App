-- Файл: 000002_create_accounts_table.up.sql

CREATE TABLE accounts (
                          id SERIAL PRIMARY KEY,
                          created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          deleted_at TIMESTAMP WITH TIME ZONE,
                          type TEXT NOT NULL,
                          name TEXT NOT NULL,
                          user_id INTEGER NOT NULL REFERENCES users(id),
                          balance_kzt INTEGER NOT NULL DEFAULT 0,
                          balance_eur INTEGER NOT NULL DEFAULT 0,
                          balance_usd INTEGER NOT NULL DEFAULT 0
);
