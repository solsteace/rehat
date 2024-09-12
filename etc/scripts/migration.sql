DROP DATABASE IF EXISTS `gorest`;
CREATE DATABASE `gorest`;
USE `gorest`;

CREATE TABLE `users` (
    `user_id` INT PRIMARY KEY AUTO_INCREMENT,
    `email` VARCHAR(63) UNIQUE,
    `username` VARCHAR(31) UNIQUE,
    `name` VARCHAR(63),
    `password` VARBINARY(63),
    `role` ENUM("superadmin", "admin", "customer") DEFAULT "customer",
    `active` BOOLEAN DEFAULT TRUE
);

CREATE TABLE `motels` (
    `motel_id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(100),
    `location` VARCHAR(255),
    `contact_number` VARCHAR(15),
    `email` VARCHAR(100),
    `rating` INT(1)
);

CREATE TABLE `motel_admins` (
    `admin_id` INT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `motel_id` INT NOT NULL,

    FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`),
    FOREIGN KEY (`motel_id`) REFERENCES `motels`(`motel_id`)
);

CREATE TABLE  `classes`(
    `class_id` INT PRIMARY KEY AUTO_INCREMENT,
    `motel_id` INT,
    `class_name` VARCHAR(15),
    `price` INT,
    FOREIGN KEY (`motel_id`) REFERENCES `motels`(`motel_id`)
);

CREATE TABLE `rooms` (
    `room_id` INT PRIMARY KEY AUTO_INCREMENT,
    `class_id` INT,
    `room_number` INT,
    `status` ENUM("open", "reserved", "maintenance") DEFAULT "open",

    FOREIGN KEY (`class_id`) REFERENCES `classes`(`class_id`)
);

CREATE TABLE `reservation` (
    `reservation_id` INT PRIMARY KEY AUTO_INCREMENT,
    `room_id` INT,
    `user_id` INT,
    `reserve_start` DATETIME,
    `reserve_end` DATETIME,
    `checkout` DATETIME,
    `total` INT,

    FOREIGN KEY (`user_id`) REFERENCES `users`(`user_id`),
    FOREIGN KEY (`room_id`) REFERENCES `rooms`(`room_id`)
);