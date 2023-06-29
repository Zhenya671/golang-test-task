-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    add column password varchar(255) not null;

ALTER TABLE users
    drop column login;
ALTER TABLE users
    add column login VARCHAR(255) NOT NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    drop column password;

ALTER TABLE users
    drop column login;
ALTER TABLE users
    add column login VARCHAR(255) NOT NULL default '';
-- +goose StatementEnd
