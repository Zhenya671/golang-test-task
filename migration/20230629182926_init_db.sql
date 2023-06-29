-- +goose Up
-- +goose StatementBegin
-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    lastname VARCHAR(255) NOT NULL,
    firstname VARCHAR(255) NOT NULL,
    fathersname VARCHAR(255) NOT NULL DEFAULT '',
    group_number VARCHAR(255) NOT NULL,
    login VARCHAR(255) NOT NULL DEFAULT ''
    );

-- Create the debt table
CREATE TABLE IF NOT EXISTS debt (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
drop table if exists debt;
-- +goose StatementEnd
