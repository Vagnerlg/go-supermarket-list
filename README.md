# Go supermarket list
O projeto é uma Api de lista de supermercado feita em golang com mongodb.

## Required

Ter o docker/docker-compose instalado

> Opcional: ter o golang instalado + vscode

## Configuração e instalação

Após clonar este repositorio entre no diretorio do projeto.

```sh
cd go-supermarket-list
```

Copiar o env dev 

```sh
cp .env.dev .env
```

É necessário buildar o projeto go-supermarket-list.

```sh
docker-compose run --rm golang go build src/main.go
```

Se tudo deu certo vc terar um arquivo com o nome de main na pasta raiz do projeto

![Alt text](./.docker/bash1.png?raw=true "Title")

Agora é só inicializar o serviço com docker-compose

```sh
docker-compose up -d
```

## Uso

A aplicação tem 3 rotas

- GET - localhost:3000/item
    Lista todos os itens cadastrados
- GET - localhost:3000/item/:id
    Lista um item expecifico se existir
- POST - localhost:3000/item
    Este rota tem que ter um corpo parecido com este

```json
{
    "product" : "coffee",
    "description" : "extra strong roast",
    "amount": 10
}
```