CREATE TABLE `earning_details` (
    `id` mediumint unsigned AUTO_INCREMENT,
    `earning_id` mediumint unsigned NOT NULL,
    `nominal` varchar(255) NOT NULL,
    `amount` int NOT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`earning_id`) REFERENCES `earnings` (`id`)
);