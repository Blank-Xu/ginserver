-- MySql scripts

CREATE DATABASE IF NOT EXISTS ginserver DEFAULT CHARACTER SET = utf8mb4 COLLATE utf8mb4_0900_ai_ci;

USE ginserver;

-- for mysql version 8.0-
-- CREATE USER IF NOT EXISTS ginserver IDENTIFIED BY '123456';
-- GRANT ALL PRIVILEGES ON ginserver.* TO 'ginserver'@'localhost' IDENTIFIED BY '123456';

-- for mysql version 8.0+
CREATE USER IF NOT EXISTS 'ginserver'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456';
GRANT ALL PRIVILEGES ON ginserver.* TO 'ginserver'@'localhost';

FLUSH PRIVILEGES;


-- table s_param
CREATE TABLE IF NOT EXISTS s_param
(
  id      int AUTO_INCREMENT PRIMARY KEY,
  ptype   tinyint(2)   NOT NULL COMMENT '0:string, 1:json',
  name    varchar(32)  NOT NULL,
  param   text         NOT NULL,
  remark  varchar(255) NOT NULL DEFAULT '',
  updater int          NOT NULL,
  updated timestamp    NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  UNIQUE uq_s_param_name (name)
) DEFAULT CHARSET = utf8mb4;


-- table log
CREATE TABLE IF NOT EXISTS log
(
  id      bigint AUTO_INCREMENT PRIMARY KEY,
  level   tinyint(2)   NOT NULL,
  type    tinyint(2)   NOT NULL,
  user_id int          NOT NULL COMMENT 'user id',
  role_id int          NOT NULL COMMENT 'user role',
  method  varchar(10)  NOT NULL COMMENT 'http method',
  path    varchar(50)  NOT NULL COMMENT 'router path',
  params  text         NOT NULL COMMENT 'request params',
  ip      varchar(32)  NOT NULL DEFAULT '' COMMENT 'request ip address',
  remark  varchar(255) NOT NULL DEFAULT '',
  created timestamp    NOT NULL DEFAULT current_timestamp,
  v1      varchar(255) NOT NULL DEFAULT '',
  v2      varchar(255) NOT NULL DEFAULT '',
  v3      varchar(255) NOT NULL DEFAULT '',
  INDEX idx_log_level (level),
  INDEX idx_log_type (type),
  INDEX idx_log_user_id (user_id),
  INDEX idx_log_role_id (role_id)
) DEFAULT CHARSET = utf8mb4;


-- table s_role_group
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
  DEFAULT CHARSET = utf8mb4;


-- table s_role
CREATE TABLE IF NOT EXISTS s_role
(
  id       int AUTO_INCREMENT PRIMARY KEY,
  group_id int          NOT NULL DEFAULT 0,
  name     varchar(32)  NOT NULL,
  state    tinyint(1)   NOT NULL DEFAULT 0 COMMENT '0:disable, 1:enable',
  remark   varchar(255) NOT NULL DEFAULT '',
  created  timestamp    NOT NULL DEFAULT current_timestamp,
  updater  int          NOT NULL,
  updated  timestamp    NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  UNIQUE uq_s_role_name (name)
) AUTO_INCREMENT = 100
  DEFAULT CHARSET = utf8mb4;

INSERT INTO s_role(name, state, updater)
VALUES ('admin', 1, 10000);


-- table s_menu
CREATE TABLE IF NOT EXISTS s_menu
(
  id          int AUTO_INCREMENT PRIMARY KEY,
  type        tinyint(2)   NOT NULL DEFAULT 0 COMMENT '0:main, 1:button, 2:href',
  name        varchar(64)  NOT NULL,
  description varchar(255) NOT NULL DEFAULT '',
  method      varchar(32)  NOT NULL COMMENT 'GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE',
  path        varchar(255) NOT NULL DEFAULT '',
  icon        varchar(255) NOT NULL DEFAULT '',
  level       tinyint(3)   NOT NULL DEFAULT 0,
  order_no    tinyint(4)   NOT NULL DEFAULT 1,
  state       tinyint(1)   NOT NULL DEFAULT 0 COMMENT '0:disable, 1:enable',
  parent_id   tinyint(4)   NOT NULL DEFAULT 0,
  created     timestamp    NOT NULL DEFAULT current_timestamp,
  updater     int          NOT NULL,
  updated     timestamp    NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  INDEX idx_s_menu_parent_id (parent_id),
  INDEX idx_s_menu_method (method),
  INDEX idx_s_menu_path (path)
) DEFAULT CHARSET = utf8mb4;

INSERT INTO s_menu(name, method, path, icon, level, order_no, state, parent_id, updater)
VALUES ('Home', 'GET', '/admin', '', 0, 1, 1, 0, 10000),
       ('About', 'GET', '/admin/about', '', 0, 5, 1, 0, 10000),
       ('User Set', 'MENU', '', '', 0, 2, 1, 0, 10000),
       ('Info', 'GET', '/admin/info', '', 1, 1, 1, 3, 10000),
       ('Change Password', 'GET', '/admin/change_pwd', '', 1, 2, 1, 3, 10000);


-- table s_role_menu
CREATE TABLE IF NOT EXISTS s_role_menu
(
  id      int AUTO_INCREMENT PRIMARY KEY,
  role_id int NOT NULL,
  menu_id int NOT NULL,
  INDEX idx_s_role_menu (role_id),
  UNIQUE uq_s_role_menu (role_id, menu_id)
) DEFAULT CHARSET = utf8mb4;

INSERT INTO s_role_menu(role_id, menu_id)
VALUES (100, 1),
       (100, 2),
       (100, 3),
       (100, 4),
       (100, 5);


-- table s_user
CREATE TABLE IF NOT EXISTS s_user
(
  id          int AUTO_INCREMENT PRIMARY KEY COMMENT 'user id',
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
  deleted     timestamp,
  register_ip varchar(64)  NOT NULL DEFAULT '',
  login_time  timestamp,
  login_ip    varchar(64)  NOT NULL DEFAULT '',
  UNIQUE uq_s_user_username (username)
) AUTO_INCREMENT = 10000
  DEFAULT CHARSET = utf8mb4;

# login password: 123456
# salt: dc83c7d015da92a93b0bd90144604d04
# salted password:  bfba91e771a65b4f0a10ba358d9c7655
INSERT INTO s_user (username, password, salt,
                    state, nickname, icon, email,
                    updater, register_ip, login_ip)
VALUES ('admin', 'bfba91e771a65b4f0a10ba358d9c7655', 'dc83c7d015da92a93b0bd90144604d04',
        1, 'blank', '/assets/img/avatar/avatar5.png', 'blank.xu@qq.com',
        10000, '127.0.0.1', '127.0.0.1');


-- table s_user_role
CREATE TABLE IF NOT EXISTS s_user_role
(
  id      int AUTO_INCREMENT PRIMARY KEY,
  user_id int NOT NULL,
  role_id int NOT NULL,
  INDEX idx_s_user_role (user_id),
  UNIQUE uq_s_user_role (user_id, role_id)
) DEFAULT CHARSET = utf8mb4;

INSERT INTO s_user_role(user_id, role_id)
VALUES (10000, 100);