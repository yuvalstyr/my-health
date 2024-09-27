-- Create "meals" table
CREATE TABLE `meals` (`id` text NULL, `type` text NOT NULL, `chef` text NOT NULL DEFAULT 'me', `date` date NOT NULL, PRIMARY KEY (`id`), CHECK ("type" IN ('breakfast', 'lunch', 'dinner')), CHECK ("chef" IN ('me', 'ordered', 'outsourced', 'work')));
-- Create "meal_dishes" table
CREATE TABLE `meal_dishes` (`id` text NULL, `name` text NOT NULL, `meal_id` text NOT NULL, `score` text NOT NULL DEFAULT 'green', PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`meal_id`) REFERENCES `meals` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE, CHECK ("score" IN ('green', 'orange', 'cheat')));
-- Create "kpi_types" table
CREATE TABLE `kpi_types` (`id` text NULL, `name` text NOT NULL, `value_type` text NOT NULL DEFAULT 'sum', PRIMARY KEY (`id`), CHECK ("value_type" IN ('sum', 'count')));
-- Create "counters" table
CREATE TABLE `counters` (`id` text NULL, `week_number` integer NOT NULL, `kpi_type_id` text NOT NULL, `value` integer NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`kpi_type_id`) REFERENCES `kpi_types` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "counters_week_number_kpi_type_id" to table: "counters"
CREATE UNIQUE INDEX `counters_week_number_kpi_type_id` ON `counters` (`week_number`, `kpi_type_id`);
