# Please

Please I ask you to see this file in a web version, I use markdown formatting to make the text more visible for you to read on the web, but seeing this here as a markdown file will give you a polluted view, or not if you already are used to seeing.

# Por favor

Por favor eu pesso que você veja esse arquivo em uma versão para web, eu ultilizo formatações markdown para deixar o texto mais visivel para você ler na web, mas vendo isso aqui como um arquivo markdown você terá uma visão poluída, ou não se você já estiver acostumado a ver.

# English Version
# How it works
During this markdown file you will see something like this:
`func (user *User) CreateUser() http.HandlerFunc`

`func` This keyword is used to declare a function, that is, the word before `()` is a function, so `CreateUser()` is a occupation.

`(user *User)` means that the function `CreateUser()` "is part" of the struct (Structure) User, if you understand Object Oriented Programming, we can call this here a method of the User class, although it's not, but if you're new to Go this will help you understand.

``http.HandlerFunc`` This is the function return.

So `CreateUser` is a function that returns a `http.HandlerFunc` and is a "User class method" or User struct function.

# User Struct
## type User struct {}
The User struct is responsible for getting the req.Body and inserting it into the database and the same is responsible for "porpulating" the JSON that returns from the database

Please you from the frontend, redirect the user to the route
`/api/v{n}/authenticate`
Here I just create the user, I don't have any JWT authenticate here

# controller.User.Name
I put Name, because if I put UserName when going to use
would have to call user.UserName and I don't like that
user.Name is already implied
```go
Name string `json:"username" db:"username"`
```

# CreateUser Function
## func (user *User) CreateUser() http.HandlerFunc

There is an error here in returning user data
it doesn't show the ID correctly, nor the Creation Date
even though the user was created
if you go in the route that shows all users you will see that
he was raised
Don't worry
this could be changed in new feature

```json
"user_id": 0,
"created_at": "",
"updated_at": ""
```

# ListAllUser Function
## func (user *User) ListAllUser() http.HandlerFunc

`ListAllUser()` will list ALL users that are registered in the database
that's right, if there are 1 thousand users I advise you to put another 8 or 16 GB's of RAM
on your machine.

In a new feature, a `DESC LIMIT {NUMBER} OFFSET {NUMBER}` could be placed in the query
This will prevent server crashes or slowdowns.

Can you imagine having to list 1000 users for 20 people at the same time?

Consider being very careful with this.

Why wasn't this implemented?

This function shouldn't even exist!
I ended up creating this function by accident, now it's a feature


# UpdateUser Function
## func (user *User) UpdateUser() http.HandlerFunc

This function will update the user data that is in the database.

I was using **Insomnia** and when I updated the `user 1` data it was no longer listed at the beginning of the function.

Why?!

The last `user` will be modified to the end of the `ListAllUser() function`

So don't be scared.

# If err != nil

the return after the error the application continues that prevents

``` go
if err != nil {
  return
}
```

# user.Password

I'm putting the `""` (empty string), to overwrite password, and don't display it to the end user `user.Password = ""`

Please do not use this in frontend application













# Portuguese (Brazil) Version

# Como funciona
Ao decorrer desse arquivo markdown você verá algo parecido com isso:
`func (usuário *Usuário) CreateUser() http.HandlerFunc`

`func` Essa palavra chave é usada para declarar uma função, ou seja, a palavra que vem antes do `()` é uma função, logo `CreateUser()` é uma função.

`(user *User)` significa que a função `CreateUser()` "faz parte" da struct (Estrutura) User, se você entende de Programação Orientada a Objetos, podemos chamar isso aqui de um método da classe User, apesar de não ser, mas se você é novo no Go isso vai te ajudar a entender.

`http.HandlerFunc` Isso aqui é o retorno da função.

Então `CreateUser` é uma função que retorna um `http.HandlerFunc` e é um "metodo da classe User" ou função da User struct.

# Estrutura User
## type User struct {}

User é o responsavel por pegar o req.Body e inserir no banco de dados e o mesmo é responsavel por "porpular" o JSON que retorna do banco de dados.

Por favor, você do frontend, redirecione o usuário para a rota
`/api/v{n}/authenticate`

Aqui eu somente crio o usuário, não tenho nenhum JWT authenticate aqui

# controller.User.Name

Coloquei Name, porque se eu colocasse UserName quando fosse usar iria ter que chamar `user.UserName` e eu não gosto disso.

`user.Name` já fica subentendido.

```go
Name string `json:"username" db:"username"`
```

# Função CreateUser
## func (user *User) CreateUser() http.HandlerFunc

Existe um erro aqui no retorno dos dados do usuario
ele não mostra o ID corretamente, nem a Data de criação mesmo que o user foi criado se você for na rota que mostra todos os usuarios você vai ver que ele foi criado.

Não se preocupe.

Isso poderia ser mudado em uma nova feature

```json
"user_id": 0,
"created_at": "",
"updated_at": ""
```

# Função ListAllUser
## func (user *User) ListAllUser() http.HandlerFunc

`ListAllUser()` vai listar TODOS os usários que estão cadastrados no banco de dados
isso mesmo, se houver 1 mil usários eu aconselho você colocar mais uns 8 ou 16GB's de RAM
na sua máquina.

Em uma nova feature poderia ser colocado um `DESC LIMIT {NUMBER} OFFSET {NUMBER}` na query
Isso ira impedir travamentos ou lentidão no servidor.

Imagina ter que listar 1 mil usuários para 20 pessoas ao mesmo tempo?

Considere tomar bastante cuidado com isso.

Porque isso não foi implementado?

Essa função nem deveria exitir!
Acabei criando essa função por acidente, agora é uma feature


# Função UpdateUser
## func (user *User) UpdateUser() http.HandlerFunc

Essa função vai atualizar os dados do usuário que estão no banco de dados.

Eu estava usando o **Insomnia** e quando eu atualizei o dado do `user 1` ele não era mais listado no inicio da função.

Porquê?!

O ultimo `user` a ser modificado vai para o final da `ListAllUser() function`

Portanto não se assuste.


# If err != nil

o `return` impede que após o erro a aplicação continue executando
``` go
if err != nil {
  return
}
```

# user.Password
`user.Password = ""`
Eu estou colocando o `""` (string vazia), para sobrescrever o password, e não exibir para ao usuário final.

Por favor não use isso no frontend
