-- modify "squads" table
ALTER TABLE "squads" ADD COLUMN "acted" boolean NOT NULL DEFAULT false;
-- modify "votes" table
ALTER TABLE "votes" ADD COLUMN "voted" boolean NOT NULL DEFAULT false;
