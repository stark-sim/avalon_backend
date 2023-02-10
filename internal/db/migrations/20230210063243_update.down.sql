-- reverse: modify "games" table
ALTER TABLE "games" ALTER COLUMN "the_assassinated_ids" SET NOT NULL;
-- reverse: modify "cards" table
ALTER TABLE "cards" DROP COLUMN "red";
