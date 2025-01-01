-- +goose Up
CREATE TABLE user_users
(
id            UUID NOT NULL PRIMARY KEY,
created_at    TIMESTAMP(0) NOT NULL DEFAULT NOW(),
last_login_at TIMESTAMP(0) DEFAULT NULL,
email         VARCHAR(320) NOT NULL,
password      VARCHAR(60)  NOT NULL
);

CREATE UNIQUE INDEX user_users_email_idx
ON user_users (email);

-- +goose Down
DROP TABLE IF EXISTS user_users;