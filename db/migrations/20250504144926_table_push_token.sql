-- +goose Up
-- +goose StatementBegin
CREATE TABLE `push_tokens` (
    `account_id` VARCHAR(36) NOT NULL,
    `exponent_push_token` VARCHAR(255) NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`account_id`, `exponent_push_token`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `push_tokens`;
-- +goose StatementEnd
