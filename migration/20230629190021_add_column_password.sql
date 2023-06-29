-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    add column password varchar(255) not null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    drop column password varchar(255) not null;
-- +goose StatementEnd
