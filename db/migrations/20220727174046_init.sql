-- +migrate Up
CREATE table users (
    id SERIAL,
    name text,
    username text,
    password text
);

-- +migrate Down
DROP TABLE users;