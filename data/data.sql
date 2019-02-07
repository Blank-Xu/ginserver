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
  ptype   tinyint(2)   NOT NULL COMMENT '0:string, 1:json',
  name    varchar(32)  NOT NULL,
  param   text         NOT NULL,
  remark  varchar(255) NOT NULL DEFAULT '',
  updater int          NOT NULL,
  updated timestamp    NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  UNIQUE uq_s_params_name (name)
) DEFAULT CHARACTER SET = utf8mb4;

-- admin log
CREATE TABLE IF NOT EXISTS s_log
(
  id       bigint AUTO_INCREMENT PRIMARY KEY,
  log_type tinyint(2)   NOT NULL,
  user_id  int          NOT NULL COMMENT 'user id',
  role_id  int          NOT NULL COMMENT 'user role',
  method   varchar(10)  NOT NULL COMMENT 'http method',
  path     varchar(50)  NOT NULL COMMENT 'router path',
  params   text         NOT NULL COMMENT 'request params',
  ip       varchar(32)  NOT NULL DEFAULT '' COMMENT 'request ip address',
  remark   varchar(255) NOT NULL DEFAULT '',
  created  timestamp    NOT NULL DEFAULT current_timestamp,
  v1       varchar(255) NOT NULL DEFAULT '',
  v2       varchar(255) NOT NULL DEFAULT '',
  v3       varchar(255) NOT NULL DEFAULT '',
  INDEX idx_s_log_user_id (user_id),
  INDEX idx_s_log_role_id (role_id)
) DEFAULT CHARACTER SET = utf8mb4;

CREATE TABLE IF NOT EXISTS s_role_group
(
  id      int AUTO_INCREMENT PRIMARY KEY,
  name    varchar(32)  NOT NULL,
  state   tinyint(1)   NOT NULL DEFAULT 0 COMMENT '0:disable, 1:enable',
  remark  varchar(255) NOT NULL DEFAULT '',
  created timestamp    NOT NULL DEFAULT current_timestamp,
  updater int          NOT NULL,
  updated timestamp    NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
) AUTO_INCREMENT = 100
  DEFAULT CHARACTER SET = utf8mb4;

CREATE TABLE IF NOT EXISTS s_role
(
  id       int AUTO_INCREMENT PRIMARY KEY,
  group_id int          NOT NULL DEFAULT 0,
  name     varchar(32)  NOT NULL,
  state    tinyint(1)   NOT NULL DEFAULT 0 COMMENT '0:disable, 1:enable',
  remark   varchar(255) NOT NULL DEFAULT '',
  created  timestamp    NOT NULL DEFAULT current_timestamp,
  updater  int          NOT NULL,
  updated  timestamp    NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
) AUTO_INCREMENT = 100
  DEFAULT CHARACTER SET = utf8mb4;

CREATE TABLE IF NOT EXISTS s_menu
(
  id       int AUTO_INCREMENT PRIMARY KEY,
  name     varchar(64)  NOT NULL,
  state    tinyint(1)   NOT NULL DEFAULT 0 COMMENT '0:disable, 1:enable',
  order_no tinyint(3)   NOT NULL DEFAULT 0,
  icon     varchar(255) NOT NULL DEFAULT '',
  method   varchar(32)  NOT NULL DEFAULT '',
  path     varchar(255) NOT NULL DEFAULT '',
  created  timestamp    NOT NULL DEFAULT current_timestamp,
  updater  int          NOT NULL,
  updated  timestamp    NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
) DEFAULT CHARACTER SET = utf8mb4;

-- admin users
CREATE TABLE IF NOT EXISTS s_admin
(
  id          int AUTO_INCREMENT PRIMARY KEY COMMENT 'admin id',
  role_id     int          NOT NULL DEFAULT 0,
  username    varchar(32)  NOT NULL,
  password    varchar(32)  NOT NULL,
  salt        varchar(32)  NOT NULL COMMENT 'salt for encrypt password',
  state       tinyint(1)   NOT NULL DEFAULT 0 COMMENT 'user state, 0:disable, 1:enable',
  nickname    varchar(32)  NOT NULL DEFAULT '',
  icon        varchar(64)  NOT NULL DEFAULT '',
  email       varchar(64)  NOT NULL DEFAULT '',
  phone       varchar(32)  NOT NULL DEFAULT '',
  remark      varchar(255) NOT NULL DEFAULT '',
  created     timestamp    NOT NULL DEFAULT current_timestamp,
  updater     int          NOT NULL,
  updated     timestamp    NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  deleted     int          NOT NULL DEFAULT 0,
  register_ip varchar(64)  NOT NULL DEFAULT '',
  login_time  timestamp,
  login_ip    varchar(64)  NOT NULL DEFAULT '',
  UNIQUE uq_s_admin_username (username)
) AUTO_INCREMENT = 10000
  DEFAULT CHARACTER SET = utf8mb4;

# login password: 123456
# salt: dc83c7d015da92a93b0bd90144604d04
# salted password:  bfba91e771a65b4f0a10ba358d9c7655
INSERT INTO s_admin (username, password, salt,
                     state, nickname, icon, email,
                     updater, register_ip, login_ip)
VALUES ('admin', 'bfba91e771a65b4f0a10ba358d9c7655', 'dc83c7d015da92a93b0bd90144604d04',
        1, 'blank', 'statics/img/avatar/avatar5.png', 'blank.xu@qq.com',
        10000, '127.0.0.1', '127.0.0.1');


