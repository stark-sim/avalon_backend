-- modify "games" table
ALTER TABLE "games" ADD COLUMN "end_by" character varying NOT NULL DEFAULT 'none', ADD COLUMN "capacity" smallint NOT NULL DEFAULT 0, ADD COLUMN "room_id" bigint NOT NULL;
-- modify "rooms" table
ALTER TABLE "rooms" ADD COLUMN "closed" boolean NOT NULL DEFAULT false;
-- create "missions" table
CREATE TABLE "missions" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "sequence" smallint NOT NULL, "status" character varying NOT NULL DEFAULT 'picking', "failed" boolean NOT NULL DEFAULT false, "capacity" smallint NOT NULL DEFAULT 0, "leader" bigint NOT NULL, "game_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- create "records" table
CREATE TABLE "records" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "user_id" bigint NOT NULL, "score" integer NOT NULL DEFAULT 0, "room_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- create "squads" table
CREATE TABLE "squads" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "user_id" bigint NOT NULL, "rat" boolean NOT NULL DEFAULT false, "mission_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- create "votes" table
CREATE TABLE "votes" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "user_id" bigint NOT NULL, "pass" boolean NOT NULL DEFAULT false, "mission_id" bigint NOT NULL, PRIMARY KEY ("id"));
