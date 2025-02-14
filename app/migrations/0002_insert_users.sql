-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, name, age, gender) VALUES
(1, 'Alice', 25, 'female'),
(2, 'Bob', 30, 'male'),
(3, 'Charlie', 35, 'male'),
(4, 'Diana', 28, 'female'),
(5, 'Eve', 22, 'female');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE id IN (1, 2, 3, 4, 5);
-- +goose StatementEnd