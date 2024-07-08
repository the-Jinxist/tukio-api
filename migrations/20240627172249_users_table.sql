-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "users" (
  "id" uuid PRIMARY KEY,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "verified" boolean NOT NULL,
  "created_at" timestamp NOT NULL
);

CREATE UNIQUE INDEX idx_users_email ON users (email);

CREATE TABLE IF NOT EXISTS "profiles" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "phone_number" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,

  FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "profiles";
DROP TABLE IF EXISTS "users";

-- +goose StatementEnd
