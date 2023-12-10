-- Create enum type "meal_type"
CREATE TYPE "meal_type" AS ENUM ('breakfast', 'lunch', 'dinner');
-- Create enum type "meal_classification"
CREATE TYPE "meal_classification" AS ENUM ('green', 'orange', 'cheat_day');
-- Create "activities" table
CREATE TABLE "activities" ("id" serial NOT NULL, "name" text NOT NULL, "date" date NOT NULL, PRIMARY KEY ("id"));
-- Create "meals" table
CREATE TABLE "meals" ("id" serial NOT NULL, "type" "meal_type" NOT NULL, "classification" "meal_classification" NOT NULL, "date" date NOT NULL, PRIMARY KEY ("id"));
-- Create "meal_items" table
CREATE TABLE "meal_items" ("id" serial NOT NULL, "name" text NOT NULL, "meal_id" integer NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "meal_items_meal_id_fkey" FOREIGN KEY ("meal_id") REFERENCES "meals" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
