# Please

Please I ask you to see this file in a web version, I use markdown formatting to make the text more visible for you to read on the web, but seeing this here as a markdown file will give you a polluted view, or not if you already are used to seeing.

# Por favor

Por favor eu pesso que você veja esse arquivo em uma versão para web, eu ultilizo formatações markdown para deixar o texto mais visivel para você ler na web, mas vendo isso aqui como um arquivo markdown você terá uma visão poluída, ou não se você já estiver acostumado a ver.

# English Version
# How it works
During this markdown file you will see something like this:
`func (listItem *ListItem) CreateListItem() http.HandlerFunc`

`func` This keyword is used to declare a function, that is, the word before `()` is a function, so `CreateListItem()` is a occupation.

`(user *User)` means that the function `CreateListItem()` "is part" of the struct (Structure) User, if you understand Object Oriented Programming, we can call this here a method of the User class, although it's not, but if you're new to Go this will help you understand.

``http.HandlerFunc`` This is the function return.

So `CreateListItem` is a function that returns a `http.HandlerFunc` and is a "User class method" or User struct function.

# ListItem Struct
## type ListItem struct {}

Here will contain the items that were inside the "parent" List

That is, Every item belongs to a list, but not every list will have an item, at least it's not mandatory

So there can be empty lists

I mean: `/api/v{1}`

# ShowListItem Function
## func (listItem *ListItem) ShowListItem() http.HandlerFunc 

Will show all items in the list the WHERE here refers to the list_id, there is no lookup to see
only one item in a list, there was no such requirement, maybe that wasn't the point


## UpdateListItem function
## func (listItem *ListItem) UpdateListItem() http.HandlerFunc

Will update an item in a list.

You need to pass item id and list id.

In the code below you will see:
`params["id"] params["item_id"]`
Where `params["id"]` refers to list_id, that is, the id of the list.

And `params["item_id"]` refers to list_item_id, that is, the id of the item in the list



# UpdateListItem Function
## func (listItem *ListItem) UpdateListItem() http.HandlerFunc

This function is responsible for deleting an item from the list.

She is a little different from the others.
It won't just delete using `list_item_id`, it also needs `list_id`

That is, I will delete an item referring to the parent list, I would not need to pass the id
from the parent list, as there is only one id referring to `list_item_id`.

But maybe this is a safer way to delete an item from the list



# Portuguese (Brazil) Version


# Estrutura ListItem
## type ListItem struct {}

Aqui irá conter os itens que ficaram dentro da Lista "pai".

Ou seja, Todo item pertence a uma lista, mas nem toda lista terá um item, pelo menos não é obrigatório

Então pode haver listas vazias

Refiro-me: `/api/v{1}`

# ShowListItem
## func (listItem *ListItem) ShowListItem() http.HandlerFunc

Irá mostrar todos os itens que há dentro da lista o WHERE aqui é referente ao list_id, não existe uma pesquisa para ver apenas um item de uma lista, não havia esse requisito, talvez não era esse o objetivo


# Função UpdateListItem
## func (listItem *ListItem) UpdateListItem() http.HandlerFunc

Vai fazer o update de uma item de uma lista.

Você precisa passar o id do item e o id da lista.

No codigo abaixo você vai ver: `params["id"], params["item_id"]`

Onde `params["id"]` é referente a `list_id`, ou seja, o id da lista,

E `params["item_id"]` é refereinte a `list_item_id`, ou seja, o id do item que está na lista


# DeleteListItem
## func (listItem *ListItem) DeleteListItem() http.HandlerFunc

Esta função é responsavel por deletar um item da lista.

Ela é um pouco diferente das outras.
Ela não vai deletar apenas usando o `list_item_id`, ela pricisa também do `list_id`

Ou seja, eu vou deletar um item referente a lista pai, eu não precisaria passar o id
da lista pai, já que só exite um id referente ao `list_item_id`.

Mas talvez essa seja uma forma mais segura de deletar um item da lista

