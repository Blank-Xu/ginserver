CREATE SCHEMA IF NOT EXISTS ginserver DEFAULT CHARSET utf8mb4;

USE ginserver;

CREATE USER IF NOT EXISTS ginserver IDENTIFIED BY '123456';

GRANT ALL PRIVILEGES ON ginserver TO 'ginserver'@'127.0.0.1:3306' IDENTIFIED BY '123456';
FLUSH PRIVILEGES;

CREATE TABLE IF NOT EXISTS s_admin
(
  id          int AUTO_INCREMENT PRIMARY KEY,
  username    varchar(20)  NOT NULL,
  password    varchar(32)  NOT NULL,
  salt        varchar(32)  NOT NULL COMMENT 'salt for encrypt password',
  email       varchar(40)  NOT NULL DEFAULT '',
  phone       varchar(20)  NOT NULL DEFAULT '',
  remark      varchar(255) NOT NULL DEFAULT '',
  creat_time  timestamp             DEFAULT CURRENT_TIMESTAMP,
  update_time timestamp             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  login_time  timestamp,
  login_ip    varchar(60)  NOT NULL DEFAULT '',
  UNIQUE uq_s_admin_name (username)
);