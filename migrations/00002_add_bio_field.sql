-- +goose Up
-- +goose StatementBegin
ALTER TABLE wizards ADD COLUMN bio TEXT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE wizards DROP COLUMN bio;
-- +goose StatementEnd