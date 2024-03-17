CREATE DATABASE IF NOT EXISTS usuarios;
USE usuarios;

CREATE TABLE IF NOT EXISTS cliente  (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL
);

INSERT INTO cliente VALUES(1,'Jarvic Balanta', 'jarvicf@gmail.com')
