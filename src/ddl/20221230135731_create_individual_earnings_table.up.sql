CREATE TABLE `individual_earnings` (
    `id` mediumint unsigned AUTO_INCREMENT,
    `nominal` varchar(255) NOT NULL,
    `amount` int NOT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`)
);