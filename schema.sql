CREATE TABLE `tbl` (
  `id` BIGINT UNSIGNED NOT NULL,
  `is_active` BOOLEAN NOT NULL,
  PRIMARY KEY (`id`),
  KEY `is_active` (`is_active`)
) ENGINE=InnoDB;
