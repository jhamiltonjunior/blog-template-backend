# You simply need insert in your terminal this code:

# Em um projeto real

Eu acredito que em um projeto real, eu mesmo iria hospedar o servidor, sem a necessidade do cliente ter que fazer algum dos seguintes passos.

Provavelmente eu iria usar o Google Cloud ou AWS

Mas como é um teste, segue o passo a posso de como executar o projeto em sua máquina.

# Preparando o ambiente

## Pré-requisitos

Você precisará do Go, air (Não é obrigatório) e PostgreSQL instalados em sua máquina. 
A instalação para cada sistema operacional é diferente. 
Por isso eu vou deixar aqui um link para você instalar o Go na sua maquina.

## Instale o Golang

![](./img/go-install.png)

[Clique aqui para ir ao site](https://go.dev/doc/install/source)

## Ou você pode baixar o Golang

![](./img/go-download.png)

[Clique aqui para ir ao site](https://go.dev/doc/install)

## O levantamento do banco de dados

Lembra quando eu disse que você precisaria do PostgresSQL? 
Ele é o seu banco de dados. 

Você pode ver uma lista de downloads do site oficial do PostgreSQL
[Clique aqui](https://www.postgresql.org/download/).

## Para ir diretamente no download para o seu OS

## Instale o PostgreSQL no Windows
[Acesse clicando aqui](https://www.postgresql.org/download/windows/)

## Instale o PostgreSQL no Linux
[Acesse clicando aqui](https://www.postgresql.org/download/linux/)

## Instale o PostgreSQL no MacOS
[Acesse clicando aqui](https://www.postgresql.org/download/macosx/)

# Para iniciar o projeto em sua máquina
## Abra um terminal e execute os comandos
```bash
git clone https://github.com/jhamiltonjunior/todo-list-backend.git
cd todo-list-backend
```

`git clone <link_do_repositório>`
Vai baixar o repositório em sua maquina

`cd todo-list-backend`
Isso fará você navegar até o diretório que foi clonado


# Configurando banco de dados
## modifique o arquivo env

![](./img/dotenv.png)

Retire a extensão `.exemple` remova os comentarios e deixe o `DB_SOURCE` como mostrado abaixo, :

```
DB_SOURCE = "postgres://postgres:0000@localhost/vibbra?sslmode=disable"
```

## Mudando a senha do postgres

Em seu terminal linux digite:
```bash
sudo passwd postgres
Nova senha: 0000
Redigite a nova senha: 0000
```

Se você colocar uma senha diferente de `0000` não esqueça de altera no `DB_SOURCE` em `...postgres:0000@...`

Ou você pode ver [esse video](https://youtu.be/HbAJQ1WuGRE)

### No Windows
Veja [esse video](https://www.youtube.com/watch?v=YMj424O3eog) para alterar a senha no Windows.

O video não é meu Ok?!

### Tutorial para o MacOS
Veja [esse video](https://youtu.be/Ox3PSDloGu0) para alterar a senha no MacOS.

O video não é meu Ok?!

# Entre no Postgres

No seu terminal linux use o comando :
```bash
sudo -u postgres psql
```

## Agora você precisa criar o banco de dados e as tabelas

Ainda dentro do postgres rode o camando para criar o banco de dados:
```postgres
CREATE DATABASE vibbra;
```

`vibbra` é o nome do banco de dados, se você alterar o nome do banco de dados, altere também o `DB_SOURCE` em `.../vibbra...`

## Cria as tabelas

Para criar a tabela de Usuário, Lista e Item Da Lista, rode os seguintes comandos, na mesma ordem listada abaixo:

```postgres
CREATE TABLE user_schema(
  user_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  
  username VARCHAR(30) UNIQUE NOT NULL,
  fullname VARCHAR(30) NOT NULL,
  email TEXT UNIQUE NOT NULL,
  passwd VARCHAR(20) NOT NULL,
  
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP DEFAULT NOW()
);
```

```postgres
CREATE TABLE list_schema(
  list_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  
  title VARCHAR(60) NOT NULL,
  
  user_id INT,
  FOREIGN KEY(user_id)
    REFERENCES user_schema(user_id)
      ON DELETE CASCADE
        ON UPDATE CASCADE
);
```

```postgres
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
```

## Saia do Postgres 

```postgres
\q
```

## Talvez seja necessário instalar algumas dependências

Certifique-se de estar na raiz do projeto e rode os seguintes comando:

```go
go mod tidy

go get -u github.com/cosmtrek/air
go get -u github.com/gorilla/mux v1.8.0
go get -u github.com/jmoiron/sqlx v1.3.4
go get -u github.com/joho/godotenv v1.4.0
go get -u github.com/lib/pq

go mod tidy
```

O ```air``` é um live reload para o Golang, dessa forma você não terá que ficar iniciando o servido toda hora

# Para executar o Projeto (servidor)

Apenas para testar as rotas:
```bash
go run main.go
```
```go run main.go``` 
Vai executar o servidor 

Se você quiser executar o servidor em modo de desenvolvimento, use o `air`, na raiz do projeto rode o comando no terminal:
```
air
```

### A partir de agora você poderá acessar a aplicação