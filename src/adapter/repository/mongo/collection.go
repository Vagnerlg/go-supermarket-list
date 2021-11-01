package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func collection() *mongo.Collection {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/").SetAuth(
		options.Credential{
			Username: "root",
			Password: "123456",
		},
	)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		panic("Erro ao conectar com o banco de dados")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		panic("Erro ao conectar com o banco de dados")
	}

	return client.Database("supermaketlist").Collection("list")
}
