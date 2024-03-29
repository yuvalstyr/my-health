CREATE TYPE "meal_type" AS ENUM ('breakfast', 'lunch', 'dinner');
CREATE TYPE "meal_score" AS ENUM ('green', 'orange', 'cheat');
CREATE TYPE "chefs" AS ENUM ('me','ordered','outsourced','work');
CREATE TYPE "value_type" AS ENUM ('sum','count');

CREATE TABLE "meals" (
    "id" text PRIMARY KEY,
    "type" meal_type  NOT NULL,
    "chef" chefs NOT NULL DEFAULT 'me',
    "date" DATE NOT NULL
);

CREATE TABLE "meal_dishes" (
    "id" text PRIMARY KEY,
    "name" text NOT NULL,
    "meal_id" text NOT NULL,
    "score" meal_score NOT NULL DEFAULT 'green',
    FOREIGN KEY ("meal_id") REFERENCES "meals" ("id") ON DELETE CASCADE
);

CREATE TABLE "activity_types" (
    "id" text PRIMARY KEY,
    "name" text NOT NULL,
    "value_type" value_type NOT NULL default 'sum',
);

CREATE TABLE "activities" (
  "id" text,
  "date" DATE NOT NULL,
  "type_id" text NOT NULL,
  "value" integer NOT NULL,
  PRIMARY KEY ("id"),
  FOREIGN KEY ("type_id") REFERENCES "activity_types" ("id") ON DELETE CASCADE,
  UNIQUE (date,type_id)
);
