-- Файл: 000003_create_transactions_table.up.sql

CREATE TABLE transactions (
                              id SERIAL PRIMARY KEY,
                              created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              deleted_at TIMESTAMP WITHOUT TIME ZONE,
                              from INTEGER NOT NULL REFERENCES users(id),
                              to INTEGER NOT NULL REFERENCES users(id),
                              amount INTEGER NOT NULL,
                              currency CHAR(3)
);
