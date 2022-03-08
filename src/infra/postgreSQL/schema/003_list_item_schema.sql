DROP TABLE IF EXISTS list_item_schema;

CREATE TABLE list_item_schema(
  list_item_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  list_id INT,
  user_id INT,
  
  title VARCHAR(60) NOT NULL,
  description TEXT NOT NULL,
  
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,

  FOREIGN KEY(list_id)
    REFERENCES list_schema(list_id)
      ON UPDATE CASCADE
        ON DELETE CASCADE,
  FOREIGN KEY(user_id)
    REFERENCES user_schema(user_id)
      ON UPDATE CASCADE
);

INSERT INTO list_item_schema (list_id, user_id, title, description)
VALUES (1, 1, 'Começar o quanto antes', 'Organizar o projeto por completo');


INSERT INTO list_item_schema (list_id, user_id, title, description)
VALUES (2, 1, 'Primeiro perguntar sobre o TODO list', 'Olá sr vibbraneo...');