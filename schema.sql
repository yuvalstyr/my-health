CREATE TYPE "meal_type" AS ENUM ('breakfast', 'lunch', 'dinner');
-- need a meal classification as green, orange and cheat_day
CREATE TYPE "meal_classification" AS ENUM ('green', 'orange', 'cheat');
CREATE TABLE "meals" (
    id SERIAL PRIMARY KEY,
    "type" meal_type NOT NULL,
    "date" DATE NOT NULL
);

CREATE TABLE "meal_dishes" (
    "id" SERIAL PRIMARY KEY,
    "name" text NOT NULL,
    "meal_id" integer NOT NULL,
    "classification" meal_classification NOT NULL,
    FOREIGN KEY ("meal_id") REFERENCES "meals" ("id") ON DELETE CASCADE
);

CREATE TABLE "activities" (
    "id" SERIAL PRIMARY KEY,
    "name" text NOT NULL,
    "date" DATE NOT NULL
);
