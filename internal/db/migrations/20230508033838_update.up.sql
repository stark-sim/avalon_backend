-- modify "game_users" table
ALTER TABLE "game_users" ADD COLUMN "exited" boolean NOT NULL DEFAULT false;
