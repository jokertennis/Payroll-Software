CREATE TABLE `earnings` (
    `id` mediumint unsigned AUTO_INCREMENT,
    `nominal` varchar(255) NOT NULL,
    `amount` int NOT NULL,
    `earning_type` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`)
);