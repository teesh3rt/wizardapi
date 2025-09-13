-- +goose Up
-- +goose StatementBegin
CREATE TABLE wizards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name TEXT NOT NULL,
    level INTEGER NOT NULL DEFAULT 1
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE wizards;
-- +goose StatementEnd