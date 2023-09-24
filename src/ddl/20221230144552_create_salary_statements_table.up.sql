CREATE TABLE `salary_statements` (
    `id` mediumint unsigned AUTO_INCREMENT,
    `earning_id` mediumint unsigned NOT NULL,
    `deduction_id` mediumint unsigned NOT NULL,
    `employee_id` mediumint unsigned NOT NULL,
    `nominal` varchar(255) NOT NULL,
    `payday` date NOT NULL,
    `target_period` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`earning_id`) REFERENCES `earnings` (`id`),
    FOREIGN KEY (`deduction_id`) REFERENCES `deductions` (`id`),
    FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`)
);