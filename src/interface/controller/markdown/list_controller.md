# Please

Please I ask you to see this file in a web version, I use markdown formatting to make the text more visible for you to read on the web, but seeing this here as a markdown file will give you a polluted view, or not if you already are used to seeing.

# Por favor

Por favor eu pesso que você veja esse arquivo em uma versão para web, eu ultilizo formatações markdown para deixar o texto mais visivel para você ler na web, mas vendo isso aqui como um arquivo markdown você terá uma visão poluída, ou não se você já estiver acostumado a ver.

# English Version
# How it works
During this markdown file you will see something like this:
`func (list *List) CreateList() http.HandlerFunc`

`func` This keyword is used to declare a function, that is, the word before `()` is a function, so `CreateList()` is a occupation.

`(list *List)` means that the function `CreateList()` "is part" of the struct (Structure) User, if you understand Object Oriented Programming, we can call this here a method of the User class, although it's not, but if you're new to Go this will help you understand.

``http.HandlerFunc`` This is the function return.

So `CreateList` is a function that returns a `http.HandlerFunc` and is a "List class method" or List struct function.


# CreateList
## func (list *List) CreateList() http.HandlerFunc

This function will create the user by inserting his data into the database.

If everything goes correctly it will return a `JSON` with user data (minus `password`) and status code 201 created



# ShowListItem Function
## func (listItem *ListItem) ShowListItem() http.HandlerFunc 

Will show all items in the list the WHERE here refers to the list_id, there is no lookup to see
only one item in a list, there was no such requirement, maybe that wasn't the point


# DeleteList
## func (list *List) DeleteList() http.HandlerFunc 

DeleteList will delete a list from the database and if it really works, it will return a simple `JSON` informed that the list has been deleted


# Portuguese (Brazil) Version


# CreateList
## func (list *List) CreateList() http.HandlerFunc

Essa função vai criar o usuário inserindo os dados dele no banco de dados.

Se tudo ocorrer corretamente ela irá retornar um `JSON` com os dados do usário (menos o `password`) e o status code 201 created


# DeleteList() Function
## func (list *List) DeleteList() http.HandlerFunc

DeleteList irá deletar uma lista do banco de dados e se realmente funcionar, irá retornar um simples `JSON` informado que a lista foi deletada
