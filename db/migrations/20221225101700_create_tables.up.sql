CREATE TABLE IF NOT EXISTS companies
(
    id unsigned AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    PRIMARY KEY (id),
    UNIQUE(name)
);

CREATE TABLE IF NOT EXISTS employees
(
    id unsigned AUTO_INCREMENT,
    company_id unsigned NOT NULL,
    name varchar(255) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY(company_id) REFERENCES companies (id)
);