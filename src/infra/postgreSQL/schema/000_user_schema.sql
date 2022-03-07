/*
  Se quiser deletar essa tabela primeiro vai ter
  que deletar as outras tabelas primeiro
*/

DROP TABLE IF EXISTS user_schema;

CREATE TABLE user_schema(
  user_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  
  username VARCHAR(30) UNIQUE NOT NULL,
  fullname VARCHAR(30) NOT NULL,
  email TEXT UNIQUE NOT NULL,
  passwd VARCHAR(20) NOT NULL,
  
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO user_schema (username, fullname, email, passwd)
VALUES ('Hamilton', 'José Hamilton', 'hamilton@gmail.com', '12345');

INSERT INTO user_schema (username, fullname, email, passwd)
VALUES ('jose', 'Luiz Ricardo', 'jose@gmail.com', '123456');