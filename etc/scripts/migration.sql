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
    `verified` BOOLEAN DEFAULT FALSE,
    `deleted_at` DATETIME DEFAULT NULL
);

CREATE TABLE `motels` (
    `motel_id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(100),
    `location` VARCHAR(255),
    `contact_number` VARCHAR(15),
    `email` VARCHAR(100)
);

CREATE TABLE `motel_admins` (
    `admin_id` INT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT,
    `motel_id` INT,

    FOREIGN KEY (`user_id`) 
        REFERENCES `users`(`user_id`) 
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (`motel_id`) 
        REFERENCES `motels`(`motel_id`) 
        ON UPDATE CASCADE
        ON DELETE SET NULL
);


CREATE TABLE  `room_classes`(
    `class_id` INT PRIMARY KEY AUTO_INCREMENT,
    `motel_id` INT,
    `name` VARCHAR(15),
    `price` INT,

    FOREIGN KEY (`motel_id`) 
        REFERENCES `motels`(`motel_id`)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE `rooms` (
    `room_id` INT PRIMARY KEY AUTO_INCREMENT,
    `motel_id` INT,
    `class_id` INT,
    `room_number` INT,
    `status` ENUM("open", "reserved", "maintenance") DEFAULT "open",

    FOREIGN KEY (`class_id`) 
        REFERENCES `room_classes`(`class_id`)
        ON UPDATE CASCADE
        ON DELETE SET NULL,
    FOREIGN KEY (`motel_id`) 
        REFERENCES `motels`(`motel_id`)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE `reservations` (
    `reservation_id` INT PRIMARY KEY AUTO_INCREMENT,
    `room_id` INT,
    `user_id` INT,
    `reserve_start` DATETIME,
    `reserve_end` DATETIME,
    `checkout` DATETIME,
    `total` INT,

    FOREIGN KEY (`user_id`) 
        REFERENCES `users`(`user_id`)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    FOREIGN KEY (`room_id`) 
        REFERENCES `rooms`(`room_id`)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);