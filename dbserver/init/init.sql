DROP TABLE IF EXISTS todos;

CREATE TABLE todos
(
    id INT PRIMARY KEY AUTO_INCREMENT,
    content TEXT DEFAULT NULL,
    created_at datetime DEFAULT NULL,
    updated_at datetime DEFAULT NULL,
    deleted_at datetime DEFAULT NULL
);

INSERT INTO todos (content) VALUES ("Nagaoka");
INSERT INTO todos (content) VALUES ("Tanaka");
