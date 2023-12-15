CREATE TYPE "meal_type" AS ENUM ('breakfast', 'lunch', 'dinner');
CREATE TYPE "meal_score" AS ENUM ('green', 'orange', 'cheat');
CREATE TYPE "chefs" AS ENUM ('me','ordered','outsourced');

CREATE TABLE "meals" (
    id SERIAL PRIMARY KEY,
    "type" meal_type  NOT NULL,
    "chef" chefs NOT NULL DEFAULT 'me',
    "date" DATE NOT NULL
);

CREATE TABLE "meal_dishes" (
    "id" SERIAL PRIMARY KEY,
    "name" text NOT NULL,
    "meal_id" integer NOT NULL,
    "score" meal_score NOT NULL DEFAULT 'green',
    FOREIGN KEY ("meal_id") REFERENCES "meals" ("id") ON DELETE CASCADE
);

CREATE TABLE "activities" (
    "id" SERIAL PRIMARY KEY,
    "name" text NOT NULL,
    "date" DATE NOT NULL
);
