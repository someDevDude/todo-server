CREATE DATABASE IF NOT EXISTS todolist;
USE todolist;

CREATE USER IF NOT EXISTS 'user'@% IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON todolist TO 'user'@'localhost';

CREATE TABLE `todo` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `TITLE` varchar(100) NOT NULL,
  `DESCRIPTION` varchar(400) DEFAULT NULL,
  `DONE` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;