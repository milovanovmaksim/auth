-- +goose Up
CREATE TYPE Role As ENUM('USER', 'ADMIN');

CREATE TABLE users (
  id serial PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  username TEXT NOT NULL,
  password TEXT NOT NULL,
  role Role NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE users;
