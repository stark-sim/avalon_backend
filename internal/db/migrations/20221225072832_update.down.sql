-- reverse: modify "votes" table
ALTER TABLE "votes" DROP COLUMN "voted";
-- reverse: modify "squads" table
ALTER TABLE "squads" DROP COLUMN "acted";
