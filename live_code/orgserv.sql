-- Active: 1736252549122@@127.0.0.1@3306@orgserv

-- Active: 1736250362930@@127.0.0.1@3306@users
show SCHEMAS;

DROP TABLE users;

CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id CHAR(36) NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  kind_id INT NOT NULL,
  #kind VARCHAR(50) UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE kinds (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE rooms (
  id INT AUTO_INCREMENT PRIMARY KEY,
  room VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE users_rooms (
  id INT AUTO_INCREMENT PRIMARY KEY,
  room_id INT NOT NULL,
  user_id INT NOT NULL
);

-- GIVEN: users table and rooms table. Create relation between them
-- TASK: find total number of rooms for each users
SELECT users.id, users.name, COUNT(ur.room_id) as user_rooms_cnt FROM users LEFT JOIN users_rooms as ur ON users.id = ur.user_id GROUP BY users.id

SHOW TABLES;

SELECT * FROM users;