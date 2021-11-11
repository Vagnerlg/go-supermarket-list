package mongo

import (
	"errors"
	"log"
	"time"

	"github.com/vagnerlg/supermaketlist/src/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type item struct {
	Id          primitive.ObjectID `bson:"_id"`
	Product     string             `bson:"product"`
	Description string             `bson:"description"`
	Amount      int                `bson:"amount"`
	CheckedAt   time.Time          `bson:"checked_at"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type MongoItem struct{}

func (m MongoItem) Insert(i domain.Item) domain.Item {

	MongoItem := item{
		Id:          primitive.NewObjectID(),
		Product:     i.Product,
		Description: i.Description,
		Amount:      i.Amount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result, err := collection("list").InsertOne(ctx, MongoItem)

	if err != nil {
		log.Fatal(err)
		panic("Erro ao inserir novo item")
	}

	var oid primitive.ObjectID = result.InsertedID.(primitive.ObjectID)

	i.Id = oid.Hex()

	return i
}

func (m MongoItem) All() []domain.Item {
	filter := bson.D{{}}
	result, err := collection("list").Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var response []domain.Item

	for result.Next(ctx) {
		var i item
		err := result.Decode(&i)
		if err != nil {
			log.Fatal(err)
			continue
		}

		dItem := domain.Item{
			Id:          i.Id.Hex(),
			Product:     i.Product,
			Description: i.Description,
			Amount:      i.Amount,
			CheckedAt:   i.CheckedAt,
			CreatedAt:   i.CreatedAt,
			UpdatedAt:   i.UpdatedAt,
		}

		response = append(response, dItem)
	}

	return response
}

func (m MongoItem) First(id string) (domain.Item, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Item{}, errors.New("item not found")
	}

	i := item{}
	err = collection("list").FindOne(ctx, bson.D{{"_id", oid}}).Decode(&i)
	if err != nil {
		return domain.Item{}, errors.New("item not found")
	}

	return domain.Item{
		Id:          i.Id.Hex(),
		Product:     i.Product,
		Description: i.Description,
		Amount:      i.Amount,
		CheckedAt:   i.CheckedAt,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
	}, nil
}
