DROP TABLE IF EXISTS list_child_schema;

CREATE TABLE list_child_schema(
  list_child_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  list_id INT,
  user_id INT,
  
  title VARCHAR(60) NOT NULL,
  description TEXT NOT NULL,
  checked BOOLEAN DEFAULT false,
  
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP DEFAULT NULL,

  FOREIGN KEY(list_id)
    REFERENCES list_schema(list_id)
      ON UPDATE CASCADE
        ON DELETE CASCADE,
  FOREIGN KEY(user_id)
    REFERENCES user_schema(user_id)
      ON UPDATE CASCADE
);

INSERT INTO list_child_schema (list_id, user_id, title, description)
VALUES (1, 2, 'Se hoje fosse amanhã que dia seria ontem?', 'Essa é minha dúvida');

INSERT INTO list_child_schema (list_id, user_id, title, description)
VALUES (1, 2, 'Como Hackear a NASA completo pt-BR', 'METODO INFALIVEL');

INSERT INTO list_child_schema (list_id, user_id, title, description)
VALUES (2, 1, 'Como fazer SQL Inject no MongoDB', 'Será que em 2022 isso ainda é possível?');