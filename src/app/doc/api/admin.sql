CREATE TABLE roles (
                       id INT PRIMARY KEY AUTO_INCREMENT,
                       name VARCHAR(50) NOT NULL
);

CREATE TABLE permissions (
                             id INT PRIMARY KEY AUTO_INCREMENT,
                             name VARCHAR(50) NOT NULL
);

CREATE TABLE role_permissions (
                                  id INT PRIMARY KEY AUTO_INCREMENT,
                                  role_id INT,
                                  permission_id INT,
                                  FOREIGN KEY (role_id) REFERENCES roles(id),
                                  FOREIGN KEY (permission_id) REFERENCES permissions(id)
);

CREATE TABLE users (
                       id INT PRIMARY KEY AUTO_INCREMENT,
                       username VARCHAR(50) NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       role_id INT,
                       FOREIGN KEY (role_id) REFERENCES roles(id)
);
