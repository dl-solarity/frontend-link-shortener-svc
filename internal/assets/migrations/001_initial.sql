-- +migrate Up

CREATE TABLE links (
    id char(8) primary key,
    path text NOT NULL,
    value jsonb NOT NULL,
    expired_at timestamp without time zone NOT NULL
);

-- +migrate Down

DROP TABLE IF EXISTS links;
