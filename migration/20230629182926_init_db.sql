-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    group_number VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    lastname VARCHAR(255) NOT NULL,
    firstname VARCHAR(255) NOT NULL,
    fathersname VARCHAR(255) NOT NULL DEFAULT '',
    group_id INT NOT NULL,
    login VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups (id)
);

CREATE TABLE IF NOT EXISTS debt (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL UNIQUE,
    amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE debt
DROP CONSTRAINT IF EXISTS debt_user_id_fkey;
ALTER TABLE users
DROP CONSTRAINT IF EXISTS users_group_id_fkey;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS debt;
DROP TABLE IF EXISTS groups;
-- +goose StatementEnd
