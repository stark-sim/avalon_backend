-- create "games" table
CREATE TABLE "games" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- create "game_users" table
CREATE TABLE "game_users" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "user_id" bigint NOT NULL, "game_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- create "rooms" table
CREATE TABLE "rooms" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- create "room_users" table
CREATE TABLE "room_users" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "user_id" bigint NOT NULL, "room_id" bigint NOT NULL, PRIMARY KEY ("id"));
