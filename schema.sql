CREATE TABLE "meals" (
    "id" TEXT PRIMARY KEY,
    "type" TEXT CHECK("type" IN ('breakfast', 'lunch', 'dinner')) NOT NULL,
    "chef" TEXT CHECK("chef" IN ('me', 'ordered', 'outsourced', 'work')) NOT NULL DEFAULT 'me',
    "date" DATE NOT NULL
);

CREATE TABLE "meal_dishes" (
    "id" TEXT PRIMARY KEY,
    "name" TEXT NOT NULL,
    "meal_id" TEXT NOT NULL,
    "score" TEXT CHECK("score" IN ('green', 'orange', 'cheat')) NOT NULL DEFAULT 'green',
    FOREIGN KEY ("meal_id") REFERENCES "meals" ("id") ON DELETE CASCADE
);

CREATE TABLE "kpi_types" (
    "id" TEXT PRIMARY KEY,
    "name" TEXT NOT NULL,
    "value_type" TEXT CHECK("value_type" IN ('sum', 'count')) NOT NULL DEFAULT 'sum'
);

CREATE TABLE "counters" (
    "id" TEXT PRIMARY KEY,
    "week_number" INTEGER NOT NULL,
    "kpi_type_id" TEXT NOT NULL,
    "value" INTEGER NOT NULL,
    FOREIGN KEY ("kpi_type_id") REFERENCES "kpi_types" ("id") ON DELETE CASCADE,
    UNIQUE ("week_number", "kpi_type_id")
);
