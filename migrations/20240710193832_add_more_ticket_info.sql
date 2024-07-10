-- +goose Up
-- +goose StatementBegin
ALTER TABLE "user_tickets" 
ADD "seat_number" int NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "user_tickets" DROP COLUMN "seat_number";
-- +goose StatementEnd
