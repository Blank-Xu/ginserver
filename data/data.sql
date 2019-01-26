-- MySql scripts

CREATE DATABASE IF NOT EXISTS ginserver DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

USE ginserver;

-- for mysql version 8.0-
-- CREATE USER IF NOT EXISTS ginserver IDENTIFIED BY '123456';
-- GRANT ALL PRIVILEGES ON ginserver.* TO 'ginserver'@'localhost' IDENTIFIED BY '123456';

-- for mysql version 8.0+
CREATE USER IF NOT EXISTS 'ginserver'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456';
GRANT ALL PRIVILEGES ON ginserver.* TO 'ginserver'@'localhost';

FLUSH PRIVILEGES;

-- system params
CREATE TABLE IF NOT EXISTS s_params
(
  id      int AUTO_INCREMENT PRIMARY KEY,
  name    varchar(32)  NOT NULL,
  param   text         NOT NULL,
  remark  varchar(255) NOT NULL DEFAULT '',
  updated timestamp    NOT NULL,
  UNIQUE uq_s_params_name (name)
) DEFAULT CHARACTER SET = utf8mb4;

-- admin log
CREATE TABLE IF NOT EXISTS s_log
(
  id      bigint AUTO_INCREMENT PRIMARY KEY,
  uid     int          NOT NULL COMMENT 'user id',
  role    varchar(32)  NOT NULL COMMENT 'user role',
  method  varchar(10)  NOT NULL COMMENT 'http method',
  router  varchar(50)  NOT NULL COMMENT 'http router',
  params  text         NOT NULL COMMENT 'request params',
  created timestamp    NOT NULL,
  ip      varchar(32)  NOT NULL DEFAULT '' COMMENT 'request ip address',
  remark  varchar(255) NOT NULL DEFAULT '',
  INDEX idx_s_log_uid (uid)
) DEFAULT CHARACTER SET = utf8mb4;

-- admin users
CREATE TABLE IF NOT EXISTS s_admin
(
  id         int AUTO_INCREMENT PRIMARY KEY COMMENT 'user id',
  username   varchar(32)  NOT NULL,
  password   varchar(32)  NOT NULL,
  salt       varchar(32)  NOT NULL COMMENT 'salt for encrypt password',
  nickname   varchar(32)  NOT NULL DEFAULT '',
  email      varchar(40)  NOT NULL DEFAULT '',
  phone      varchar(20)  NOT NULL DEFAULT '',
  remark     varchar(255) NOT NULL DEFAULT '',
  created    timestamp    NOT NULL,
  updated    timestamp    NOT NULL,
  login_time timestamp,
  login_ip   varchar(60)  NOT NULL DEFAULT '',
  UNIQUE uq_s_admin_username (username)
) DEFAULT CHARACTER SET = utf8mb4;