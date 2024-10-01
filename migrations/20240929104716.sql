-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_counters" table
CREATE TABLE `new_counters` (`id` text NULL, `week_number` integer NOT NULL, `kpi_type_id` text NOT NULL, `value` integer NOT NULL, `target` integer NOT NULL DEFAULT 0, `icon` text NOT NULL DEFAULT 'burger', PRIMARY KEY (`id`), CONSTRAINT `0` FOREIGN KEY (`kpi_type_id`) REFERENCES `kpi_types` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE);
-- Copy rows from old table "counters" to new temporary table "new_counters"
INSERT INTO `new_counters` (`id`, `week_number`, `kpi_type_id`, `value`) SELECT `id`, `week_number`, `kpi_type_id`, `value` FROM `counters`;
-- Drop "counters" table after copying rows
DROP TABLE `counters`;
-- Rename temporary table "new_counters" to "counters"
ALTER TABLE `new_counters` RENAME TO `counters`;
-- Create index "counters_week_number_kpi_type_id" to table: "counters"
CREATE UNIQUE INDEX `counters_week_number_kpi_type_id` ON `counters` (`week_number`, `kpi_type_id`);
-- Create "new_kpi_types" table
CREATE TABLE `new_kpi_types` (`id` text NULL, `name` text NOT NULL, `value_type` text NOT NULL DEFAULT 'sum', PRIMARY KEY (`id`), CHECK ("value_type" IN ('sum', 'count')));
-- Copy rows from old table "kpi_types" to new temporary table "new_kpi_types"
INSERT INTO `new_kpi_types` (`id`, `name`, `value_type`) SELECT `id`, `name`, `value_type` FROM `kpi_types`;
-- Drop "kpi_types" table after copying rows
DROP TABLE `kpi_types`;
-- Rename temporary table "new_kpi_types" to "kpi_types"
ALTER TABLE `new_kpi_types` RENAME TO `kpi_types`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
