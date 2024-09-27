-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_kpi_types" table
CREATE TABLE `new_kpi_types` (`id` text NULL, `name` text NOT NULL, `value_type` text NOT NULL DEFAULT 'sum', `target` integer NOT NULL DEFAULT 0, PRIMARY KEY (`id`), CHECK ("value_type" IN ('sum', 'count')));
-- Copy rows from old table "kpi_types" to new temporary table "new_kpi_types"
INSERT INTO `new_kpi_types` (`id`, `name`, `value_type`, `target`) SELECT `id`, `name`, `value_type`, IFNULL(`target`, 0) AS `target` FROM `kpi_types`;
-- Drop "kpi_types" table after copying rows
DROP TABLE `kpi_types`;
-- Rename temporary table "new_kpi_types" to "kpi_types"
ALTER TABLE `new_kpi_types` RENAME TO `kpi_types`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
