package mongo

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func collection(model string) *mongo.Collection {
	hostConn := fmt.Sprintf("mongodb://%v/", os.Getenv("DB_HOST"))

	fmt.Println("env colection", hostConn)

	clientOptions := options.Client().ApplyURI(hostConn).SetAuth(
		options.Credential{
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
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

	return client.Database("supermaketlist").Collection(model)
}
