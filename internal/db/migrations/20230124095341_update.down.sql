-- reverse: modify "room_users" table
ALTER TABLE "room_users" DROP COLUMN "host";
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "protected";
-- reverse: modify "games" table
ALTER TABLE "games" DROP COLUMN "assassin_chance", DROP COLUMN "the_assassinated_ids", ADD COLUMN "the_assassinated_id" bigint NOT NULL DEFAULT 0;
