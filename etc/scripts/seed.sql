CREATE DATABASE IF NOT EXISTS `gorest`;
USE `gorest`;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `user_id` INT PRIMARY KEY AUTO_INCREMENT,
    `email` VARCHAR(63) UNIQUE,
    `username` VARCHAR(31) UNIQUE,
    `name` VARCHAR(63),
    `password` VARBINARY(63),
    `role` ENUM("admin", "staff", "customer") DEFAULT "customer"
);


DROP TABLE IF EXISTS `motels`;
CREATE TABLE `motels` (
    `motel_id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(100),
    `location` VARCHAR(255),
    `contact_number` VARCHAR(15),
    `email` VARCHAR(100)
);
INSERT INTO `motels` (`name`, `location`, `contact_number`, `email`)
    VALUES
        ('Sunset Motel', '123 Beach Rd, Ocean City', '555-123-4567', 'contact@sunsetmotel.com'),
        ('Mountain View Inn', '456 Hilltop Ave, Mountainville', '555-234-5678', 'info@mountainviewinn.com'),
        ('City Central Hotel', '789 Downtown St, Metropolis', '555-345-6789', 'reservations@citycentralhotel.com'),
        ('Seaside Retreat', '321 Ocean Blvd, Coastal Town', '555-456-7890', 'info@seasideretreat.com'),
        ('Lakeside Lodge', '654 Lake Rd, Lakeside', '555-567-8901', 'contact@lakesidelodge.com'),
        ('Countryside Inn', '987 Country Ln, Ruralville', '555-678-9012', 'reservations@countrysideinn.com'),
        ('Urban Stay', '111 Metro St, Big City', '555-789-0123', 'info@urbanstay.com'),
        ('Historic Hotel', '222 Heritage Ave, Oldtown', '555-890-1234', 'contact@historichotel.com');
