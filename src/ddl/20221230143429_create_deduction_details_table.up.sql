CREATE TABLE `deduction_details` (
    `id` mediumint unsigned AUTO_INCREMENT,
    `deduction_id` mediumint unsigned NOT NULL,
    `nominal` varchar(255) NOT NULL,
    `amount` int NOT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`deduction_id`) REFERENCES `deductions` (`id`)
);