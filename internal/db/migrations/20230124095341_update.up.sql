-- modify "games" table
ALTER TABLE "games" DROP COLUMN "the_assassinated_id", ADD COLUMN "the_assassinated_ids" jsonb NOT NULL, ADD COLUMN "assassin_chance" smallint NOT NULL DEFAULT 1;
-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "protected" boolean NOT NULL DEFAULT false;
-- modify "room_users" table
ALTER TABLE "room_users" ADD COLUMN "host" boolean NOT NULL DEFAULT false;
