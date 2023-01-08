/* 
Reference : https://dev.mysql.com/doc/refman/8.0/ja/create-table.html
*/

CREATE TABLE `companies` (
    `id` smallint unsigned AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE(`name`)
);