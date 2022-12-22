-- modify "cards" table
ALTER TABLE "cards" ALTER COLUMN "name" SET DEFAULT 'Merlin', ADD COLUMN "role" character varying NOT NULL, ADD COLUMN "tale" character varying NOT NULL DEFAULT '';
