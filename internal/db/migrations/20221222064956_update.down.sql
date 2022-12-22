-- reverse: modify "cards" table
ALTER TABLE "cards" DROP COLUMN "tale", DROP COLUMN "role", ALTER COLUMN "name" SET DEFAULT '';
