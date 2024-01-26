-- Modify "meal_dishes" table
ALTER TABLE "public"."meal_dishes" ALTER COLUMN "score" SET DEFAULT 'green';
-- Create "activities_type" table
CREATE TABLE "public"."activities_type" ("id" serial NOT NULL, "name" text NOT NULL, "date" date NOT NULL, PRIMARY KEY ("id"));
-- Modify "activities" table
ALTER TABLE "public"."activities" DROP COLUMN "name", ADD COLUMN "type_id" integer NOT NULL, ADD COLUMN "count" integer NOT NULL, ADD CONSTRAINT "activities_type_id_fkey" FOREIGN KEY ("type_id") REFERENCES "public"."activities_type" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
