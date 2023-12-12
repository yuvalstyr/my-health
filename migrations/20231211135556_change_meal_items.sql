-- Add value to enum type: "meal_classification"
ALTER TYPE "public"."meal_classification" ADD VALUE 'cheat' AFTER 'orange';
-- Modify "meal_items" table
ALTER TABLE "public"."meal_items" ADD COLUMN "classification" "public"."meal_classification" NOT NULL;
-- Modify "meals" table
ALTER TABLE "public"."meals" DROP COLUMN "classification";
