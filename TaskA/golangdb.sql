DROP DATABASE IF EXISTS golangdb;
CREATE DATABASE golangdb;
USE golangdb;

CREATE TABLE person (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  age INT
);

CREATE TABLE phone (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  number VARCHAR(255),
  person_id INT,
  FOREIGN KEY (person_id) REFERENCES person(id)
);

CREATE TABLE address (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  city VARCHAR(255),
  state VARCHAR(255),
  street1 VARCHAR(255),
  street2 VARCHAR(255),
  zip_code VARCHAR(255)
);

CREATE TABLE address_join (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  person_id INT,
  address_id INT,
  FOREIGN KEY (person_id) REFERENCES person(id),
  FOREIGN KEY (address_id) REFERENCES address(id)
);

INSERT INTO person (id, name, age) VALUES
(1, 'mike', 31),
(2, 'John', 45),
(3, 'Joseph', 20);

INSERT INTO phone (id, number, person_id) VALUES
(1, '444-444-4444', 1),
(8, '123-444-7777', 2),
(3, '445-444-4444', 3);

INSERT INTO address (id, city, state, street1, street2, zip_code) VALUES
(1, 'Eugene', 'OR', '111 Main St', '', '98765'),
(2, 'Sacramento', 'CA', '432 First St', 'Apt 1', '22221'),
(3, 'Austin', 'TX', '213 South 1st St', NULL, '78704');

INSERT INTO address_join (id, person_id, address_id) VALUES
(1, 1, 1),
(2, 2, 1),  
(3, 3, 2);  

