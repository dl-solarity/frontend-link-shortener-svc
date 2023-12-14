-- +migrate Up

CREATE TABLE links (
    id char(64) primary key,
    path text NOT NULL,
    value jsonb NOT NULL,
    created_at timestamp without time zone
);

-- +migrate Down

DROP TABLE IF EXISTS links;
