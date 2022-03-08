/*
DROP TABLE IF EXISTS user_schema;
*/
DROP TABLE IF EXISTS list_schema;

CREATE TABLE list_schema(
  list_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  
  title VARCHAR(60) NOT NULL,
  
  user_id INT,
  FOREIGN KEY(user_id)
    REFERENCES user_schema(user_id)
      ON DELETE CASCADE
        ON UPDATE CASCADE
);

INSERT INTO list_schema (title, user_id)
VALUES ('Projeto Vibbra', 1);

INSERT INTO list_schema (title, user_id)
VALUES ('Entrar em contato com o sr vibbraneo', 2);