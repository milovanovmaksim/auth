-- +goose Up
CREATE TABLE users (
  id serial PRIMARY KEY,
  email VARCHAR(32) NOT NULL UNIQUE,
  username VARCHAR(32) NOT NULL,
  password VARCHAR(64) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE users;
