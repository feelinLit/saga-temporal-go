-- +goose Up
CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    status INT NOT NULL,
    client_id BIGINT NOT NULL,
    item_id BIGINT NOT NULL,
    item_count INT NOT NULL,
    total_price INT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS orders;
