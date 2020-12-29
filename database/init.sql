CREATE DATABASE IF NOT EXISTS example_db;
USE example_db;
CREATE TABLE IF NOT EXISTS `users` (
  `user_id` varchar(50) NOT NULL,
  `mobile_no` varchar(12) DEFAULT NULL,
  `first_name` varchar(500) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `idx_mobile_no` (`mobile_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;