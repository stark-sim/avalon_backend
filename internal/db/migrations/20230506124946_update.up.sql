-- modify "games" table
ALTER TABLE "games" DROP COLUMN "end_by", ADD COLUMN "result" character varying NOT NULL DEFAULT 'none', ADD COLUMN "closed" boolean NOT NULL DEFAULT false;
