-- +goose Up
-- +goose StatementBegin
ALTER TABLE "events_ticket_categories" 
ADD "seat_number" int NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "events_ticket_categories" DROP COLUMN "seat_number";
-- +goose StatementEnd
