DROP TABLE IF EXISTS user_schema;

CREATE TABLE user_schema(
  user_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  
  username VARCHAR(30) UNIQUE NOT NULL,
  email TEXT UNIQUE NOT NULL,
  passwd VARCHAR(20) NOT NULL,
  
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO user_schema (username, email, passwd)
VALUES ('Hamilton', 'hamilton@gmail.com', '12345');

INSERT INTO user_schema (username, email, passwd)
VALUES ('jose', 'jose@gmail.com', '123456');