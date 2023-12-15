-- Create enum type "chefs"
CREATE TYPE "public"."chefs" AS ENUM ('me', 'ordered', 'outsourced');
-- Modify "meals" table
ALTER TABLE "public"."meals" ADD COLUMN "chef" "public"."chefs" NOT NULL;
-- Create "meal_dishes" table
CREATE TABLE "public"."meal_dishes" ("id" serial NOT NULL, "name" text NOT NULL, "meal_id" integer NOT NULL, "classification" "public"."meal_classification" NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "meal_dishes_meal_id_fkey" FOREIGN KEY ("meal_id") REFERENCES "public"."meals" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Drop "meal_items" table
DROP TABLE "public"."meal_items";
