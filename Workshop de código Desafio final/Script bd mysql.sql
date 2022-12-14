select * from dentists ;
select * from patients ;
select * from appointments ;
SELECT * FROM patients WHERE id=1;

CREATE DATABASE  IF NOT EXISTS `my_db`;

USE `my_db`;

DROP TABLE IF EXISTS `dentists`;

CREATE TABLE `dentists` (
	`id` 			int AUTO_INCREMENT PRIMARY KEY NOT NULL,
    `firstname` 	varchar(50) DEFAULT NULL,
    `lastname`      varchar(50) DEFAULT NULL,
    `registration`  varchar(50) DEFAULT NULL
   );

DROP TABLE IF EXISTS `patients`;

CREATE TABLE `patients` (
	`id` 				int AUTO_INCREMENT PRIMARY KEY NOT NULL,
    `firstname` 		varchar(50) DEFAULT NULL,
    `lastname`			varchar(50) DEFAULT NULL,
    `rg` 				varchar(50) DEFAULT NULL,
    `registrationDate`  Date DEFAULT NULL
   );

DROP TABLE IF EXISTS `appointments`;

CREATE TABLE `appointments` (
	`id` 				int AUTO_INCREMENT PRIMARY KEY  NOT  NULL,
    `id_pacient` 		int DEFAULT NULL,
    `id_dentist`		int DEFAULT NULL,
    `datetime` 			DATETIME DEFAULT NULL,
    `description`  		text DEFAULT NULL,
    FOREIGN KEY (id_pacient) REFERENCES patients(id),
    FOREIGN KEY (id_dentist) REFERENCES dentists(id)
   );

