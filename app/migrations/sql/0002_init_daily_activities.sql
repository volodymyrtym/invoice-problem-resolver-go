-- +goose Up

CREATE TABLE daily_activity_daily_activities
(
    id          UUID NOT NULL PRIMARY KEY,
    user_id     UUID NOT NULL,
    start_at    TIMESTAMP(0) NOT NULL,
    end_at      TIMESTAMP(0) NOT NULL,
    description TEXT NOT NULL,
    created_at  TIMESTAMP(0) NOT NULL DEFAULT NOW(),
    project     VARCHAR (16) DEFAULT NULL
);

CREATE INDEX user_created_at
    ON daily_activity_daily_activities (user_id, created_at);

CREATE INDEX user_start_date_idx
    ON daily_activity_daily_activities (user_id, start_at);

CREATE INDEX user_activity_range_idx
    ON daily_activity_daily_activities (user_id, start_at, end_at);

-- +goose Down
DROP TABLE IF EXISTS daily_activity_daily_activities;