-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users_profiles" (
	"uid"	TEXT NOT NULL UNIQUE,
	"surname"	TEXT,
	"name"	TEXT,
	"secondname"	TEXT,
	"age"	INTEGER,
	"created_at"	INTEGER NOT NULL,
	"updated_at"	INTEGER NOT NULL,
	"deleted_at"	INTEGER NOT NULL,
	PRIMARY KEY("uid")
)
-- +goose StatementEnd