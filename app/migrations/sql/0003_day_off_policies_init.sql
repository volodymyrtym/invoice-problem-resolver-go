-- +goose Up
CREATE TABLE day_off_policies (
                        id UUID PRIMARY KEY,           -- Поле id типу UUID
                        user_id UUID NOT NULL,         -- Поле user_id типу UUID
                        half_day BOOLEAN NOT NULL,     -- Поле half_day (логічне значення)
                        approvable BOOLEAN NOT NULL,   -- Поле approvable (логічне значення)
                        name VARCHAR(32) NOT NULL      -- Поле name обмежене довжиною 32 символи (VARCHAR)
);

CREATE INDEX idx_policy_user_id ON day_off_policies(user_id);

-- +goose Down
DROP TABLE day_off_policies;