-- modify "cards" table
ALTER TABLE "cards" ADD COLUMN "red" boolean NOT NULL DEFAULT false;
-- modify "games" table
ALTER TABLE "games" ALTER COLUMN "the_assassinated_ids" DROP NOT NULL;
