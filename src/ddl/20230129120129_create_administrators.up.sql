CREATE TABLE `administrators` (
    `id` mediumint unsigned AUTO_INCREMENT,
    `company_id` smallint unsigned NOT NULL,
    `name` varchar(255) NOT NULL,
    `mail_address` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`),
    UNIQUE(`mail_address`)
);