CREATE TABLE `fixed_deduction_details` (
    `id` mediumint unsigned AUTO_INCREMENT,
    `fixed_deduction_id` mediumint unsigned NOT NULL,
    `nominal` varchar(255) NOT NULL,
    `amount` int NOT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`fixed_deduction_id`) REFERENCES `fixed_deductions` (`id`)
);