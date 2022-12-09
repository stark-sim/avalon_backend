-- reverse: create "votes" table
DROP TABLE "votes";
-- reverse: create "squads" table
DROP TABLE "squads";
-- reverse: create "records" table
DROP TABLE "records";
-- reverse: create "missions" table
DROP TABLE "missions";
-- reverse: modify "rooms" table
ALTER TABLE "rooms" DROP COLUMN "closed";
-- reverse: modify "games" table
ALTER TABLE "games" DROP COLUMN "room_id", DROP COLUMN "capacity", DROP COLUMN "end_by";
