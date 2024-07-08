-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "events" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "desc" varchar NOT NULL,
  "picture" varchar NOT NULL,
  "location" varchar NOT NULL,
  "dress_code" varchar,
  "event_time" timestamp NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,

  FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE TABLE IF NOT EXISTS "events_ticket_categories" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "desc" varchar NOT NULL,
  "price" numeric NOT NULL,
  "event_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,

  FOREIGN KEY ("event_id") REFERENCES "events" ("id")
)

CREATE TABLE IF NOT EXISTS "user_tickets" (
  "id" uuid PRIMARY KEY,
  "event_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "category_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "desc" varchar NOT NULL,
  "price" numeric NOT NULL,
  "quantity" int NOT NULL,
  "valid" boolean NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,

  FOREIGN KEY ("event_id") REFERENCES "events" ("id")
  FOREIGN KEY ("user_id") REFERENCES "users" ("id")
  FOREIGN KEY ("category_id") REFERENCES "events_ticket_categories" ("id")

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "events";
DROP TABLE IF EXISTS "events_ticket_categories";
DROP TABLE IF EXISTS "user_tickets";
-- +goose StatementEnd
