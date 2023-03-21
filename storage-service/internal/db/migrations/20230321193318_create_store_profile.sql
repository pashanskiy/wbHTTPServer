-- +goose Up
-- +goose StatementBegin
CREATE TABLE "stores_profiles" (
    "uid" TEXT NOT NULL UNIQUE,
    "name" TEXT NOT NULL,
    "address" TEXT,
    "working" INTEGER,
    "owner" TEXT,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
    "deleted_at" TIMESTAMP,
    PRIMARY KEY("uid")
);
-- +goose StatementEnd