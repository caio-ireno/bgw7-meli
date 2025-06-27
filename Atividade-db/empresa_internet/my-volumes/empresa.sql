SET NAMES 'utf8mb4';

DROP DATABASE IF EXISTS empresa_internet;
CREATE DATABASE empresa_internet CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE empresa_internet;

DROP TABLE IF EXISTS `internet`;

CREATE TABLE internet (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `velocity` int  NOT NULL, 
  `price` float  NOT NULL, 
  `discount` float  NOT NULL,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `customers`;
CREATE TABLE customers (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100)  NOT NULL,
  `last_name` varchar(100)  NOT NULL,
  `province` varchar(100)  NOT NULL,
  `city` varchar(100)  NOT NULL,
  `date_birthday` date NOT NULL,
  `plan_id` int(10) unsigned,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`plan_id`) REFERENCES `internet`(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table `internet`
INSERT INTO `internet` (`id`, `velocity`, `price`, `discount`) VALUES
(1, 100, 99.90, 10.0),
(2, 200, 149.90, 15.0),
(3, 300, 199.90, 20.0);


-- Dumping data for table `customers`

INSERT INTO `customers` (`id`, `name`, `last_name`, `province`, `city`, `date_birthday`, `plan_id`) VALUES
(1, 'João', 'Silva', 'MG', 'Belo Horizonte', '1990-01-01', 1),
(2, 'Maria', 'Santos', 'SP', 'São Paulo', '1985-03-15', 1),
(3, 'Carlos', 'Oliveira', 'RJ', 'Rio de Janeiro', '1978-11-22', 2);

