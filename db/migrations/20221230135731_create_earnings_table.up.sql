CREATE TABLE `earnings` (
    `id` mediumint unsigned AUTO_INCREMENT,
    `amount` int NOT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`)
);