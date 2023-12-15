-- Create enum type "meal_score"
CREATE TYPE "public"."meal_score" AS ENUM ('green', 'orange', 'cheat');
-- Modify "meal_dishes" table
ALTER TABLE "public"."meal_dishes" DROP COLUMN "classification", ADD COLUMN "score" "public"."meal_score" NOT NULL;
-- Drop enum type "meal_classification"
DROP TYPE "public"."meal_classification";
