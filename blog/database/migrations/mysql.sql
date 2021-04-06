-- CREATE DATABASE
CREATE DATABASE IF NOT EXISTS blog;

-- CREATE TABLES

-- USERS TABLE
DROP TABLE IF EXISTS `users`;

CREATE TABLE IF NOT EXISTS `users` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NULL,
    `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `password` varchar(255) COLLATE utf8mb4_unicode_ci NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- CATEGORIES TABLE
DROP TABLE IF EXISTS `categories`;

CREATE TABLE IF NOT EXISTS `categories` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `is_active`	tinyint(4) NOT NULL DEFAULT 0,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- POSTS TABLE
DROP TABLE IF EXISTS `posts`;

CREATE TABLE IF NOT EXISTS `posts` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `content` text COLLATE utf8mb4_unicode_ci NOT NULL,
    `category_id` bigint(20) unsigned NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;