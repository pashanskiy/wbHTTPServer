-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users_profiles" (
	"uid"	TEXT NOT NULL UNIQUE,
	"surname"	TEXT NOT NULL,
	"name"	TEXT,
	"secondname"	TEXT,
	"age"	INTEGER,
	"created_at"	TIMESTAMP NOT NULL,
	"updated_at"	TIMESTAMP NOT NULL,
	"deleted_at"	TIMESTAMP,
	PRIMARY KEY("uid")
);
-- +goose StatementEnd