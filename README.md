# You simply need insert in your terminal this code:

## You need be in project root

~~~ bash
go run main.go
~~~

# Em um projeto real

Eu acredito que em um projeto real, eu mesmo iria hospedar o servidor, sem a necessidade do cliente ter que fazer algum dos seguintes passos. \
Provavelmente eu iria usar o Google Cloud ou AWS \
\
Mas como é um teste, segue o passo a posso de como executar o projeto em sua máquina.

# Preparando o ambiente

## Pré-requisitos
\
Você precisará do GO, air e PostgreSQL instalados em sua máquina. \
Existem diferentes maneiras de instalar o NodeJS na sua maquina.\
E para cada sistema operacional uma maneira diferente. \
Por isso eu vou listar alguns links do site oficial de instalação do NodeJS.

## A versão do Node que estou usando é a 16.14.0
Esta é a versão que eu estou usando, então é a versão que você vai baixar nos links abaixo, caso queira usar outra versão acesse o site oficial [clicando aqui](https://nodejs.org/en/download/)


## Instale o Node 16.14 no Windows
[Acesse clicando aqui](https://nodejs.org/dist/v16.14.0/node-v16.14.0-x86.msi)

## Instale o Node 16.14 no MacOS
[Acesse clicando aqui](https://nodejs.org/dist/v16.14.0/node-v16.14.0.pkg)

## Baixe o source code
[Acesse clicando aqui](https://nodejs.org/dist/v16.14.0/node-v16.14.0.tar.gz)

## Instale o Node na sua distribuição Linux
[Acesse clicando aqui](https://nodejs.org/en/download/package-manager/)

## O levantamento do banco de dados

Lembra quando eu disse que você precisaria do PostgresSQL? \
Ele é o seu banco de dados. \
\
Você pode ver uma lista de downloads do site oficial do PostgreSQL
[clicando aqui](https://www.postgresql.org/download/).

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
git clone <link_do_repositório>
cd vibbra-test
npm i
```

```git clone <link_do_repositório>``` \
Vai baixar o repositório em sua maquina

```cd vibbra-test``` \
Isso fará você navegar até o diretório que foi clonado

```npm i``` \
Irá instalar todas as dependências do projeto 

# Você precisará do banco de dados

## A seguir você verá os comandos para "levantar" o banco de dados




# Para executar o Projeto (servidor)

Ainda com o seu terminal aberto, execute: \
```npm start```

### A partir de agora você poderá acessar a aplicação