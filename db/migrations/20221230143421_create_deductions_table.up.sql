CREATE TABLE `deductions` (
    `id` mediumint unsigned AUTO_INCREMENT,
    `amount` int NOT NULL,
    `nominal` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`)
);