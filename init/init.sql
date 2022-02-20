use `project-capstone`;

CREATE TABLE IF NOT EXISTS `roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(255) NOT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(1000) NOT NULL,
  `divisi` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `id_role` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `users_name` (`name`),
  KEY `users_FK` (`id_role`),
  CONSTRAINT `users_FK` FOREIGN KEY (`id_role`) REFERENCES `roles` (`id`)
);

CREATE TABLE IF NOT EXISTS `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(255) NOT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `status_check` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(255) NOT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `assets` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_category` int NOT NULL,
  `is_maintenance` BOOL DEFAULT false,
  `name` varchar(255) NOT NULL,
  `description` text DEFAULT NULL,
  `initial_quantity` int NOT NULL,
  `avail_quantity` int NOT NULL,
  `photo` varchar(1000) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `assets_FK` FOREIGN KEY (`id_category`) REFERENCES `categories` (`id`)
);

CREATE TABLE IF NOT EXISTS `requests` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_user` int NOT NULL,
  `id_asset` int NOT NULL,
  `id_status` int NOT NULL,
  `request_date` datetime DEFAULT NULL,
  `return_date` datetime DEFAULT NULL,
  `description` text DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `updated_date` datetime DEFAULT NULL,
  `deleted_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `requests_users_FK` FOREIGN KEY (`id_user`) REFERENCES `users` (`id`),
  CONSTRAINT `requests_assets_FK` FOREIGN KEY (`id_asset`) REFERENCES `assets` (`id`),
  CONSTRAINT `requests_status_FK` FOREIGN KEY (`id_status`) REFERENCES `status_check` (`id`)
);

