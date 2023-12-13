-- create and use db for order before execute this script
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`
(
    `id`          VARCHAR(36)  NOT NULL,
    `username`    VARCHAR(255) NOT NULL,
    `price`       FLOAT        NOT NULL,
    `status`      VARCHAR(36)  NOT NULL,
    `created_by`  VARCHAR(36)  NOT NULL,
    `created_at`  INT(11)      NOT NULL,
    `modified_by` VARCHAR(36) DEFAULT NULL,
    `modified_at` INT(11)     DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = `latin1`;


-- create and use db for payment before execute this script
DROP TABLE IF EXISTS `payment`;
CREATE TABLE `payment`
(
    `id`          VARCHAR(36) NOT NULL,
    `order_Id`    VARCHAR(36) NOT NULL,
    `price`       FLOAT       NOT NULL,
    `status`      VARCHAR(36) NOT NULL,
    `created_by`  VARCHAR(36) NOT NULL,
    `created_at`  INT(11)     NOT NULL,
    `modified_by` VARCHAR(36) DEFAULT NULL,
    `modified_at` INT(11)     DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = `latin1`;