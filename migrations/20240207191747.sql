-- Create enum type "value_type"
CREATE TYPE "public"."value_type" AS ENUM ('sum', 'count');
-- Modify "activities" table
ALTER TABLE "public"."activities" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE text, ALTER COLUMN "type_id" TYPE text, DROP COLUMN "count", ADD COLUMN "value" integer NOT NULL;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."activities_id_seq";
-- Create index "activities_date_type_id_key" to table: "activities"
CREATE UNIQUE INDEX "activities_date_type_id_key" ON "public"."activities" ("date", "type_id");
-- Modify "meal_dishes" table
ALTER TABLE "public"."meal_dishes" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE text, ALTER COLUMN "meal_id" TYPE text;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."meal_dishes_id_seq";
-- Modify "meals" table
ALTER TABLE "public"."meals" ALTER COLUMN "id" DROP DEFAULT, ALTER COLUMN "id" TYPE text;
-- Drop sequence used by serial column "id"
DROP SEQUENCE IF EXISTS "public"."meals_id_seq";
-- Create "activity_types" table
CREATE TABLE "public"."activity_types" ("id" text NOT NULL, "name" text NOT NULL, "value_type" "public"."value_type" NOT NULL DEFAULT 'sum', PRIMARY KEY ("id"));
-- Drop "activities_type" table
DROP TABLE "public"."activities_type";
