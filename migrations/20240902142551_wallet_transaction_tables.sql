-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "wallets" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "amount" varchar NOT NULL,
  "updated_at" timestamp NOT NULL,

  FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE TABLE IF NOT EXISTS "transactions" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "amount" varchar NOT NULL,
  "reference" varchar NOT NULL,
  "type" varchar NOT NULL,
  "updated_at" timestamp NOT NULL,
  "created_at" timestamp NOT NULL,

  FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE UNIQUE INDEX idx_transactions_reference ON transactions (reference);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "wallets";
DROP TABLE IF EXISTS "transactions";
-- +goose StatementEnd
