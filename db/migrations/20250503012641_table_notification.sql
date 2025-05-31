-- +goose Up
-- +goose StatementBegin
CREATE TABLE `notifications` (
    `id` varchar(36) NOT NULL,
    `account_id` varchar(36) DEFAULT NULL,
    `sub_id` varchar(36) DEFAULT NULL,
    `content` longtext NOT NULL,
    `route` varchar(255) NOT NULL,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `read_at` datetime,
    PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `notifications`;
-- +goose StatementEnd
