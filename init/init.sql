CREATE DATABASE IF NOT EXISTS `capstone_project`;

CREATE TABLE IF NOT EXISTS `capstone_project`.`users`(
	`id` int NOT NULL AUTO_INCREMENT,
	`name` varchar(255) NOT NULL,
	`email` varchar(255) NOT NULL,
	`password` varchar(1000) NOT NULL,
	`birth_date` varchar(255),
	`phone_number` varchar(50),
	`photo` varchar(1000),
	`gender` varchar(10),
	`address` varchar(255),
	`created_at` DATETIME DEFAULT NULL,
	`updated_at` DATETIME DEFAULT NULL,
	`deleted_at` DATETIME DEFAULT NULL,
	PRIMARY KEY(`id`)
);