CREATE DATABASE db_bisnis;

CREATE TABLE product (
    product_id VARCHAR(8) NOT NULL UNIQUE,
    product_name VARCHAR(100) NOT NULL,
    premium DECIMAL(17, 2) NOT NULL,
    active TINYINT(1) NOT NULL,
    create_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE agent (
    agent_id VARCHAR(8) NOT NULL UNIQUE,
    agent_name VARCHAR(100) NOT NULL,
    password VARCHAR(25) NOT NULL,
    active TINYINT(1) NOT NULL,
    create_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_parameter (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id VARCHAR(8) NOT NULL,
    parameter_name VARCHAR(100) NOT NULL,
    parameter_value VARCHAR(8) NOT NULL,
    active TINYINT(1) NOT NULL,
    create_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transaction (
    trans_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    agent_id VARCHAR(8) NOT NULL,
    product_id VARCHAR(8) NOT NULL,
    nama VARCHAR(100) NOT NULL,
    usia INT NOT NULL,
    premium DECIMAL(17,2) NOT NULL,
    create_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

