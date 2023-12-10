CREATE TYPE "meal_type" AS ENUM ('breakfast', 'lunch', 'dinner');
-- need a meal classification as green, orange and cheat_day
CREATE TYPE "meal_classification" AS ENUM ('green', 'orange', 'cheat_day');
CREATE TABLE "meals" (
    "id" integer PRIMARY KEY AUTOINCREMENT NOT NULL,
    "type" meal_type NOT NULL,
    "classification" meal_classification NOT NULL,
    "date" datetime NOT NULL
)

CREATE TABLE "meal_items" (
    "id" integer PRIMARY KEY AUTOINCREMENT NOT NULL,
    "name" text NOT NULL,
    "meal_id" integer NOT NULL,
    FOREIGN KEY ("meal_id") REFERENCES "meals" ("id") ON DELETE CASCADE
)

CREATE TABLE "activities" (
    "id" integer PRIMARY KEY AUTOINCREMENT NOT NULL,
    "name" text NOT NULL
    "date" datetime NOT NULL
)