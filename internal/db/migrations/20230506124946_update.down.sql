-- reverse: modify "games" table
ALTER TABLE "games" DROP COLUMN "closed", DROP COLUMN "result", ADD COLUMN "end_by" character varying NOT NULL DEFAULT 'none';
