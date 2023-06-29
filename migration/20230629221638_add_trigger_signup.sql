-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION set_initial_debt()
RETURNS TRIGGER AS $$
BEGIN
INSERT INTO debt (user_id, amount)
VALUES (NEW.id, 0);
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER user_signup_trigger
    AFTER INSERT ON users
    FOR EACH ROW
    EXECUTE FUNCTION set_initial_debt();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS user_signup_trigger ON users;
-- +goose StatementEnd
