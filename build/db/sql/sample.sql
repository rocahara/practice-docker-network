DROP DATABASE IF EXISTS sample;
CREATE DATABASE sample;
USE sample;

DROP TABLE IF EXISTS UserInfo;

CREATE TABLE UserInfo (
    Id INT(10),
    Name VARCHAR(40)
);

INSERT INTO UserInfo (Id, Name) VALUES (1, "ryuichi");
INSERT INTO UserInfo (Id, Name) VALUES (2, "hiroshi");
INSERT INTO UserInfo (Id, Name) VALUES (3, "mamoru");

CREATE USER testUser@localhost IDENTIFIED BY 'test';
GRANT ALL PRIVILEGES ON sample.* TO testUser@localhost;

CREATE USER testUser2@localhost IDENTIFIED BY 'test2';
GRANT ALL PRIVILEGES ON sample.* TO testUser2@localhost;