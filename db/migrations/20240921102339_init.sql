-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_login text NOT NULL UNIQUE,
    user_email text NOT NULL UNIQUE,
    user_hashed_password text NOT NULL
);
CREATE TABLE IF NOT EXISTS users_info(
    user_id uuid PRIMARY KEY NOT NULL,
    user_phone text UNIQUE,
    user_name text,
    user_surname text,
    user_city text,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);
CREATE TABLE IF NOT EXISTS user_addresses(
    user_id uuid  KEY NOT NULL,
    user_address text NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT pk_user_address PRIMARY KEY (user_id, user_address)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_addresses;
DROP TABLE IF EXISTS users_info;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
