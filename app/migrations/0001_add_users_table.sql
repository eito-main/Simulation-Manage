-- +goose Up
-- +goose StatementBegin
CREATE TYPE gender_type AS ENUM ('male', 'female');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INT,
    gender gender_type
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS gender_type;
-- +goose StatementEnd
