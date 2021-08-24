-- 创建用户信息表

DROP TABLE t_user_info;

CREATE TABLE t_user_info (
     Id INT NOT NULL AUTO_INCREMENT PRIMARY KEY ,
     Name VARCHAR(64) NOT NULL,
     Phone VARCHAR(12) NOT NULL,
     Status INT NOT NULL,
     Procdate DATE
)